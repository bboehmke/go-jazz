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
	"fmt"
)

type feedEntry struct {
	Id    string
	Title string
	Alias string
}

type rawEntry struct {
	Id    string `json:"id"`
	Title struct {
		Content string `json:"content"`
	} `json:"title"`
	Content struct {
		Project struct {
			Alias struct {
				Content string `json:"content"`
			} `json:"alias"`
		} `json:"project"`
	} `json:"content"`
}

func (e *rawEntry) entry() feedEntry {
	return feedEntry{
		Id:    e.Id,
		Title: e.Title.Content,
		Alias: e.Content.Project.Alias.Content,
	}
}

type link struct {
	Name string `json:"rel"`
	Href string `json:"href"`
}

type linkList []link

func (l *linkList) get(name string) string {
	for _, link := range *l {
		if link.Name == name {
			return link.Href
		}
	}
	return ""
}

type subFeed struct {
	Entries []rawEntry
	NextURL string
}

func (s *subFeed) UnmarshalJSON(p []byte) error {
	feed := struct {
		Entries []rawEntry `json:"Entry"`
		Links   linkList   `json:"link"`
	}{}
	err := json.Unmarshal(p, &feed)

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		feed2 := struct {
			Entry rawEntry `json:"Entry"`
			Links linkList `json:"link"`
		}{}
		err = json.Unmarshal(p, &feed2)
		if err == nil {
			s.Entries = []rawEntry{feed2.Entry}
			s.NextURL = feed2.Links.get("next")
		}
	} else {
		s.Entries = feed.Entries
		s.NextURL = feed.Links.get("next")
	}
	return err
}

type rawFeed struct {
	Feed subFeed `json:"feed"`
}

func (c *Client) requestFeed(url string, entries chan feedEntry, noGc bool) error {
	// request list until last page reached
	for url != "" {
		response, err := c.Get(url, "application/json", noGc)
		if err != nil {
			return err
		}

		var feed rawFeed
		err = json.NewDecoder(response.Body).Decode(&feed)
		response.Body.Close()
		if err != nil {
			return fmt.Errorf("failed to parse feed: %w", err)
		}
		for _, entry := range feed.Feed.Entries {
			entries <- entry.entry()
		}
		url = feed.Feed.NextURL
	}
	return nil
}
