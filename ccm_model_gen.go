package jazz

// Code generated! DO NOT EDIT

import (
	"reflect"
	"time"
)

func init() {
	ccmRegisterType(new(CCMProjectArea))
	ccmRegisterType(new(CCMTeamAreaHierarchyRecord))
	ccmRegisterType(new(CCMTeamArea))
	ccmRegisterType(new(CCMContributor))
	ccmRegisterType(new(CCMIteration))
	ccmRegisterType(new(CCMDevelopmentLine))
	ccmRegisterType(new(CCMAuditableLink))
	ccmRegisterType(new(CCMReference))
	ccmRegisterType(new(CCMReferenceType))
	ccmRegisterType(new(CCMReadAccess))
	ccmRegisterType(new(CCMRole))
	ccmRegisterType(new(CCMRoleAssignment))
	ccmRegisterType(new(CCMWorkspace))
	ccmRegisterType(new(CCMProperty))
	ccmRegisterType(new(CCMComponent))
	ccmRegisterType(new(CCMChangeSet))
	ccmRegisterType(new(CCMBuildDefinition))
	ccmRegisterType(new(CCMBuildResult))
	ccmRegisterType(new(CCMCompilationResult))
	ccmRegisterType(new(CCMUnitTestResult))
	ccmRegisterType(new(CCMUnitTestEvent))
	ccmRegisterType(new(CCMBuildEngine))
	ccmRegisterType(new(CCMWorkItem))
	ccmRegisterType(new(CCMComment))
	ccmRegisterType(new(CCMAttribute))
	ccmRegisterType(new(CCMApproval))
	ccmRegisterType(new(CCMApprovalDescriptor))
	ccmRegisterType(new(CCMState))
	ccmRegisterType(new(CCMResolution))
	ccmRegisterType(new(CCMWorkItemType))
	ccmRegisterType(new(CCMLiteral))
	ccmRegisterType(new(CCMCategory))
	ccmRegisterType(new(CCMDeliverable))
	ccmRegisterType(new(CCMExtensionEntry))
	ccmRegisterType(new(CCMTimeSheetEntry))
	ccmRegisterType(new(CCMItem))
	ccmRegisterType(new(CCMBooleanExtensionEntry))
	ccmRegisterType(new(CCMIntExtensionEntry))
	ccmRegisterType(new(CCMLongExtensionEntry))
	ccmRegisterType(new(CCMStringExtensionEntry))
	ccmRegisterType(new(CCMMediumStringExtensionEntry))
	ccmRegisterType(new(CCMLargeStringExtensionEntry))
	ccmRegisterType(new(CCMTimestampExtensionEntry))
	ccmRegisterType(new(CCMBigDecimalExtensionEntry))
	ccmRegisterType(new(CCMItemExtensionEntry))
	ccmRegisterType(new(CCMMultiItemExtensionEntry))
}

// CCMProjectArea (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#projectArea_type_com_ibm_team_pr)
// This element represents a Project Area.
type CCMProjectArea struct {
	CCMBaseObject

	// The human-readable name of the project area (e.g. "My Project")
	Name string `jazz:"name"`

	// A list of members of this project
	TeamMembers []*CCMMultiItemExtensionEntry `jazz:"teamMembers"`

	// A list of records reflecting the team area hierarchy for this project area
	TeamAreaHierarchy []*CCMMultiItemExtensionEntry `jazz:"teamAreaHierarchy"`

	// A list of development lines for this project area
	DevelopmentLines []*CCMMultiItemExtensionEntry `jazz:"developmentLines"`

	// The main development line for this project area
	ProjectDevelopmentLine *CCMMultiItemExtensionEntry `jazz:"projectDevelopmentLine"`

	// The roles defined in the project area
	Roles []*CCMMultiItemExtensionEntry `jazz:"roles"`

	// The role assignments defined in the project area
	RoleAssignments []*CCMMultiItemExtensionEntry `jazz:"roleAssignments"`

	// All the team areas contained in the project area
	AllTeamAreas []*CCMMultiItemExtensionEntry `jazz:"allTeamAreas"`
}

// CCMProjectAreaType contains the reflection type of CCMProjectArea
var goCCMProjectAreaType = reflect.TypeOf(CCMProjectArea{})

// Spec returns the specification object for CCMProjectArea
func (o *CCMProjectArea) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "projectArea",
		TypeID:     "com.ibm.team.process.ProjectArea",
		Type:       goCCMProjectAreaType,
	}
}

// Load CCMProjectArea object
func (o *CCMProjectArea) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMProjectArea object
func (o *CCMProjectArea) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.TeamMembers,
		o.TeamAreaHierarchy,
		o.DevelopmentLines,
		o.ProjectDevelopmentLine,
		o.Roles,
		o.RoleAssignments,
		o.AllTeamAreas,
	)
}

// CCMTeamAreaHierarchyRecord (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_process_TeamAreaHie)
// This element appears only inside a Project Area, and represents a piece of
// a team area hierarchy.
type CCMTeamAreaHierarchyRecord struct {
	CCMBaseObject

	// The parent team area
	Parent *CCMMultiItemExtensionEntry `jazz:"parent"`

	// The children team areas of the parent team area
	Children []*CCMMultiItemExtensionEntry `jazz:"children"`
}

// CCMTeamAreaHierarchyRecordType contains the reflection type of CCMTeamAreaHierarchyRecord
var goCCMTeamAreaHierarchyRecordType = reflect.TypeOf(CCMTeamAreaHierarchyRecord{})

// Spec returns the specification object for CCMTeamAreaHierarchyRecord
func (o *CCMTeamAreaHierarchyRecord) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.process.TeamAreaHierarchyRecord",
		Type:       goCCMTeamAreaHierarchyRecordType,
	}
}

// LoadAllFields of CCMTeamAreaHierarchyRecord object
func (o *CCMTeamAreaHierarchyRecord) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Parent,
		o.Children,
	)
}

// CCMTeamArea (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#teamArea_type_com_ibm_team_proce)
// This element represents a Team Area.
type CCMTeamArea struct {
	CCMBaseObject

	// The human-readable name of the project area (e.g. "My Team")
	Name string `jazz:"name"`

	// A fully-qualified team area name, slash-separated, including all parent
	// team areas (e.g. "/My Parent Team/My Team").
	QualifiedName string `jazz:"qualifiedName"`

	// A list of members of this team area
	TeamMembers []*CCMMultiItemExtensionEntry `jazz:"teamMembers"`

	// The project area containing this team area
	ProjectArea *CCMMultiItemExtensionEntry `jazz:"projectArea"`

	// The roles defined in the team area
	Roles []*CCMMultiItemExtensionEntry `jazz:"roles"`

	// The role assignments defined in the team area
	RoleAssignments []*CCMMultiItemExtensionEntry `jazz:"roleAssignments"`

	// The parent team area
	ParentTeamArea *CCMMultiItemExtensionEntry `jazz:"parentTeamArea"`
}

// CCMTeamAreaType contains the reflection type of CCMTeamArea
var goCCMTeamAreaType = reflect.TypeOf(CCMTeamArea{})

// Spec returns the specification object for CCMTeamArea
func (o *CCMTeamArea) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "teamArea",
		TypeID:     "com.ibm.team.process.TeamArea",
		Type:       goCCMTeamAreaType,
	}
}

