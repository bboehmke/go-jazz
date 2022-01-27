package jazz

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

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

// List object of the given type
func (a *CCMApplication) List(data interface{}) error {
	// get specification of object
	spec, err := LoadObjectSpec(data)
	if err != nil {
		return fmt.Errorf("failed to list elements: %w", err)
	}

	// prepare list for results
	objList := reflect.ValueOf(spec.NewList())

	// load object returned by list
	requestChan := make(chan string, 100*2)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var getErr error
	for i := 0; i < a.client.Worker; i++ {
		wg.Add(1)
		go func() {
			for id := range requestChan {
				val, err := a.get(spec, id)

				mutex.Lock()
				if err == nil {
					objList = reflect.Append(objList, reflect.ValueOf(val).Elem())
				} else {
					getErr = err
				}
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	// get initial URL request
	url := spec.ListURL()

	var counter uint64 // TODO remove

	// request list until last page reached
	for url != "" {
		resp, root, err := a.client.SimpleGet(url, "application/xml",
			"failed get element list", 0)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return errors.New(root.FindElement("//qm:message/text()").Text())
		}

		entries := root.FindElements(spec.ElementID + "/itemId")
		for _, entry := range entries {
			requestChan <- entry.Text()
			counter++ // TODO remove
		}

		if len(entries) >= 100 {
			url = root.SelectAttrValue("href", "")
		} else {
			url = ""
		}

		// TODO remove
		if counter >= 200 {
			break
		}
	}

	// stop background worker and wait for work is done
	close(requestChan)
	wg.Wait()

	if getErr != nil {
		return getErr
	}

	// write object list back
	reflect.ValueOf(data).Elem().Set(objList)
	return nil
}

// Get object with the given id
func (a *CCMApplication) Get(data interface{}, id string) error {
	spec, err := LoadObjectSpec(data)
	if err != nil {
		return fmt.Errorf("failed to list elements: %w", err)
	}

	entry, err := a.get(spec, id)

	reflect.ValueOf(data).Elem().Set(reflect.ValueOf(entry).Elem())

	return err
}

func (a *CCMApplication) get(spec *ObjectSpec, id string) (interface{}, error) {
	resp, root, err := a.client.SimpleGet(spec.GetURL(id),
		"application/xml",
		"failed get element "+id, 0)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(root.FindElement("//qm:message/text()").Text())
	}

	return spec.Load(root.FindElement(spec.ElementID))
}
