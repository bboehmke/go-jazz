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
	dataType := reflect.TypeOf(data)

	// get specification of object
	spec, err := LoadObjectSpec(reflect.TypeOf(data))
	if err != nil {
		return fmt.Errorf("failed to list elements: %w", err)
	}

	// prepare list for results
	objList := reflect.MakeSlice(reflect.SliceOf(dataType.Elem().Elem()), 0, 0)

	// load object returned by list
	requestChan := make(chan string, 100*2)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var getErr error
	for i := 0; i < a.client.Worker; i++ {
		wg.Add(1)
		go func() {
			for id := range requestChan {

				value := reflect.New(dataType.Elem().Elem())
				err := a.get(spec, value, id)

				mutex.Lock()
				if err == nil {
					objList = reflect.Append(objList, value.Elem())
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
		resp, root, err := a.client.SimpleGet(url, "application/xml", //nolint:bodyclose
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

// CCMGet object of the given type
func CCMGet[T CCMObject](ccm *CCMApplication, id string) (T, error) {
	var value T
	spec := value.Spec()

	resp, root, err := ccm.client.SimpleGet(spec.GetURL(id),
		"application/xml",
		"failed get element "+id, 0)
	if err != nil {
		return value, err
	}
	if resp.StatusCode != 200 {
		return value, errors.New(root.FindElement("//qm:message/text()").Text())
	}

	return value, spec.Load(ccm, reflect.ValueOf(&value), root.FindElement(spec.ElementID))
}

// Get object with the given id
func (a *CCMApplication) Get(data interface{}, id string) error {
	spec, err := LoadObjectSpec(reflect.TypeOf(data))
	if err != nil {
		return fmt.Errorf("failed to load object spec: %w", err)
	}

	return a.get(spec, reflect.ValueOf(data), id)
}

func (a *CCMApplication) get(spec *ObjectSpec, value reflect.Value, id string) error {
	resp, root, err := a.client.SimpleGet(spec.GetURL(id),
		"application/xml",
		"failed get element "+id, 0)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(root.FindElement("//qm:message/text()").Text())
	}

	return spec.Load(a, value, root.FindElement(spec.ElementID))
}
