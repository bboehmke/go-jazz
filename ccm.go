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
	"context"
	"errors"
	"reflect"

	"github.com/beevik/etree"
	"golang.org/x/sync/errgroup"
)

// CCMErrorEmptyResponse is returned if an empty XML response was received
var CCMErrorEmptyResponse = errors.New("empty response XML -> item maybe deleted")

// CCMApplication interface
type CCMApplication struct {
	client *Client
}

// Name of application
func (a *CCMApplication) Name() string {
	return "Change and Configuration Management"
}

// ID of application
func (a *CCMApplication) ID() string {
	return "ccm"
}

// Client instance used for communication
func (a *CCMApplication) Client() *Client {
	return a.client
}

// CCMList object of the given type
func CCMList[T CCMObject](ctx context.Context, ccm *CCMApplication, filter CCMFilter) ([]T, error) {
	return Chan2List[T](func(ch chan T) error {
		return CCMListChan[T](ctx, ccm, filter, ch)
	})
}

// CCMListChan object of the given type returned via a channel
func CCMListChan[T CCMObject](ctx context.Context, ccm *CCMApplication, filter CCMFilter, results chan T) error {
	// load object returned by list
	requestChan := make(chan string, 100*2)
	g := new(errgroup.Group)
	for i := 0; i < ccm.client.Worker; i++ {
		g.Go(func() error {
			for id := range requestChan {
				obj, err := CCMGet[T](ctx, ccm, id)
				if err != nil {
					return err
				} else {
					results <- obj
				}
			}
			return nil
		})
	}

	err := CCMListEntryChan[T](ctx, ccm, filter, requestChan)
	if err != nil {
		return err
	}

	// stop background worker and wait for work is done
	close(requestChan)
	err = g.Wait()
	return err
}

// CCMListEntryChan queries only the references of objects (without loading)
func CCMListEntryChan[T CCMObject](ctx context.Context, ccm *CCMApplication, filter CCMFilter, results chan string) error {
	spec := (*new(T)).Spec()

	// get initial URL request
	url, err := spec.ListURL(filter)
	if err != nil {
		return err
	}

	// request list until last page reached
	for url != "" {
		resp, root, err := ccm.client.getEtree(ctx, url, "application/xml", //nolint:bodyclose
			"failed get element list", 0)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return ccmResponse2error(root)
		}

		// extract item IDs from result
		entries := root.FindElements(spec.ElementID + "/itemId")
		for _, entry := range entries {
			results <- entry.Text()
		}

		if len(entries) >= 100 {
			url = root.SelectAttrValue("href", "")
		} else {
			url = ""
		}
	}
	return nil
}

// CCMGet object of the given type
func CCMGet[T CCMObject](ctx context.Context, ccm *CCMApplication, id string) (T, error) {
	var value T
	spec := value.Spec()

	err := ccm.get(ctx, spec, reflect.ValueOf(&value), id)
	return value, err
}

// CCMGetFilter object of the given filter
func CCMGetFilter[T CCMObject](ctx context.Context, ccm *CCMApplication, filter CCMFilter) (T, error) {
	return listOnlyOnce(CCMList[T](ctx, ccm, filter))
}

func (a *CCMApplication) get(ctx context.Context, spec *CCMObjectSpec, value reflect.Value, id string) error {
	resp, root, err := a.client.getEtree(ctx,
		spec.GetURL(id),
		"application/xml",
		"failed get element "+id, 0)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return ccmResponse2error(root)
	}

	// catch empty elements
	if len(root.ChildElements()) == 0 {
		return CCMErrorEmptyResponse
	}

	return spec.Load(a, value, root.FindElement(spec.ElementID))
}

func ccmResponse2error(root *etree.Element) error {
	element := root.FindElement("//qm:message/text()")
	if element != nil {
		return errors.New(element.Text())
	}

	element = root.FindElement("/error/text()")
	if element != nil {
		return errors.New(element.Text())
	}
	return errors.New("unknown error")
}
