package jazz

import (
	"errors"
	"reflect"
	"sync"

	"golang.org/x/sync/errgroup"
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

// CCMList object of the given type
func CCMList[T CCMObject](ccm *CCMApplication, filter CCMFilter) ([]T, error) {
	results := make(chan T)
	objects := make([]T, 0)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for obj := range results {
			objects = append(objects, obj)
		}
		wg.Done()
	}()

	err := CCMListChan[T](ccm, filter, results)
	close(results)

	// wait for all results to be handled
	wg.Wait()
	return objects, err
}

// CCMListChan object of the given type returned via a channel
func CCMListChan[T CCMObject](ccm *CCMApplication, filter CCMFilter, results chan T) error {
	spec := (*new(T)).Spec()

	// load object returned by list
	requestChan := make(chan string, 100*2)
	g := new(errgroup.Group)
	for i := 0; i < ccm.client.Worker; i++ {
		g.Go(func() error {
			for id := range requestChan {
				var obj T
				if err := ccm.get(spec, reflect.ValueOf(&obj), id); err != nil {
					return err
				} else {
					results <- obj
				}
			}
			return nil
		})
	}

	// get initial URL request
	url, err := spec.ListURL(filter)
	if err != nil {
		return err
	}

	var counter uint64 // TODO remove
	// request list until last page reached
	for url != "" {
		resp, root, err := ccm.client.SimpleGet(url, "application/xml", //nolint:bodyclose
			"failed get element list", 0)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return errors.New(root.FindElement("//qm:message/text()").Text())
		}

		// extract item IDs from result
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
	err = g.Wait()
	return err
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