// Load CCMTeamArea object
func (o *CCMTeamArea) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMTeamArea object
func (o *CCMTeamArea) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.TeamMembers,
		o.ProjectArea,
		o.Roles,
		o.RoleAssignments,
		o.ParentTeamArea,
	)
}

// CCMContributor (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#contributor)
// This element represents a Contributor (user).
type CCMContributor struct {
	CCMBaseObject

	// The human-readable name of the contributor (e.g. "James Moody")
	Name string `jazz:"name"`

	// The email address of the contributor
	EmailAddress string `jazz:"emailAddress"`

	// The userId of the contributor, unique in this application (e.g. "jmoody")
	UserId string `jazz:"userId"`
}

// CCMContributorType contains the reflection type of CCMContributor
var goCCMContributorType = reflect.TypeOf(CCMContributor{})

// Spec returns the specification object for CCMContributor
func (o *CCMContributor) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "contributor",
		TypeID:     "com.ibm.team.repository.Contributor",
		Type:       goCCMContributorType,
	}
}

// Load CCMContributor object
func (o *CCMContributor) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMContributor object
func (o *CCMContributor) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMIteration (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#iteration_type_com_ibm_team_proc)
// This element represents a single iteration (milestone, sprint).
type CCMIteration struct {
	CCMBaseObject

	// The human-readable name of this iteration (e.g. "M1")
	Name string `jazz:"name"`

	// The identifier of this iteration (e.g. "3.0M1")
	Id string `jazz:"id"`

	// The start date of this iteration
	StartDate *time.Time `jazz:"startDate"`

	// The end date of this iteration
	EndDate *time.Time `jazz:"endDate"`

	// The parent iteration of this iteration, if any
	Parent *CCMMultiItemExtensionEntry `jazz:"parent"`

	// The immediate child iterations of this iteration, if any
	Children []*CCMMultiItemExtensionEntry `jazz:"children"`

	// The development line in which this iteration appears
	DevelopmentLine *CCMMultiItemExtensionEntry `jazz:"developmentLine"`

	// Whether or not this iteration is marked as having deliverables associated
	// with it
	HasDeliverable bool `jazz:"hasDeliverable"`
}

// CCMIterationType contains the reflection type of CCMIteration
var goCCMIterationType = reflect.TypeOf(CCMIteration{})

// Spec returns the specification object for CCMIteration
func (o *CCMIteration) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "iteration",
		TypeID:     "com.ibm.team.process.Iteration",
		Type:       goCCMIterationType,
	}
}

// Load CCMIteration object
func (o *CCMIteration) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMIteration object
func (o *CCMIteration) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Parent,
		o.Children,
		o.DevelopmentLine,
	)
}

// CCMDevelopmentLine (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#developmentLine_type_com_ibm_tea)
// This element represents a development line.
type CCMDevelopmentLine struct {
	CCMBaseObject

	// The human-readable name of this development line (e.g. "Maintenance
	// Development")
	Name string `jazz:"name"`

	// The start date of this development line
	StartDate *time.Time `jazz:"startDate"`

	// The end date of this development line
	EndDate *time.Time `jazz:"endDate"`

	// The child iterations of this development line
	Iterations []*CCMMultiItemExtensionEntry `jazz:"iterations"`

	// The project area containing this development line
	ProjectArea *CCMMultiItemExtensionEntry `jazz:"projectArea"`

	// The iteration marked as current in this development line
	CurrentIteration *CCMMultiItemExtensionEntry `jazz:"currentIteration"`
}

// CCMDevelopmentLineType contains the reflection type of CCMDevelopmentLine
var goCCMDevelopmentLineType = reflect.TypeOf(CCMDevelopmentLine{})

// Spec returns the specification object for CCMDevelopmentLine
func (o *CCMDevelopmentLine) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "developmentLine",
		TypeID:     "com.ibm.team.process.DevelopmentLine",
		Type:       goCCMDevelopmentLineType,
	}
}

// Load CCMDevelopmentLine object
func (o *CCMDevelopmentLine) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMDevelopmentLine object
func (o *CCMDevelopmentLine) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Iterations,
		o.ProjectArea,
		o.CurrentIteration,
	)
}

// CCMAuditableLink (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#auditableLink)
// This element represents a link from one artifact to another. These links
// may be either within the same repository, or between one artifact in this
// repository and one external artifact. References (source and target) may be
// made either by uri (for any artifact) or by referencedItem (in the case of
// local artifacts).
type CCMAuditableLink struct {
	CCMBaseObject

	// The id of this link type (e.g. "com.ibm.team.workitem.parentChild"). This
	// describes the relationship represented by this link.
	Name string `jazz:"name"`

	// The source of the link
	SourceRef *CCMMultiItemExtensionEntry `jazz:"sourceRef"`

	// The target of the link
	TargetRef *CCMMultiItemExtensionEntry `jazz:"targetRef"`
}

// CCMAuditableLinkType contains the reflection type of CCMAuditableLink
var goCCMAuditableLinkType = reflect.TypeOf(CCMAuditableLink{})

// Spec returns the specification object for CCMAuditableLink
func (o *CCMAuditableLink) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "auditableLink",
		TypeID:     "",
		Type:       goCCMAuditableLinkType,
	}
}

// Load CCMAuditableLink object
func (o *CCMAuditableLink) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMAuditableLink object
func (o *CCMAuditableLink) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.SourceRef,
		o.TargetRef,
	)
}

// CCMReference (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_links_Reference)
// This element is always contained in an auditableLink, and represents either
// the source or target reference of a link. The reference may be either by
// uri (for any artifact) or by referencedItem (in the case of local
// artifacts). Which one can be determined by the referenceType field.
type CCMReference struct {
	CCMBaseObject

	// A human-readable comment about the reference. In some cases the comment may
	// suffice rather than fetching the content on the other end of the link. For
	// example, a reference pointing to a work item may contain the id and summary
	// of the work item ("12345: Summary of my work item").
	Comment string `jazz:"comment"`

	// This element indicates whether the reference is by uri or by itemId.
	ReferenceType *CCMMultiItemExtensionEntry `jazz:"referenceType"`

	// The URI of the element referenced. This is only valid if this Reference is
	// a URI reference.
	Uri string `jazz:"uri"`

	// The referenced item. This is only valid if this Reference is an Item
	// reference.
	ReferencedItem *CCMMultiItemExtensionEntry `jazz:"referencedItem"`

	// Get the extra information associated with the reference. May be null.
	ExtraInfo string `jazz:"extraInfo"`

	// Internal.
	ContentType string `jazz:"contentType"`
}

// CCMReferenceType contains the reflection type of CCMReference
var goCCMReferenceType = reflect.TypeOf(CCMReference{})

// Spec returns the specification object for CCMReference
func (o *CCMReference) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.links.Reference",
		Type:       goCCMReferenceType,
	}
}

// LoadAllFields of CCMReference object
func (o *CCMReference) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.ReferenceType,
		o.ReferencedItem,
	)
}

// CCMReferenceType (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_links_ReferenceType)
// This element represents a reference type, indicating whether a reference is
// by URI or itemID.
type CCMReferenceType struct {
	CCMBaseObject

	// Either "ITEM_REFERENCE" or "URI_REFERENCE"
	Literal string `jazz:"literal"`

	// Either 0 (for ITEM_REFERENCE) or 2 (for URI_REFERENCE). Use literal
	// instead.
	Value int `jazz:"value"`
}

