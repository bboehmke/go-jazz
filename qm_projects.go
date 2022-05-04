package jazz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/sync/errgroup"
)

type QMProject struct {
	Title string
	Alias string

	qm *QMApplication
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

// QMGetList object of the given type
func QMGetList[T QMObject](proj *QMProject, ids []string) ([]T, error) {
	return Chan2List[T](func(ch chan T) error {
		return QMGetListChan[T](proj, ids, ch)
	})
}

// QMGetListChan object of the given type returned via a channel
func QMGetListChan[T QMObject](proj *QMProject, ids []string, results chan T) error {
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

	response, err := proj.qm.client.Get(spec.GetURL(proj, id), "application/json", false)
	if err != nil {
		return value, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return value, fmt.Errorf("failed to get %s: %s", spec.ResourceID, response.Status)
	}

	bla, _ := io.ReadAll(response.Body) // TODO remove
	buffer := bytes.NewBuffer(bla)

	var tmpData map[string]T
	err = json.NewDecoder(buffer).Decode(&tmpData)
	if err != nil {
		return value, fmt.Errorf("failed to parse %s: %w", spec.ResourceID, err)
	}

	tmpData[spec.ResourceID].SetProj(proj)
	return tmpData[spec.ResourceID], nil
}
