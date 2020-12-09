package jazz

import "time"

type BaseObject struct {
	// Common fields of every object
	// https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#Common_properties

	//  The UUID representing the item in storage. This is technically an internal detail, and resources should mostly be referred to by their unique URLs. In some cases the itemId may be the only unique identifier, however.
	ItemId string `jazz:"itemId"`

	// An MD5 hash of the URI for this element
	UniqueId string `jazz:"uniqueId"`

	// The UUID of the state for this item in storage. This is an internal detail.
	StateId string `jazz:"stateId"`

	// The UUID of a context object used for read access. This is an internal detail.
	ContextId string `jazz:"contextId"`

	// The timestamp of the last modification date of this resource.
	Modified time.Time `jazz:"modified"`

	// A boolean indicating whether or not the resource is "archived". Archived resources are typically hidden from the UI and filtered out of queries.
	Archived bool `jazz:"archived"`

	ReportableUrl string `jazz:"reportableUrl"`

	ModifiedBy *Contributor `jazz:"modifiedBy"`
}
