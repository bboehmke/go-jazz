// Copyright 2022 Benjamin Böhmke <benjamin@boehmke.net>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jazz

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"

	"golang.org/x/sync/errgroup"
)

type QMProject struct {
	Title string
	Alias string

	qm *QMApplication
}

// NewUUID returns a new UUID generated on the server
func (p *QMProject) NewUUID() (string, error) {
	response, err := p.qm.client.Get(
		"qm/service/com.ibm.rqm.integration.service.IIntegrationService/UUID/new",
		"application/json",
		true)
	if err != nil {
		return "", fmt.Errorf("failed to get UUID: %w", err)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to get UUID: %w", err)
	}
	return string(data), nil
}

// QMList object of the given type
func QMList[T QMObject](proj *QMProject, filter QMFilter) ([]T, error) {
	return Chan2List[T](func(ch chan T) error {
		return QMListChan[T](proj, filter, ch)
	})
}

// QMListChan object of the given type returned via a channel
func QMListChan[T QMObject](proj *QMProject, filter QMFilter, results chan T) error {
	spec := (*new(T)).Spec()

	// load object returned by list
	requestChan := make(chan feedEntry, 100)
	g := new(errgroup.Group)
	for i := 0; i < proj.qm.client.Worker; i++ {
		g.Go(func() error {
			for entry := range requestChan {
				obj, err := QMGet[T](proj, entry.Id)
				if err != nil {
					return err
				}
				results <- obj
			}
			return nil
		})
	}

	// get initial URL request
	url, err := spec.ListURL(proj, filter)
	if err != nil {
		return err
	}

	// request object list
	err = proj.qm.client.requestFeed(url, requestChan, false)
	if err != nil {
		return err
	}

	// stop background worker and wait for work is done
	close(requestChan)
	err = g.Wait()
	return err
}

// qmGetList object of the given type
func qmGetList[T QMObject](proj *QMProject, ids []string) ([]T, error) {
	return Chan2List[T](func(ch chan T) error {
		return qmGetListChan[T](proj, ids, ch)
	})
}

// qmGetListChan object of the given type returned via a channel
func qmGetListChan[T QMObject](proj *QMProject, ids []string, results chan T) error {
	// load object returned by list
	idChan := make(chan string, proj.qm.client.Worker*2)
	g := new(errgroup.Group)
	for i := 0; i < proj.qm.client.Worker; i++ {
		g.Go(func() error {
			for id := range idChan {
				obj, err := QMGet[T](proj, id)
				if err != nil {
					return err
				}
				results <- obj
			}
			return nil
		})
	}

	// add ids to channel
	for _, id := range ids {
		idChan <- id
	}

	// stop background worker and wait for work is done
	close(idChan)
	return g.Wait()
}

// QMGet object of the given type
func QMGet[T QMObject](proj *QMProject, id string) (T, error) {
	var value T
	spec := value.Spec()

	response, err := proj.qm.client.Get(spec.GetURL(proj, id),
		"application/xml", false)
	if err != nil {
		return value, fmt.Errorf("failed to get %s: %w", spec.ResourceID, err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return value, fmt.Errorf("failed to get %s: %s", spec.ResourceID, response.Status)
	}

	err = xml.NewDecoder(response.Body).Decode(&value)
	if err != nil {
		return value, fmt.Errorf("failed to parse %s: %w", spec.ResourceID, err)
	}

	value.SetProj(proj)
	return value, nil
}

// QMGetFilter object of the given filter
func QMGetFilter[T QMObject](proj *QMProject, filter QMFilter) (T, error) {
	var nul T
	entries, err := QMList[T](proj, filter)
	if err != nil {
		return nul, err
	}
	if len(entries) == 0 {
		return nul, fmt.Errorf("no object matching filter found")
	}
	if len(entries) > 1 {
		return nul, fmt.Errorf("more then one object (%d) found", len(entries))
	}

	return entries[0], nil
}

// QMSave object of the given type
func QMSave[T QMObject](proj *QMProject, obj T) (T, error) {
	// create a new resource URL if not already set
	if obj.Ref().Href == "" {
		uuid, err := proj.NewUUID()
		if err != nil {
			return obj, fmt.Errorf("failed to save object: %w", err)
		}

		obj.SetRef(obj.Spec().GetURL(proj, "go_"+uuid))
	}

	// encode object
	data := obj.Spec().DumpXml(obj)

	// send request to server
	response, err := proj.qm.client.Put(obj.Ref().Href, "application/xml", bytes.NewBuffer(data))
	if err != nil {
		return obj, fmt.Errorf("failed to save object: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode >= 300 {
		return obj, fmt.Errorf("failed to save object: %s", response.Status)
	}

	// load created object from server
	return QMGet[T](proj, obj.Ref().Href)
}

// UploadAttachment with the given file name and content
func (p *QMProject) UploadAttachment(fileName string, fileReader io.Reader) (*QMAttachment, error) {
	// get new UUID
	uuid, err := p.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to upload attachment: %w", err)
	}

	// create multipart writer
	r, w := io.Pipe()
	defer r.Close()
	m := multipart.NewWriter(w)

	// copy file content to multipart writer
	go func() {
		part, err := m.CreateFormFile("file", fileName)
		if err != nil {
			// The error is returned from read on the pipe.
			w.CloseWithError(err)
			return
		}
		if _, err := io.Copy(part, fileReader); err != nil {
			// The error is returned from read on the pipe.
			w.CloseWithError(err)
			return
		}

		// add closing boundary (missing in multipart writer)
		fmt.Fprintf(w, "\r\n--%s--", m.Boundary())

		// The http.Post function reads the pipe until
		// an error or EOF. Close to return an EOF to
		// http.Post.
		w.Close()
	}()

	// send response to server
	url := new(QMAttachment).Spec().GetURL(p, "go_"+uuid)
	response, err := p.qm.client.Put(url, m.FormDataContentType(), r)
	if err != nil {
		return nil, fmt.Errorf("failed to save object: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to save object: %s", response.Status)
	}

	// load uploaded attachment
	return QMGet[*QMAttachment](p, url)
}