// CCMReferenceTypeType contains the reflection type of CCMReferenceType
var goCCMReferenceTypeType = reflect.TypeOf(CCMReferenceType{})

// Spec returns the specification object for CCMReferenceType
func (o *CCMReferenceType) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.links.ReferenceType",
		Type:       goCCMReferenceTypeType,
	}
}

// LoadAllFields of CCMReferenceType object
func (o *CCMReferenceType) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMReadAccess (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#readAccess)
// The readAccess element represents a mapping of contributors to project
// areas that each contributor has permissions to read.
type CCMReadAccess struct {
	CCMBaseObject

	// The itemId of the Contributor
	ContributorItemId string `jazz:"contributorItemId"`

	// The itemID of the context object associated with the contributor (i.e. the
	// project area)
	ContributorContextId string `jazz:"contributorContextId"`
}

// CCMReadAccessType contains the reflection type of CCMReadAccess
var goCCMReadAccessType = reflect.TypeOf(CCMReadAccess{})

// Spec returns the specification object for CCMReadAccess
func (o *CCMReadAccess) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "readAccess",
		TypeID:     "",
		Type:       goCCMReadAccessType,
	}
}

// Load CCMReadAccess object
func (o *CCMReadAccess) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMReadAccess object
func (o *CCMReadAccess) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMRole (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_process_Role)
type CCMRole struct {
	CCMBaseObject

	// The role Id
	Id string `jazz:"id"`

	// The role name
	Name string `jazz:"name"`

	// The role description
	Description string `jazz:"description"`
}

// CCMRoleType contains the reflection type of CCMRole
var goCCMRoleType = reflect.TypeOf(CCMRole{})

// Spec returns the specification object for CCMRole
func (o *CCMRole) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.process.Role",
		Type:       goCCMRoleType,
	}
}

// LoadAllFields of CCMRole object
func (o *CCMRole) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMRoleAssignment (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_process_RoleAssignm)
type CCMRoleAssignment struct {
	CCMBaseObject

	// The contributor with assigned roles
	Contributor *CCMMultiItemExtensionEntry `jazz:"contributor"`

	// The roles assigned to the contributor
	ContributorRoles []*CCMMultiItemExtensionEntry `jazz:"contributorRoles"`
}

// CCMRoleAssignmentType contains the reflection type of CCMRoleAssignment
var goCCMRoleAssignmentType = reflect.TypeOf(CCMRoleAssignment{})

// Spec returns the specification object for CCMRoleAssignment
func (o *CCMRoleAssignment) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.process.RoleAssignment",
		Type:       goCCMRoleAssignmentType,
	}
}

// LoadAllFields of CCMRoleAssignment object
func (o *CCMRoleAssignment) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Contributor,
		o.ContributorRoles,
	)
}

// CCMWorkspace (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#workspace_type_com_ibm_team_scm)
// This element represents an SCM Workspace or Stream
type CCMWorkspace struct {
	CCMBaseObject

	// The name of the workspace or stream
	Name string `jazz:"name"`

	// True if this is a stream, false if this is a workspace
	Stream bool `jazz:"stream"`

	// A description of the workspace or stream
	Description string `jazz:"description"`

	// Whether or not ETL data collection is configured for this stream
	CollectData bool `jazz:"collectData"`

	// A collection of key/value properties associated with the workspace or
	// stream
	Properties []*CCMMultiItemExtensionEntry `jazz:"properties"`

	// The owner of the workspace or stream
	Contributor *CCMMultiItemExtensionEntry `jazz:"contributor"`
}

// CCMWorkspaceType contains the reflection type of CCMWorkspace
var goCCMWorkspaceType = reflect.TypeOf(CCMWorkspace{})

// Spec returns the specification object for CCMWorkspace
func (o *CCMWorkspace) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "scm",
		ElementID:  "workspace",
		TypeID:     "com.ibm.team.scm.Workspace",
		Type:       goCCMWorkspaceType,
	}
}

// Load CCMWorkspace object
func (o *CCMWorkspace) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMWorkspace object
func (o *CCMWorkspace) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Properties,
		o.Contributor,
	)
}

// CCMProperty (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_scm_Property)
// This element only occurs in a workspace, and represents a property of a
// Workspace or Stream
type CCMProperty struct {
	CCMBaseObject

	// The property key
	Key string `jazz:"key"`
}

// CCMPropertyType contains the reflection type of CCMProperty
var goCCMPropertyType = reflect.TypeOf(CCMProperty{})

// Spec returns the specification object for CCMProperty
func (o *CCMProperty) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "scm",
		ElementID:  "",
		TypeID:     "com.ibm.team.scm.Property",
		Type:       goCCMPropertyType,
	}
}

// LoadAllFields of CCMProperty object
func (o *CCMProperty) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMComponent (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#component_type_com_ibm_team_scm)
// This element represents an SCM Component
type CCMComponent struct {
	CCMBaseObject

	// The name of the component
	Name string `jazz:"name"`
}

// CCMComponentType contains the reflection type of CCMComponent
var goCCMComponentType = reflect.TypeOf(CCMComponent{})

// Spec returns the specification object for CCMComponent
func (o *CCMComponent) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "scm",
		ElementID:  "component",
		TypeID:     "com.ibm.team.scm.Component",
		Type:       goCCMComponentType,
	}
}

// Load CCMComponent object
func (o *CCMComponent) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMComponent object
func (o *CCMComponent) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMChangeSet (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#changeSet_type_com_ibm_team_scm)
// This element represents an SCM Change Set
type CCMChangeSet struct {
	CCMBaseObject

	// The comment on the change set
	Comment string `jazz:"comment"`

	// The owner of the change set
	Owner *CCMMultiItemExtensionEntry `jazz:"owner"`
}

// CCMChangeSetType contains the reflection type of CCMChangeSet
var goCCMChangeSetType = reflect.TypeOf(CCMChangeSet{})

// Spec returns the specification object for CCMChangeSet
func (o *CCMChangeSet) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "scm",
		ElementID:  "changeSet",
		TypeID:     "com.ibm.team.scm.ChangeSet",
		Type:       goCCMChangeSetType,
	}
}

// Load CCMChangeSet object
func (o *CCMChangeSet) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMChangeSet object
func (o *CCMChangeSet) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Owner,
	)
}

// CCMBuildDefinition (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#buildDefinition_type_com_ibm_tea)
// This element represents a Build Definition.
type CCMBuildDefinition struct {
	CCMBaseObject

	// The id of the build definition
	Id string `jazz:"id"`

	// The description of the build definition
	Description string `jazz:"description"`

	// The project area containing the build definition
	ProjectArea *CCMMultiItemExtensionEntry `jazz:"projectArea"`

	// The team area containing the build definition
	TeamArea *CCMMultiItemExtensionEntry `jazz:"teamArea"`
}

// CCMBuildDefinitionType contains the reflection type of CCMBuildDefinition
var goCCMBuildDefinitionType = reflect.TypeOf(CCMBuildDefinition{})

// Spec returns the specification object for CCMBuildDefinition
func (o *CCMBuildDefinition) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "build",
		ElementID:  "buildDefinition",
		TypeID:     "com.ibm.team.build.BuildDefinition",
		Type:       goCCMBuildDefinitionType,
	}
}

// Load CCMBuildDefinition object
func (o *CCMBuildDefinition) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMBuildDefinition object
func (o *CCMBuildDefinition) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.ProjectArea,
		o.TeamArea,
	)
}

