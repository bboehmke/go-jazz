package jazz

import (
	"encoding/json"
	"sync"
)

// Chan2List converts a channel to a slice
func Chan2List[T any](f func(ch chan T) error) ([]T, error) {
	// create channel and return slice
	ch := make(chan T)
	entries := make([]T, 0)

	// add all entries to slice
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for entry := range ch {
			entries = append(entries, entry)
		}
		wg.Done()
	}()

	err := f(ch)
	close(ch)

	// wait for all entries to be handled
	wg.Wait()
	return entries, err
}

// UnmarshalJSONOptionalList handles broken list in json on if single response
func UnmarshalJSONOptionalList[T any](b []byte) ([]T, error) {
	var entries []T
	err := json.Unmarshal(b, &entries)
	if err == nil {
		return entries, nil
	} else if _, ok := err.(*json.UnmarshalTypeError); ok {
		var entry T
		err = json.Unmarshal(b, &entry)
		if err == nil {
			entries = append(entries, entry)
		}
	}
	return entries, err
}
