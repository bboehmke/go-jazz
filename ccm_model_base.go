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
	"reflect"
	"sync"
	"time"
)

var CCMBaseObjectType = reflect.TypeOf(CCMBaseObject{})

// CCMObject describes a CCM object implementation
type CCMObject interface {
	Spec() *CCMObjectSpec
}

// CCMLoadableObject is only implemented by objects that are loadable
type CCMLoadableObject interface {
	CCMObject
	Load(ctx context.Context) error
}

type CCMBaseObject struct {
	// Common fields of every object
	// https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Common_properties

	// The UUID representing the item in storage. This is technically an internal
	// detail, and resources should mostly be referred to by their unique URLs.
	// In some cases the itemId may be the only unique identifier, however.
	ItemId string `jazz:"itemId"`

	// An MD5 hash of the URI for this element
	UniqueId string `jazz:"uniqueId"`

	// The UUID of the state for this item in storage. This is an internal detail.
	StateId string `jazz:"stateId"`

	// The UUID of a context object used for read access. This is an internal detail.
	ContextId string `jazz:"contextId"`

	// The timestamp of the last modification date of this resource.
	Modified *time.Time `jazz:"modified"`

	// A boolean indicating whether the resource is "archived". Archived
	// resources are typically hidden from the UI and filtered out of queries.
	Archived bool `jazz:"archived"`

	ReportableUrl string `jazz:"reportableUrl"`

	ModifiedBy *CCMContributor `jazz:"modifiedBy"`

	// init ensures elements are only loaded once
	init sync.Once

	// ccm Application instance used for interactions with the server
	ccm *CCMApplication
}

// String returns the ItemId of this object (used for filter)
func (o *CCMBaseObject) String() string {
	return o.ItemId
}

// setCCM application used for read and write actions
func (o *CCMBaseObject) setCCM(ccm *CCMApplication) {
	o.ccm = ccm
}

// loadFields of the given object
func (o *CCMBaseObject) loadFields(ctx context.Context, fields ...interface{}) error {
	for _, field := range fields {
		if fields, ok := field.([]CCMLoadableObject); ok {
			for _, f := range fields {
				if err := f.Load(ctx); err != nil {
					return err
				}
			}
		} else if f, ok := field.(CCMLoadableObject); ok {
			if reflect.ValueOf(f).IsNil() {
				continue
			}
			if err := f.Load(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}
