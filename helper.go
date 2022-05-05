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
