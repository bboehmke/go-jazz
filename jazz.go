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
	"fmt"
	"io"
	"net/http"
)

type Application interface {
	Name() string
	ID() string
	Client() *Client
}

type App struct {
	Application
}

func (a *App) RootServices() *RootService {
	return &RootService{
		client: a.Client(),
		base:   a.ID(),
	}
}

// Error for responses of jazz server
type Error struct {
	Msg      string
	Details  string
	PostData []byte
}

func (e Error) Error() string {
	return e.Msg
}

func errorFromResponse(msg string, response *http.Response, data []byte) error {
	body, _ := io.ReadAll(response.Body)
	return &Error{
		Msg:      fmt.Sprintf("%s: %s", msg, response.Status),
		Details:  string(body),
		PostData: data,
	}
}
