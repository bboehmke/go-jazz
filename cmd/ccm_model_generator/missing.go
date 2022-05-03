package main

// This file contains objects that are missing in the documentation
var missingObjects = []Model{
	{
		LinkRef: "com_ibm_team_repository_Item",
		Description: []string{
			"Item The only time you're likely to see a raw Item is when using the referencedItem",
			"field of a Reference. Most of the time you'll want to fetch whichever concrete item",
			"type is represented by this artifact (e.g. a Work Item). The only standard field here",
			"likely to be useful is itemId, which can be used to look up the concrete element.",
			"This element is always contained in a com.ibm.team.links.Reference, and represents",
			"whether the reference is by uri or by itemId.",
		},
		ResourceID: "foundation",
		TypeID:     "com.ibm.team.repository.Item",
		Fields: []Field{
			{
				Name:        "itemType",
				Type:        "xs:string",
				Description: []string{"Type of item"},
			},
			{
				Name:        "itemId",
				Type:        "xs:string",
				Description: []string{"The UUID representing the item in storage"},
			},
		},
	},
	{
		ResourceID: "workitem",
		TypeID:     "com.ibm.team.repository.BooleanExtensionEntry",
		Fields: []Field{
			{
				Name:        "key",
				Type:        "xs:string",
				Description: []string{"Key of the custom attribute"},
			},
			{
				Name:        "value",
				Type:        "xs:boolean",
				Description: []string{"Value of the custom attribute"},
			},
		},
	},
	{
		ResourceID: "workitem",
		TypeID:     "com.ibm.team.repository.IntExtensionEntry",
		Fields: []Field{
			{
				Name:        "key",
				Type:        "xs:string",
				Description: []string{"Key of the custom attribute"},
			},
			{
				Name:        "value",
				Type:        "xs:integer",
				Description: []string{"Value of the custom attribute"},
			},
		},
	},
	{
		ResourceID: "workitem",
		TypeID:     "com.ibm.team.repository.LongExtensionEntry",
		Fields: []Field{
			{
				Name:        "key",
				Type:        "xs:string",
				Description: []string{"Key of the custom attribute"},
			},
			{
				Name:        "value",
				Type:        "xs:long",
				Description: []string{"Value of the custom attribute"},
			},
		},
	},
	{
		ResourceID: "workitem",
		TypeID:     "com.ibm.team.repository.StringExtensionEntry",
		Fields: []Field{
			{
				Name:        "key",
				Type:        "xs:string",
				Description: []string{"Key of the custom attribute"},
			},
			{
				Name:        "value",
				Type:        "xs:string",
				Description: []string{"Value of the custom attribute"},
			},
		},
	},
	{
		ResourceID: "workitem",
		TypeID:     "com.ibm.team.repository.MediumStringExtensionEntry",
		Fields: []Field{
			{
				Name:        "key",
				Type:        "xs:string",
				Description: []string{"Key of the custom attribute"},
			},
			{
				Name:        "value",
				Type:        "xs:string",
				Description: []string{"Value of the custom attribute"},
			},
		},
	},
	{
		ResourceID: "workitem",
		TypeID:     "com.ibm.team.repository.LargeStringExtensionEntry",
		Fields: []Field{
			{
				Name:        "key",
				Type:        "xs:string",
				Description: []string{"Key of the custom attribute"},
			},
			{
				Name:        "value",
				Type:        "xs:string",
				Description: []string{"Value of the custom attribute"},
			},
		},
	},
	{
		ResourceID: "workitem",
		TypeID:     "com.ibm.team.repository.TimestampExtensionEntry",
		Fields: []Field{
			{
				Name:        "key",
				Type:        "xs:string",
				Description: []string{"Key of the custom attribute"},
			},
			{
				Name:        "value",
				Type:        "xs:time",
				Description: []string{"Value of the custom attribute"},
			},
		},
	},
	{
		ResourceID: "workitem",
		TypeID:     "com.ibm.team.repository.BigDecimalExtensionEntry",
		Fields: []Field{
			{
				Name:        "key",
				Type:        "xs:string",
				Description: []string{"Key of the custom attribute"},
			},
			{
				Name:        "value",
				Type:        "xs:decimal",
				Description: []string{"Value of the custom attribute"},
			},
		},
	},
	{
		ResourceID: "workitem",
		TypeID:     "com.ibm.team.workitem.ItemExtensionEntry",
		Fields: []Field{
			{
				Name:        "key",
				Type:        "xs:string",
				Description: []string{"Key of the custom attribute"},
			},
			{
				Name:        "value",
				Type:        "com.ibm.team.repository.Item",
				Description: []string{"Value of the custom attribute"},
			},
		},
	},
	{
		ResourceID: "workitem",
		TypeID:     "com.ibm.team.workitem.MultiItemExtensionEntry",
		Fields: []Field{
			{
				Name:        "key",
				Type:        "xs:string",
				Description: []string{"Key of the custom attribute"},
			},
			{
				Name:        "value",
				Type:        "com.ibm.team.repository.Item, maxOccurs: unbounded",
				Description: []string{"Value of the custom attribute"},
			},
		},
	},
}