// CCMBuildResult (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#buildResult_type_com_ibm_team_bu)
// This element represents a Build Result.
type CCMBuildResult struct {
	CCMBaseObject

	// James: To Do
	BuildStatus string `jazz:"buildStatus"`

	// James: To Do
	BuildState string `jazz:"buildState"`

	// The label for the build
	Label string `jazz:"label"`

	// How long the build took, in milliseconds
	TimeTaken int64 `jazz:"timeTaken"`

	// Whether this was a personal build or not
	PersonalBuild bool `jazz:"personalBuild"`

	// The start time of the build
	StartTime *time.Time `jazz:"startTime"`

	// How long the build waited in the queue, in milliseconds
	TimeWaiting int64 `jazz:"timeWaiting"`

	// Which build definition this build was for
	BuildDefinition *CCMMultiItemExtensionEntry `jazz:"buildDefinition"`

	// The contributor who requested the build
	Creator *CCMMultiItemExtensionEntry `jazz:"creator"`

	// The engine the build ran on
	BuildEngine *CCMMultiItemExtensionEntry `jazz:"buildEngine"`

	// Code compilation results
	CompilationResults []*CCMMultiItemExtensionEntry `jazz:"compilationResults"`

	// Unit test results
	UnitTestResults []*CCMMultiItemExtensionEntry `jazz:"unitTestResults"`

	// Unit test changes from the previous build
	UnitTestEvents []*CCMMultiItemExtensionEntry `jazz:"unitTestEvents"`
}

// CCMBuildResultType contains the reflection type of CCMBuildResult
var goCCMBuildResultType = reflect.TypeOf(CCMBuildResult{})

// Spec returns the specification object for CCMBuildResult
func (o *CCMBuildResult) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "build",
		ElementID:  "buildResult",
		TypeID:     "com.ibm.team.build.BuildResult",
		Type:       goCCMBuildResultType,
	}
}

// Load CCMBuildResult object
func (o *CCMBuildResult) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMBuildResult object
func (o *CCMBuildResult) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.BuildDefinition,
		o.Creator,
		o.BuildEngine,
		o.CompilationResults,
		o.UnitTestResults,
		o.UnitTestEvents,
	)
}

// CCMCompilationResult (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_build_CompilationRe)
// This element only occurs in a buildResult. The number of errors and
// warnings for a particular component in the containing build result
type CCMCompilationResult struct {
	CCMBaseObject

	// The component for which the errors and warnings are being reported
	Component string `jazz:"component"`

	// The number of compilation errors for the component in the containing build
	// result
	Errors int64 `jazz:"errors"`

	// The umber of compilation warnings for the component in the containing build
	// result
	Warnings int64 `jazz:"warnings"`
}

// CCMCompilationResultType contains the reflection type of CCMCompilationResult
var goCCMCompilationResultType = reflect.TypeOf(CCMCompilationResult{})

// Spec returns the specification object for CCMCompilationResult
func (o *CCMCompilationResult) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "build",
		ElementID:  "",
		TypeID:     "com.ibm.team.build.CompilationResult",
		Type:       goCCMCompilationResultType,
	}
}

// LoadAllFields of CCMCompilationResult object
func (o *CCMCompilationResult) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMUnitTestResult (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_build_UnitTestResul)
// This element only occurs in a buildResult. The number of unit tests run,
// along with number of failures and errors, for a particular component in the
// containing build result
type CCMUnitTestResult struct {
	CCMBaseObject

	// The component for which the tests, errors and failures are being reported
	Component string `jazz:"component"`

	// The number of unit tests run for the component in the containing build
	// result
	Tests int64 `jazz:"tests"`

	// The number of unit test failures for the component in the containing build
	// result
	Failures int64 `jazz:"failures"`

	// The number of unit test errors for the component in the containing build
	// result
	Errors int64 `jazz:"errors"`
}

// CCMUnitTestResultType contains the reflection type of CCMUnitTestResult
var goCCMUnitTestResultType = reflect.TypeOf(CCMUnitTestResult{})

// Spec returns the specification object for CCMUnitTestResult
func (o *CCMUnitTestResult) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "build",
		ElementID:  "",
		TypeID:     "com.ibm.team.build.UnitTestResult",
		Type:       goCCMUnitTestResultType,
	}
}

// LoadAllFields of CCMUnitTestResult object
func (o *CCMUnitTestResult) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMUnitTestEvent (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_build_UnitTestEvent)
// This element only occurs in a buildResult. It represents a single unit test
// execution, along with a pass, fail or regression label
type CCMUnitTestEvent struct {
	CCMBaseObject

	// The component for which the test and event is being reported
	Component string `jazz:"component"`

	// The name of the unit test run
	Test string `jazz:"test"`

	// Indication of test passing, failing or regressing. James: To do, provide
	// the literals here.
	Event string `jazz:"event"`
}

// CCMUnitTestEventType contains the reflection type of CCMUnitTestEvent
var goCCMUnitTestEventType = reflect.TypeOf(CCMUnitTestEvent{})

// Spec returns the specification object for CCMUnitTestEvent
func (o *CCMUnitTestEvent) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "build",
		ElementID:  "",
		TypeID:     "com.ibm.team.build.UnitTestEvent",
		Type:       goCCMUnitTestEventType,
	}
}

// LoadAllFields of CCMUnitTestEvent object
func (o *CCMUnitTestEvent) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMBuildEngine (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#buildEngine_type_com_ibm_team_bu)
// This element represents a build engine.
type CCMBuildEngine struct {
	CCMBaseObject

	// The id of this build engine
	Id string `jazz:"id"`
}

// CCMBuildEngineType contains the reflection type of CCMBuildEngine
var goCCMBuildEngineType = reflect.TypeOf(CCMBuildEngine{})

// Spec returns the specification object for CCMBuildEngine
func (o *CCMBuildEngine) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "build",
		ElementID:  "buildEngine",
		TypeID:     "com.ibm.team.build.BuildEngine",
		Type:       goCCMBuildEngineType,
	}
}

