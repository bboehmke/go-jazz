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
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"golang.org/x/sync/errgroup"
)

type FeedEntry struct {
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

func (e *rawEntry) entry() FeedEntry {
	return FeedEntry{
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
	LastURL string
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
			s.LastURL = feed.Links.get("last")
		}
	} else {
		s.Entries = feed.Entries
		s.NextURL = feed.Links.get("next")
		s.LastURL = feed.Links.get("last")
	}
	return err
}

type rawFeed struct {
	Feed subFeed `json:"feed"`
}

func (c *Client) requestFeed(ctx context.Context, feedUrl string, entries chan FeedEntry, noGc bool) error {
	// request list until last page reached
	for feedUrl != "" {
		var err error
		feedUrl, _, err = c.doRequestFeed(ctx, feedUrl, entries, noGc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) requestFeedFast(ctx context.Context, feedUrl string, entries chan FeedEntry, noGc bool) error {
	// first request to get page count
	next, last, err := c.doRequestFeed(ctx, feedUrl, entries, noGc)
	if err != nil {
		return err
	}

	// single page -> stop here
	if next == "" {
		return nil
	}

	// parse URL of last entry
	parsedUrl, err := url.Parse(last)
	if err != nil {
		return fmt.Errorf("failed to parse last page URL: %w", err)
	}
	query := parsedUrl.Query()

	// extract index of last page
	lastPage, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		return fmt.Errorf("failed to parse page parameter: %w", err)
	}

	// handle page requests in parallel
	pageUrls := make(chan string, lastPage)
	var group errgroup.Group
	for i := 0; i < c.Worker; i++ {
		group.Go(func() error {
			for pageUrl := range pageUrls {
				_, _, err := c.doRequestFeed(ctx, pageUrl, entries, noGc)
				if err != nil {
					return err
				}
			}
			return nil
		})
	}

	// build page URLs
	for i := 1; i <= lastPage; i++ {
		query.Set("page", strconv.Itoa(i))
		parsedUrl.RawQuery = query.Encode()
		pageUrls <- parsedUrl.String()
	}

	// wait for result
	close(pageUrls)
	err = group.Wait()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) doRequestFeed(ctx context.Context, url string, entries chan FeedEntry, noGc bool) (string, string, error) {
	response, err := c.get(ctx, url, "application/json", noGc)
	if err != nil {
		return "", "", err
	}

	var feed rawFeed
	err = json.NewDecoder(response.Body).Decode(&feed)
	response.Body.Close()
	if err != nil {
		return "", "", fmt.Errorf("failed to parse feed: %w", err)
	}
	for _, entry := range feed.Feed.Entries {
		select {
		case <-ctx.Done():
			return "", "", ctx.Err()
		case entries <- entry.entry():
		}
	}

	return feed.Feed.NextURL, feed.Feed.LastURL, nil
}
