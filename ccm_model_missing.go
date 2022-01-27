package jazz

import "time"

// This file contains objects that are missing in the documentation

// Item The only time you're likely to see a raw Item is when using the referencedItem
// field of a Reference. Most of the time you'll want to fetch whichever concrete item
// type is represented by this artifact (e.g. a Work Item). The only standard field here
// likely to be useful is itemId, which can be used to look up the concrete element.
// This element is always contained in a com.ibm.team.links.Reference, and represents
// whether the reference is by uri or by itemId.
// see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_repository_Item
type Item struct {
	BaseObject `jazz_resource:"foundation" jazz_type:"com.ibm.team.repository.Item"`

	// Type of item
	ItemType string `jazz:"itemType"`

	// The UUID representing the item in storage
	ItemId bool `jazz:"itemId"`
}

// BooleanExtensionEntry for additional entries
type BooleanExtensionEntry struct {
	BaseObject `jazz_resource:"workitem" jazz_type:"com.ibm.team.repository.BooleanExtensionEntry"`

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value bool `jazz:"value"`
}

// IntExtensionEntry for additional entries
type IntExtensionEntry struct {
	BaseObject `jazz_resource:"workitem" jazz_type:"com.ibm.team.repository.IntExtensionEntry"`

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value int `jazz:"value"`
}

// LongExtensionEntry for additional entries
type LongExtensionEntry struct {
	BaseObject `jazz_resource:"workitem" jazz_type:"com.ibm.team.repository.LongExtensionEntry"`

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value int64 `jazz:"value"`
}

// StringExtensionEntry for additional entries
type StringExtensionEntry struct {
	BaseObject `jazz_resource:"workitem" jazz_type:"com.ibm.team.repository.StringExtensionEntry"`

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value string `jazz:"value"`
}

// MediumStringExtensionEntry for additional entries
type MediumStringExtensionEntry struct {
	BaseObject `jazz_resource:"workitem" jazz_type:"com.ibm.team.repository.MediumStringExtensionEntry"`

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value string `jazz:"value"`
}

// LargeStringExtensionEntry for additional entries
type LargeStringExtensionEntry struct {
	BaseObject `jazz_resource:"workitem" jazz_type:"com.ibm.team.repository.LargeStringExtensionEntry"`

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value string `jazz:"value"`
}

// TimestampExtensionEntry for additional entries
type TimestampExtensionEntry struct {
	BaseObject `jazz_resource:"workitem" jazz_type:"com.ibm.team.repository.TimestampExtensionEntry"`

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value time.Time `jazz:"value"`
}

// BigDecimalExtensionEntry for additional entries
type BigDecimalExtensionEntry struct {
	BaseObject `jazz_resource:"workitem" jazz_type:"com.ibm.team.repository.BigDecimalExtensionEntry"`

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value float64 `jazz:"value"`
}

// ItemExtensionEntry for additional entries
type ItemExtensionEntry struct {
	BaseObject `jazz_resource:"workitem" jazz_type:"com.ibm.team.repository.ItemExtensionEntry"`

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value *Item `jazz:"value"`
}

// MultiItemExtensionEntry for additional entries
type MultiItemExtensionEntry struct {
	BaseObject `jazz_resource:"workitem" jazz_type:"com.ibm.team.repository.MultiItemExtensionEntry"`

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value []*Item `jazz:"value"`
}