// Load CCMBuildEngine object
func (o *CCMBuildEngine) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMBuildEngine object
func (o *CCMBuildEngine) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMWorkItem (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#workItem_type_com_ibm_team_worki)
// This element represents a Work Item.
type CCMWorkItem struct {
	CCMBaseObject

	// The system-generated id number for the work item (e.g. "123")
	Id int `jazz:"id"`

	// The date and time when the work item was resolved, or null if the work item
	// has not been resolved
	ResolutionDate *time.Time `jazz:"resolutionDate"`

	// The one-line summary (or title) of the work item
	Summary string `jazz:"summary"`

	// The date and time when the work item was created
	CreationDate *time.Time `jazz:"creationDate"`

	// The date and time when the work item is scheduled for completion, or null
	// if no due date has been specified
	DueDate *time.Time `jazz:"dueDate"`

	// The multi-line description of the work item
	Description string `jazz:"description"`

	// James: To Do
	WorkflowSurrogate string `jazz:"workflowSurrogate"`

	// The tags attached to the work item. In the case of multiple tags, this
	// single string contains a comma-separated list of tags
	Tags string `jazz:"tags"`

	// The estimate specified for the work item, indicated the estimated time to
	// complete the work item. In the UI, this is called "Estimate" rather than
	// duration.
	Duration int64 `jazz:"duration"`

	// How much time has actually been spent so far on the work item
	TimeSpent int64 `jazz:"timeSpent"`

	// The corrected estimate for the work item, in the case that the user has
	// corrected the estimate
	CorrectedEstimate int64 `jazz:"correctedEstimate"`

	// The day on which the work item was last modified
	DayModified *time.Time `jazz:"dayModified"`

	// The contributor who created the work item
	Creator *CCMMultiItemExtensionEntry `jazz:"creator"`

	// The contributor who owns the work item
	Owner *CCMMultiItemExtensionEntry `jazz:"owner"`

	// The category to which the work item is assigned. In the UI, this is called
	// "Filed Against".
	Category *CCMMultiItemExtensionEntry `jazz:"category"`

	// A collection of zero or more comments appended to the work item
	Comments []*CCMMultiItemExtensionEntry `jazz:"comments"`

	// A collection of zero or more "custom attributes" attached to the work item.
	// These are user-defined attributes (as opposed to the built-in attributes
	// elsewhere in this list).
	CustomAttributes []*CCMMultiItemExtensionEntry `jazz:"customAttributes"`

	// A collection of zero or more Contributors who are subscribed to the work
	// item
	Subscriptions []*CCMMultiItemExtensionEntry `jazz:"subscriptions"`

	// The project area to which the work item belongs
	ProjectArea *CCMMultiItemExtensionEntry `jazz:"projectArea"`

	// The Contributor who resolved the work item, or null if the work item has
	// not been resolved
	Resolver *CCMMultiItemExtensionEntry `jazz:"resolver"`

	// A collection of zero or more Approvals attached to the work item
	Approvals []*CCMMultiItemExtensionEntry `jazz:"approvals"`

	// A collection of zero or more Approval Descriptors attached to the work item
	ApprovalDescriptors []*CCMMultiItemExtensionEntry `jazz:"approvalDescriptors"`

	// The iteration that the work item is "Planned For"
	Target *CCMMultiItemExtensionEntry `jazz:"target"`

	// The deliverable that the work item is "Found In"
	FoundIn *CCMMultiItemExtensionEntry `jazz:"foundIn"`

	// A collection of zero or more WorkItem elements, representing the entire
	// history of the work item. Each state the work item has ever been in is
	// reflected in this history list.
	ItemHistory []*CCMMultiItemExtensionEntry `jazz:"itemHistory"`

	// The team area to which the work item belongs
	TeamArea *CCMMultiItemExtensionEntry `jazz:"teamArea"`

	// The state of the work item (e.g. "Resolved", "In Progress", "New"). The
	// states are user-defined as part of the project area process.
	State *CCMMultiItemExtensionEntry `jazz:"state"`

	// The resolution of the work item (e.g. "Duplicate", "Invalid", "Fixed"). The
	// resolutions are user-defined as part of the project area process.
	Resolution *CCMMultiItemExtensionEntry `jazz:"resolution"`

	// The type of the work item (e.g. "Defect", "Task", "Story"). The work item
	// types are user-defined as part of the project area process.
	Type *CCMMultiItemExtensionEntry `jazz:"type"`

	// The severity of the work item (e.g. "Critical", "Normal", "Blocker"). The
	// work item severities are user-defined as part of the project area process.
	Severity *CCMMultiItemExtensionEntry `jazz:"severity"`

	// The priority of the work item (e.g. "High", "Medium", "Low"). The work item
	// priorities are user-defined as part of the project area process.
	Priority *CCMMultiItemExtensionEntry `jazz:"priority"`

	// The parent work item of this work item, if one exists
	Parent *CCMMultiItemExtensionEntry `jazz:"parent"`

	// A collection of zero or more child work items
	Children []*CCMMultiItemExtensionEntry `jazz:"children"`

	// A collection of zero or more work items which this work item blocks
	Blocks []*CCMMultiItemExtensionEntry `jazz:"blocks"`

	// A collection of zero or more work items which block this work item
	DependsOn []*CCMMultiItemExtensionEntry `jazz:"dependsOn"`

	// A collection of zero or more work items which are closed as duplicates of
	// this work item
	DuplicatedBy []*CCMMultiItemExtensionEntry `jazz:"duplicatedBy"`

	// A collection of zero or more work items which this work item is a duplicate
	// of
	DuplicateOf []*CCMMultiItemExtensionEntry `jazz:"duplicateOf"`

	// A collection of zero of more work items which this work item is related to
	Related []*CCMMultiItemExtensionEntry `jazz:"related"`

	// A collection of zero or more items linked to the work item as custom
	// attributes
	ItemExtensions []*CCMMultiItemExtensionEntry `jazz:"itemExtensions"`

	// A collection of zero or more lists of items linked to the work item as
	// custom attributes
	MultiItemExtensions []*CCMMultiItemExtensionEntry `jazz:"multiItemExtensions"`

	// A collection of zero or more custom attributes of type medium string
	MediumStringExtensions []*CCMMultiItemExtensionEntry `jazz:"mediumStringExtensions"`

	// A collection of zero or more custom attributes of type boolean
	BooleanExtensions []*CCMMultiItemExtensionEntry `jazz:"booleanExtensions"`

	// A collection of zero or more custom attributes of type timestamp
	TimestampExtensions []*CCMMultiItemExtensionEntry `jazz:"timestampExtensions"`

	// A collection of zero or more custom attributes of type long
	LongExtensions []*CCMMultiItemExtensionEntry `jazz:"longExtensions"`

	// A collection of zero or more custom attributes of type integer
	IntExtensions []*CCMMultiItemExtensionEntry `jazz:"intExtensions"`

	// A collection of zero or more custom attributes of type big decimal
	BigDecimalExtensions []*CCMMultiItemExtensionEntry `jazz:"bigDecimalExtensions"`

	// A collection of zero or more custom attributes of type large string
	LargeStringExtensions []*CCMMultiItemExtensionEntry `jazz:"largeStringExtensions"`

	// A collection of zero or more custom attributes of type string
	StringExtensions []*CCMMultiItemExtensionEntry `jazz:"stringExtensions"`

	// A collection of zero or more custom attributes of all types
	AllExtensions []*CCMMultiItemExtensionEntry `jazz:"allExtensions"`

	// A collection of zero or more timesheet entries linked to the work item
	TimeSheetEntries []*CCMMultiItemExtensionEntry `jazz:"timeSheetEntries"`

	// The work item's planned start date as specified in the plan.
	PlannedStartDate *time.Time `jazz:"plannedStartDate"`

	// The work item's planned end date as specified in the plan.
	PlannedEndDate *time.Time `jazz:"plannedEndDate"`
}

// CCMWorkItemType contains the reflection type of CCMWorkItem
var goCCMWorkItemType = reflect.TypeOf(CCMWorkItem{})

// Spec returns the specification object for CCMWorkItem
func (o *CCMWorkItem) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "workItem",
		TypeID:     "com.ibm.team.workitem.WorkItem",
		Type:       goCCMWorkItemType,
	}
}

// Load CCMWorkItem object
func (o *CCMWorkItem) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMWorkItem object
func (o *CCMWorkItem) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Creator,
		o.Owner,
		o.Category,
		o.Comments,
		o.CustomAttributes,
		o.Subscriptions,
		o.ProjectArea,
		o.Resolver,
		o.Approvals,
		o.ApprovalDescriptors,
		o.Target,
		o.FoundIn,
		o.ItemHistory,
		o.TeamArea,
		o.State,
		o.Resolution,
		o.Type,
		o.Severity,
		o.Priority,
		o.Parent,
		o.Children,
		o.Blocks,
		o.DependsOn,
		o.DuplicatedBy,
		o.DuplicateOf,
		o.Related,
		o.ItemExtensions,
		o.MultiItemExtensions,
		o.MediumStringExtensions,
		o.BooleanExtensions,
		o.TimestampExtensions,
		o.LongExtensions,
		o.IntExtensions,
		o.BigDecimalExtensions,
		o.LargeStringExtensions,
		o.StringExtensions,
		o.AllExtensions,
		o.TimeSheetEntries,
	)
}

// CCMComment (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_Comment)
// This element represents a single work item comment.
type CCMComment struct {
	CCMBaseObject

	// The date/time that the comment was saved in the work item
	CreationDate *time.Time `jazz:"creationDate"`

	// The string content of the comment
	Content string `jazz:"content"`

	// Whether or not the comment has been edited
	Edited bool `jazz:"edited"`

	// The contributor who created the comment
	Creator *CCMMultiItemExtensionEntry `jazz:"creator"`
}

// CCMCommentType contains the reflection type of CCMComment
var goCCMCommentType = reflect.TypeOf(CCMComment{})

// Spec returns the specification object for CCMComment
func (o *CCMComment) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.Comment",
		Type:       goCCMCommentType,
	}
}

// LoadAllFields of CCMComment object
func (o *CCMComment) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Creator,
	)
}

// CCMAttribute (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_Attribute)
// This element represents information about a custom attribute declaration.
// Custom attribute declarations are process-specific.
type CCMAttribute struct {
	CCMBaseObject

	// An identifier for the custom attribute, unique within a project area
	Identifier string `jazz:"identifier"`

	// The data type of the attribute value
	AttributeType string `jazz:"attributeType"`

	// Whether or not the attribute is built-in
	BuiltIn bool `jazz:"builtIn"`

	// The project in which the attribute is defined
	ProjectArea *CCMMultiItemExtensionEntry `jazz:"projectArea"`
}

// CCMAttributeType contains the reflection type of CCMAttribute
var goCCMAttributeType = reflect.TypeOf(CCMAttribute{})

// Spec returns the specification object for CCMAttribute
func (o *CCMAttribute) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.Attribute",
		Type:       goCCMAttributeType,
	}
}

// LoadAllFields of CCMAttribute object
func (o *CCMAttribute) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.ProjectArea,
	)
}

// CCMApproval (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_Approval)
// This element represents an approval from a single contributor with a
// particular state.
type CCMApproval struct {
	CCMBaseObject

	// The state of the approval
	StateIdentifier string `jazz:"stateIdentifier"`

	// The date the state was assigned
	StateDate *time.Time `jazz:"stateDate"`

	// The name of the state
	StateName string `jazz:"stateName"`

	// The contributor who is asked for approval
	Approver *CCMMultiItemExtensionEntry `jazz:"approver"`
}

// CCMApprovalType contains the reflection type of CCMApproval
var goCCMApprovalType = reflect.TypeOf(CCMApproval{})

// Spec returns the specification object for CCMApproval
func (o *CCMApproval) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.Approval",
		Type:       goCCMApprovalType,
	}
}

// LoadAllFields of CCMApproval object
func (o *CCMApproval) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Approver,
	)
}

// CCMApprovalDescriptor (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_ApprovalDe)
// This element represents an approval descriptor aggregates approvals from
// contributors.
type CCMApprovalDescriptor struct {
	CCMBaseObject

	// An identifier for this approval
	Id int `jazz:"id"`

	// The type of approval, used to distinguish Approvals, Reviews,
	// Verifications, or other types of approvals
	TypeIdentifier string `jazz:"typeIdentifier"`

	// The name of the type of approval
	TypeName string `jazz:"typeName"`

	// The display name for this approval
	Name string `jazz:"name"`

	// The cumulative state of all the approvals for this approval descriptor
	CumulativeStateIdentifier string `jazz:"cumulativeStateIdentifier"`

	// The name of the cumulative state
	CumulativeStateName string `jazz:"cumulativeStateName"`

	// The date this approval is due
	DueDate *time.Time `jazz:"dueDate"`

	// A collection of zero of more approvals aggregated by the approval
	// descriptor
	Approvals []*CCMMultiItemExtensionEntry `jazz:"approvals"`
}

// CCMApprovalDescriptorType contains the reflection type of CCMApprovalDescriptor
var goCCMApprovalDescriptorType = reflect.TypeOf(CCMApprovalDescriptor{})

// Spec returns the specification object for CCMApprovalDescriptor
func (o *CCMApprovalDescriptor) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.ApprovalDescriptor",
		Type:       goCCMApprovalDescriptorType,
	}
}

// LoadAllFields of CCMApprovalDescriptor object
func (o *CCMApprovalDescriptor) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Approvals,
	)
}

// CCMState (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_State)
// This element represents the state of a work item. States are defined by the
// user in the process specification for a project area.
type CCMState struct {
	CCMBaseObject

	// The id of the state (e.g. "com.ibm.team.workitem.defect.inProgress"),
	// unique in a repository.
	Id string `jazz:"id"`

	// The name of the state (e.g. "In Progress"). Not necessarily unique.
	Name string `jazz:"name"`

	// The "State Group" of this state. A state group is a process-independent
	// grouping of states, which is useful for creating reports which are not
	// dependent on a particular process but still need to know, for example,
	// whether work items are open or closed. Every state belongs to one of the
	// following state groups: "OPEN_STATES", "CLOSED_STATES",
	// "IN_PROGRESS_STATES".
	Group string `jazz:"group"`
}

// CCMStateType contains the reflection type of CCMState
var goCCMStateType = reflect.TypeOf(CCMState{})

// Spec returns the specification object for CCMState
func (o *CCMState) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.State",
		Type:       goCCMStateType,
	}
}

// LoadAllFields of CCMState object
func (o *CCMState) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMResolution (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_Resolution)
// This element represents the resolution of a work item. This indicates how
// or why a work item was resolved; for example, "Fixed", "Invalid", "Won't
// Fix". Resolutions are process-dependent.
type CCMResolution struct {
	CCMBaseObject

	// The id of the resolution (e.g. "com.ibm.team.workitem.defect.fixed"),
	// unique in a repository.
	Id string `jazz:"id"`

	// The name of the resolution (e.g. "Fixed"). Not necessarily unique.
	Name string `jazz:"name"`
}

// CCMResolutionType contains the reflection type of CCMResolution
var goCCMResolutionType = reflect.TypeOf(CCMResolution{})

// Spec returns the specification object for CCMResolution
func (o *CCMResolution) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.Resolution",
		Type:       goCCMResolutionType,
	}
}

// LoadAllFields of CCMResolution object
func (o *CCMResolution) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMWorkItemType (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_WorkItemTy)
// This element represents the type of a work item. Work item types are
// process-dependent.
type CCMWorkItemType struct {
	CCMBaseObject

	// The id of the type (e.g. "com.ibm.team.workitem.defect"), unique in a
	// repository.
	Id string `jazz:"id"`

	// The name of the type (e.g. "Defect"). Not necessarily unique.
	Name string `jazz:"name"`
}

// CCMWorkItemTypeType contains the reflection type of CCMWorkItemType
var goCCMWorkItemTypeType = reflect.TypeOf(CCMWorkItemType{})

// Spec returns the specification object for CCMWorkItemType
func (o *CCMWorkItemType) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.WorkItemType",
		Type:       goCCMWorkItemTypeType,
	}
}

// LoadAllFields of CCMWorkItemType object
func (o *CCMWorkItemType) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMLiteral (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_Literal)
// This element represents a user-defined literal value, used for priority and
// severity in a work item. Work item severities and priorities are
// process-dependent.
type CCMLiteral struct {
	CCMBaseObject

	// The id of the literal (e.g. "com.ibm.team.workitem.blocking"), unique in a
	// repository.
	Id string `jazz:"id"`

	// The name of the literal (e.g. "Blocking"). Not necessarily unique.
	Name string `jazz:"name"`
}

// CCMLiteralType contains the reflection type of CCMLiteral
var goCCMLiteralType = reflect.TypeOf(CCMLiteral{})

// Spec returns the specification object for CCMLiteral
func (o *CCMLiteral) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.Literal",
		Type:       goCCMLiteralType,
	}
}

// LoadAllFields of CCMLiteral object
func (o *CCMLiteral) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMCategory (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#category_type_com_ibm_team_worki)
// This element represents a work item Category. Work item categories are
// process-dependent.
type CCMCategory struct {
	CCMBaseObject

	// The id of the category, unique in a repository.
	Id string `jazz:"id"`

	// The name of the category (e.g. "Reports"). Not necessarily unique.
	Name string `jazz:"name"`

	// A textual description of the category.
	Description string `jazz:"description"`

	// The slash-separated qualified name of the category, indicating its
	// containment hierarchy (e.g. "/RTC Development/Reports").
	QualifiedName string `jazz:"qualifiedName"`
}

// CCMCategoryType contains the reflection type of CCMCategory
var goCCMCategoryType = reflect.TypeOf(CCMCategory{})

// Spec returns the specification object for CCMCategory
func (o *CCMCategory) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "category",
		TypeID:     "com.ibm.team.workitem.Category",
		Type:       goCCMCategoryType,
	}
}

// Load CCMCategory object
func (o *CCMCategory) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMCategory object
func (o *CCMCategory) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// CCMDeliverable (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#deliverable_type_com_ibm_team_wo)
// This element represents a deliverable, often used in Work Items to identify
// in which deliverable a work item was found ("Found In"). Deliverables are
// process-dependent.
type CCMDeliverable struct {
	CCMBaseObject

	// The name of the deliverable (e.g. "RTC 3.0")
	Name string `jazz:"name"`

	// A textual description of the deliverable
	Description string `jazz:"description"`

	// The creation date of the deliverable
	CreationDate *time.Time `jazz:"creationDate"`

	// The project area associated with the deliverable
	ProjectArea *CCMMultiItemExtensionEntry `jazz:"projectArea"`

	// An optional link to a repository item associated with the deliverable. This
	// field should be treated as internal.
	Artifact *CCMMultiItemExtensionEntry `jazz:"artifact"`
}

// CCMDeliverableType contains the reflection type of CCMDeliverable
var goCCMDeliverableType = reflect.TypeOf(CCMDeliverable{})

// Spec returns the specification object for CCMDeliverable
func (o *CCMDeliverable) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "deliverable",
		TypeID:     "com.ibm.team.workitem.Deliverable",
		Type:       goCCMDeliverableType,
	}
}

// Load CCMDeliverable object
func (o *CCMDeliverable) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMDeliverable object
func (o *CCMDeliverable) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.ProjectArea,
		o.Artifact,
	)
}

// CCMExtensionEntry (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#allExtensions_type_com_ibm_team)
// This element represents the value of a custom attribute.
type CCMExtensionEntry struct {
	CCMBaseObject

	// The name of the custom attribute
	Key string `jazz:"key"`

	// The type of the custom attribute (e.g. timestampValue, itemValue)
	Type string `jazz:"type"`

	// Boolean value if the type of the custom attribute is booleanValue, else
	// null
	BooleanValue bool `jazz:"booleanValue"`

	// Integer value if the type of the custom attribute is integerValue, else
	// null
	IntegerValue int `jazz:"integerValue"`

	// Long value if the type of the custom attribute is longValue, else null
	LongValue int64 `jazz:"longValue"`

	// Double value if the type of the custom attribute is doubleValue, else 0.0
	DoubleValue float64 `jazz:"doubleValue"`

	// String value if the type of the custom attribute is smallStringValue, else
	// null
	SmallStringValue string `jazz:"smallStringValue"`

	// String value if the type of the custom attribute is mediumStringValue, else
	// null
	MediumStringValue string `jazz:"mediumStringValue"`

	// String value if the type of the custom attribute is largeStringValue, else
	// null
	LargeStringValue string `jazz:"largeStringValue"`

	// Timestamp value if the type of the custom attribute is timestampValue, else
	// null
	TimestampValue *time.Time `jazz:"timestampValue"`

	// Decimal value if the type of the custom attribute is decimalValue, else
	// null
	DecimalValue float64 `jazz:"decimalValue"`

	// The information of the Item assigned as the value of the custom attribute
	// if the type is itemValue, else null
	ItemValue *CCMMultiItemExtensionEntry `jazz:"itemValue"`

	// A collection of zero of more items assigned as the value of the custom
	// attribute if the type is itemList, else null
	ItemList []*CCMMultiItemExtensionEntry `jazz:"itemList"`
}

// CCMExtensionEntryType contains the reflection type of CCMExtensionEntry
var goCCMExtensionEntryType = reflect.TypeOf(CCMExtensionEntry{})

// Spec returns the specification object for CCMExtensionEntry
func (o *CCMExtensionEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.ExtensionEntry",
		Type:       goCCMExtensionEntryType,
	}
}

// LoadAllFields of CCMExtensionEntry object
func (o *CCMExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.ItemValue,
		o.ItemList,
	)
}

// CCMTimeSheetEntry (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#time_SheetEntry_type_com_ibm_tea)
// This element represents a time sheet entry, each of the cells seen in the
// Time Tracking tab of a work item.
type CCMTimeSheetEntry struct {
	CCMBaseObject

	// The date for which the time sheet entry was entered
	StartDate *time.Time `jazz:"startDate"`

	// The time (in milliseconds) entered on the time sheet entry
	TimeSpent int64 `jazz:"timeSpent"`

	// The work item type (e.g. Defect)
	WorkType string `jazz:"workType"`

	// The description of the time code (e.g. Coding)
	TimeCode string `jazz:"timeCode"`

	// The identifier of the time code (e.g. timecode.literal.l2)
	TimeCodeId string `jazz:"timeCodeId"`

	// Work item to which the time sheet entry is related to.
	WorkItem *CCMMultiItemExtensionEntry `jazz:"workItem"`
}

// CCMTimeSheetEntryType contains the reflection type of CCMTimeSheetEntry
var goCCMTimeSheetEntryType = reflect.TypeOf(CCMTimeSheetEntry{})

// Spec returns the specification object for CCMTimeSheetEntry
func (o *CCMTimeSheetEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "timeSheetEntry",
		TypeID:     "com.ibm.team.workitem.TimeSheetEntry",
		Type:       goCCMTimeSheetEntryType,
	}
}

// Load CCMTimeSheetEntry object
func (o *CCMTimeSheetEntry) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of CCMTimeSheetEntry object
func (o *CCMTimeSheetEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.WorkItem,
	)
}

// CCMItem (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_repository_Item)
// Item The only time you're likely to see a raw Item is when using the referencedItem
// field of a Reference. Most of the time you'll want to fetch whichever concrete item
// type is represented by this artifact (e.g. a Work Item). The only standard field here
// likely to be useful is itemId, which can be used to look up the concrete element.
// This element is always contained in a com.ibm.team.links.Reference, and represents
// whether the reference is by uri or by itemId.
type CCMItem struct {
	CCMBaseObject

	// Type of item
	ItemType string `jazz:"itemType"`

	// The UUID representing the item in storage
	ItemId string `jazz:"itemId"`
}

// CCMItemType contains the reflection type of CCMItem
var goCCMItemType = reflect.TypeOf(CCMItem{})

// Spec returns the specification object for CCMItem
func (o *CCMItem) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.Item",
		Type:       goCCMItemType,
	}
}

// LoadAllFields of CCMItem object
func (o *CCMItem) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type CCMBooleanExtensionEntry struct {
	CCMBaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value bool `jazz:"value"`
}

// CCMBooleanExtensionEntryType contains the reflection type of CCMBooleanExtensionEntry
var goCCMBooleanExtensionEntryType = reflect.TypeOf(CCMBooleanExtensionEntry{})

// Spec returns the specification object for CCMBooleanExtensionEntry
func (o *CCMBooleanExtensionEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.BooleanExtensionEntry",
		Type:       goCCMBooleanExtensionEntryType,
	}
}

// LoadAllFields of CCMBooleanExtensionEntry object
func (o *CCMBooleanExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type CCMIntExtensionEntry struct {
	CCMBaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value int `jazz:"value"`
}

// CCMIntExtensionEntryType contains the reflection type of CCMIntExtensionEntry
var goCCMIntExtensionEntryType = reflect.TypeOf(CCMIntExtensionEntry{})

// Spec returns the specification object for CCMIntExtensionEntry
func (o *CCMIntExtensionEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.IntExtensionEntry",
		Type:       goCCMIntExtensionEntryType,
	}
}

// LoadAllFields of CCMIntExtensionEntry object
func (o *CCMIntExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type CCMLongExtensionEntry struct {
	CCMBaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value int64 `jazz:"value"`
}

// CCMLongExtensionEntryType contains the reflection type of CCMLongExtensionEntry
var goCCMLongExtensionEntryType = reflect.TypeOf(CCMLongExtensionEntry{})

// Spec returns the specification object for CCMLongExtensionEntry
func (o *CCMLongExtensionEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.LongExtensionEntry",
		Type:       goCCMLongExtensionEntryType,
	}
}

// LoadAllFields of CCMLongExtensionEntry object
func (o *CCMLongExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type CCMStringExtensionEntry struct {
	CCMBaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value string `jazz:"value"`
}

// CCMStringExtensionEntryType contains the reflection type of CCMStringExtensionEntry
var goCCMStringExtensionEntryType = reflect.TypeOf(CCMStringExtensionEntry{})

// Spec returns the specification object for CCMStringExtensionEntry
func (o *CCMStringExtensionEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.StringExtensionEntry",
		Type:       goCCMStringExtensionEntryType,
	}
}

// LoadAllFields of CCMStringExtensionEntry object
func (o *CCMStringExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type CCMMediumStringExtensionEntry struct {
	CCMBaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value string `jazz:"value"`
}

// CCMMediumStringExtensionEntryType contains the reflection type of CCMMediumStringExtensionEntry
var goCCMMediumStringExtensionEntryType = reflect.TypeOf(CCMMediumStringExtensionEntry{})

// Spec returns the specification object for CCMMediumStringExtensionEntry
func (o *CCMMediumStringExtensionEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.MediumStringExtensionEntry",
		Type:       goCCMMediumStringExtensionEntryType,
	}
}

// LoadAllFields of CCMMediumStringExtensionEntry object
func (o *CCMMediumStringExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type CCMLargeStringExtensionEntry struct {
	CCMBaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value string `jazz:"value"`
}

// CCMLargeStringExtensionEntryType contains the reflection type of CCMLargeStringExtensionEntry
var goCCMLargeStringExtensionEntryType = reflect.TypeOf(CCMLargeStringExtensionEntry{})

// Spec returns the specification object for CCMLargeStringExtensionEntry
func (o *CCMLargeStringExtensionEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.LargeStringExtensionEntry",
		Type:       goCCMLargeStringExtensionEntryType,
	}
}

// LoadAllFields of CCMLargeStringExtensionEntry object
func (o *CCMLargeStringExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type CCMTimestampExtensionEntry struct {
	CCMBaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value *time.Time `jazz:"value"`
}

// CCMTimestampExtensionEntryType contains the reflection type of CCMTimestampExtensionEntry
var goCCMTimestampExtensionEntryType = reflect.TypeOf(CCMTimestampExtensionEntry{})

// Spec returns the specification object for CCMTimestampExtensionEntry
func (o *CCMTimestampExtensionEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.TimestampExtensionEntry",
		Type:       goCCMTimestampExtensionEntryType,
	}
}

// LoadAllFields of CCMTimestampExtensionEntry object
func (o *CCMTimestampExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type CCMBigDecimalExtensionEntry struct {
	CCMBaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value float64 `jazz:"value"`
}

// CCMBigDecimalExtensionEntryType contains the reflection type of CCMBigDecimalExtensionEntry
var goCCMBigDecimalExtensionEntryType = reflect.TypeOf(CCMBigDecimalExtensionEntry{})

// Spec returns the specification object for CCMBigDecimalExtensionEntry
func (o *CCMBigDecimalExtensionEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.BigDecimalExtensionEntry",
		Type:       goCCMBigDecimalExtensionEntryType,
	}
}

// LoadAllFields of CCMBigDecimalExtensionEntry object
func (o *CCMBigDecimalExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type CCMItemExtensionEntry struct {
	CCMBaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value *CCMMultiItemExtensionEntry `jazz:"value"`
}

// CCMItemExtensionEntryType contains the reflection type of CCMItemExtensionEntry
var goCCMItemExtensionEntryType = reflect.TypeOf(CCMItemExtensionEntry{})

// Spec returns the specification object for CCMItemExtensionEntry
func (o *CCMItemExtensionEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.ItemExtensionEntry",
		Type:       goCCMItemExtensionEntryType,
	}
}

// LoadAllFields of CCMItemExtensionEntry object
func (o *CCMItemExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Value,
	)
}

type CCMMultiItemExtensionEntry struct {
	CCMBaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value []*CCMMultiItemExtensionEntry `jazz:"value"`
}

// CCMMultiItemExtensionEntryType contains the reflection type of CCMMultiItemExtensionEntry
var goCCMMultiItemExtensionEntryType = reflect.TypeOf(CCMMultiItemExtensionEntry{})

// Spec returns the specification object for CCMMultiItemExtensionEntry
func (o *CCMMultiItemExtensionEntry) Spec() *CCMObjectSpec {
	return &CCMObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.MultiItemExtensionEntry",
		Type:       goCCMMultiItemExtensionEntryType,
	}
}

// LoadAllFields of CCMMultiItemExtensionEntry object
func (o *CCMMultiItemExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Value,
	)
}
