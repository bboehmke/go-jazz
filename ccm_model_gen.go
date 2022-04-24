package jazz

// Code generated! DO NOT EDIT

import (
	"reflect"
	"time"
)

func init() {
	ccmRegisterType(new(ProjectArea))
	ccmRegisterType(new(TeamAreaHierarchyRecord))
	ccmRegisterType(new(TeamArea))
	ccmRegisterType(new(Contributor))
	ccmRegisterType(new(Iteration))
	ccmRegisterType(new(DevelopmentLine))
	ccmRegisterType(new(AuditableLink))
	ccmRegisterType(new(Reference))
	ccmRegisterType(new(ReferenceType))
	ccmRegisterType(new(ReadAccess))
	ccmRegisterType(new(Role))
	ccmRegisterType(new(RoleAssignment))
	ccmRegisterType(new(Workspace))
	ccmRegisterType(new(Property))
	ccmRegisterType(new(Component))
	ccmRegisterType(new(ChangeSet))
	ccmRegisterType(new(BuildDefinition))
	ccmRegisterType(new(BuildResult))
	ccmRegisterType(new(CompilationResult))
	ccmRegisterType(new(UnitTestResult))
	ccmRegisterType(new(UnitTestEvent))
	ccmRegisterType(new(BuildEngine))
	ccmRegisterType(new(WorkItem))
	ccmRegisterType(new(Comment))
	ccmRegisterType(new(Attribute))
	ccmRegisterType(new(Approval))
	ccmRegisterType(new(ApprovalDescriptor))
	ccmRegisterType(new(State))
	ccmRegisterType(new(Resolution))
	ccmRegisterType(new(WorkItemType))
	ccmRegisterType(new(Literal))
	ccmRegisterType(new(Category))
	ccmRegisterType(new(Deliverable))
	ccmRegisterType(new(ExtensionEntry))
	ccmRegisterType(new(TimeSheetEntry))
	ccmRegisterType(new(Item))
	ccmRegisterType(new(BooleanExtensionEntry))
	ccmRegisterType(new(IntExtensionEntry))
	ccmRegisterType(new(LongExtensionEntry))
	ccmRegisterType(new(StringExtensionEntry))
	ccmRegisterType(new(MediumStringExtensionEntry))
	ccmRegisterType(new(LargeStringExtensionEntry))
	ccmRegisterType(new(TimestampExtensionEntry))
	ccmRegisterType(new(BigDecimalExtensionEntry))
	ccmRegisterType(new(ItemExtensionEntry))
	ccmRegisterType(new(MultiItemExtensionEntry))
}

// ProjectArea (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#projectArea_type_com_ibm_team_pr)
// This element represents a Project Area.
type ProjectArea struct {
	BaseObject

	// The human-readable name of the project area (e.g. "My Project")
	Name string `jazz:"name"`

	// A list of members of this project
	TeamMembers []*Contributor `jazz:"teamMembers"`

	// A list of records reflecting the team area hierarchy for this project area
	TeamAreaHierarchy []*TeamAreaHierarchyRecord `jazz:"teamAreaHierarchy"`

	// A list of development lines for this project area
	DevelopmentLines []*DevelopmentLine `jazz:"developmentLines"`

	// The main development line for this project area
	ProjectDevelopmentLine *DevelopmentLine `jazz:"projectDevelopmentLine"`

	// The roles defined in the project area
	Roles []*Role `jazz:"roles"`

	// The role assignments defined in the project area
	RoleAssignments []*RoleAssignment `jazz:"roleAssignments"`

	// All the team areas contained in the project area
	AllTeamAreas []*TeamArea `jazz:"allTeamAreas"`
}

// ProjectAreaType contains the reflection type of ProjectArea
var goProjectAreaType = reflect.TypeOf(ProjectArea{})

// Spec returns the specification object for ProjectArea
func (o *ProjectArea) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "projectArea",
		TypeID:     "com.ibm.team.process.ProjectArea",
		Type:       goProjectAreaType,
	}
}

// Load ProjectArea object
func (o *ProjectArea) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of ProjectArea object
func (o *ProjectArea) LoadAllFields() error {
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

// TeamAreaHierarchyRecord (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_process_TeamAreaHie)
// This element appears only inside a Project Area, and represents a piece of
// a team area hierarchy.
type TeamAreaHierarchyRecord struct {
	BaseObject

	// The parent team area
	Parent *TeamArea `jazz:"parent"`

	// The children team areas of the parent team area
	Children []*TeamArea `jazz:"children"`
}

// TeamAreaHierarchyRecordType contains the reflection type of TeamAreaHierarchyRecord
var goTeamAreaHierarchyRecordType = reflect.TypeOf(TeamAreaHierarchyRecord{})

// Spec returns the specification object for TeamAreaHierarchyRecord
func (o *TeamAreaHierarchyRecord) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.process.TeamAreaHierarchyRecord",
		Type:       goTeamAreaHierarchyRecordType,
	}
}

// LoadAllFields of TeamAreaHierarchyRecord object
func (o *TeamAreaHierarchyRecord) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Parent,
		o.Children,
	)
}

// TeamArea (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#teamArea_type_com_ibm_team_proce)
// This element represents a Team Area.
type TeamArea struct {
	BaseObject

	// The human-readable name of the project area (e.g. "My Team")
	Name string `jazz:"name"`

	// A fully-qualified team area name, slash-separated, including all parent
	// team areas (e.g. "/My Parent Team/My Team").
	QualifiedName string `jazz:"qualifiedName"`

	// A list of members of this team area
	TeamMembers []*Contributor `jazz:"teamMembers"`

	// The project area containing this team area
	ProjectArea *ProjectArea `jazz:"projectArea"`

	// The roles defined in the team area
	Roles []*Role `jazz:"roles"`

	// The role assignments defined in the team area
	RoleAssignments []*RoleAssignment `jazz:"roleAssignments"`

	// The parent team area
	ParentTeamArea *TeamArea `jazz:"parentTeamArea"`
}

// TeamAreaType contains the reflection type of TeamArea
var goTeamAreaType = reflect.TypeOf(TeamArea{})

// Spec returns the specification object for TeamArea
func (o *TeamArea) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "teamArea",
		TypeID:     "com.ibm.team.process.TeamArea",
		Type:       goTeamAreaType,
	}
}

// Load TeamArea object
func (o *TeamArea) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of TeamArea object
func (o *TeamArea) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.TeamMembers,
		o.ProjectArea,
		o.Roles,
		o.RoleAssignments,
		o.ParentTeamArea,
	)
}

// Contributor (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#contributor)
// This element represents a Contributor (user).
type Contributor struct {
	BaseObject

	// The human-readable name of the contributor (e.g. "James Moody")
	Name string `jazz:"name"`

	// The email address of the contributor
	EmailAddress string `jazz:"emailAddress"`

	// The userId of the contributor, unique in this application (e.g. "jmoody")
	UserId string `jazz:"userId"`
}

// ContributorType contains the reflection type of Contributor
var goContributorType = reflect.TypeOf(Contributor{})

// Spec returns the specification object for Contributor
func (o *Contributor) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "contributor",
		TypeID:     "com.ibm.team.repository.Contributor",
		Type:       goContributorType,
	}
}

// Load Contributor object
func (o *Contributor) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of Contributor object
func (o *Contributor) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// Iteration (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#iteration_type_com_ibm_team_proc)
// This element represents a single iteration (milestone, sprint).
type Iteration struct {
	BaseObject

	// The human-readable name of this iteration (e.g. "M1")
	Name string `jazz:"name"`

	// The identifier of this iteration (e.g. "3.0M1")
	Id string `jazz:"id"`

	// The start date of this iteration
	StartDate *time.Time `jazz:"startDate"`

	// The end date of this iteration
	EndDate *time.Time `jazz:"endDate"`

	// The parent iteration of this iteration, if any
	Parent *Iteration `jazz:"parent"`

	// The immediate child iterations of this iteration, if any
	Children []*Iteration `jazz:"children"`

	// The development line in which this iteration appears
	DevelopmentLine *DevelopmentLine `jazz:"developmentLine"`

	// Whether or not this iteration is marked as having deliverables associated
	// with it
	HasDeliverable bool `jazz:"hasDeliverable"`
}

// IterationType contains the reflection type of Iteration
var goIterationType = reflect.TypeOf(Iteration{})

// Spec returns the specification object for Iteration
func (o *Iteration) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "iteration",
		TypeID:     "com.ibm.team.process.Iteration",
		Type:       goIterationType,
	}
}

// Load Iteration object
func (o *Iteration) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of Iteration object
func (o *Iteration) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Parent,
		o.Children,
		o.DevelopmentLine,
	)
}

// DevelopmentLine (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#developmentLine_type_com_ibm_tea)
// This element represents a development line.
type DevelopmentLine struct {
	BaseObject

	// The human-readable name of this development line (e.g. "Maintenance
	// Development")
	Name string `jazz:"name"`

	// The start date of this development line
	StartDate *time.Time `jazz:"startDate"`

	// The end date of this development line
	EndDate *time.Time `jazz:"endDate"`

	// The child iterations of this development line
	Iterations []*Iteration `jazz:"iterations"`

	// The project area containing this development line
	ProjectArea *ProjectArea `jazz:"projectArea"`

	// The iteration marked as current in this development line
	CurrentIteration *Iteration `jazz:"currentIteration"`
}

// DevelopmentLineType contains the reflection type of DevelopmentLine
var goDevelopmentLineType = reflect.TypeOf(DevelopmentLine{})

// Spec returns the specification object for DevelopmentLine
func (o *DevelopmentLine) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "developmentLine",
		TypeID:     "com.ibm.team.process.DevelopmentLine",
		Type:       goDevelopmentLineType,
	}
}

// Load DevelopmentLine object
func (o *DevelopmentLine) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of DevelopmentLine object
func (o *DevelopmentLine) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Iterations,
		o.ProjectArea,
		o.CurrentIteration,
	)
}

// AuditableLink (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#auditableLink)
// This element represents a link from one artifact to another. These links
// may be either within the same repository, or between one artifact in this
// repository and one external artifact. References (source and target) may be
// made either by uri (for any artifact) or by referencedItem (in the case of
// local artifacts).
type AuditableLink struct {
	BaseObject

	// The id of this link type (e.g. "com.ibm.team.workitem.parentChild"). This
	// describes the relationship represented by this link.
	Name string `jazz:"name"`

	// The source of the link
	SourceRef *Reference `jazz:"sourceRef"`

	// The target of the link
	TargetRef *Reference `jazz:"targetRef"`
}

// AuditableLinkType contains the reflection type of AuditableLink
var goAuditableLinkType = reflect.TypeOf(AuditableLink{})

// Spec returns the specification object for AuditableLink
func (o *AuditableLink) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "auditableLink",
		TypeID:     "",
		Type:       goAuditableLinkType,
	}
}

// Load AuditableLink object
func (o *AuditableLink) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of AuditableLink object
func (o *AuditableLink) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.SourceRef,
		o.TargetRef,
	)
}

// Reference (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_links_Reference)
// This element is always contained in an auditableLink, and represents either
// the source or target reference of a link. The reference may be either by
// uri (for any artifact) or by referencedItem (in the case of local
// artifacts). Which one can be determined by the referenceType field.
type Reference struct {
	BaseObject

	// A human-readable comment about the reference. In some cases the comment may
	// suffice rather than fetching the content on the other end of the link. For
	// example, a reference pointing to a work item may contain the id and summary
	// of the work item ("12345: Summary of my work item").
	Comment string `jazz:"comment"`

	// This element indicates whether the reference is by uri or by itemId.
	ReferenceType *ReferenceType `jazz:"referenceType"`

	// The URI of the element referenced. This is only valid if this Reference is
	// a URI reference.
	Uri string `jazz:"uri"`

	// The referenced item. This is only valid if this Reference is an Item
	// reference.
	ReferencedItem *Item `jazz:"referencedItem"`

	// Get the extra information associated with the reference. May be null.
	ExtraInfo string `jazz:"extraInfo"`

	// Internal.
	ContentType string `jazz:"contentType"`
}

// ReferenceType contains the reflection type of Reference
var goReferenceType = reflect.TypeOf(Reference{})

// Spec returns the specification object for Reference
func (o *Reference) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.links.Reference",
		Type:       goReferenceType,
	}
}

// LoadAllFields of Reference object
func (o *Reference) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.ReferenceType,
		o.ReferencedItem,
	)
}

// ReferenceType (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_links_ReferenceType)
// This element represents a reference type, indicating whether a reference is
// by URI or itemID.
type ReferenceType struct {
	BaseObject

	// Either "ITEM_REFERENCE" or "URI_REFERENCE"
	Literal string `jazz:"literal"`

	// Either 0 (for ITEM_REFERENCE) or 2 (for URI_REFERENCE). Use literal
	// instead.
	Value int `jazz:"value"`
}

// ReferenceTypeType contains the reflection type of ReferenceType
var goReferenceTypeType = reflect.TypeOf(ReferenceType{})

// Spec returns the specification object for ReferenceType
func (o *ReferenceType) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.links.ReferenceType",
		Type:       goReferenceTypeType,
	}
}

// LoadAllFields of ReferenceType object
func (o *ReferenceType) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// ReadAccess (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#readAccess)
// The readAccess element represents a mapping of contributors to project
// areas that each contributor has permissions to read.
type ReadAccess struct {
	BaseObject

	// The itemId of the Contributor
	ContributorItemId string `jazz:"contributorItemId"`

	// The itemID of the context object associated with the contributor (i.e. the
	// project area)
	ContributorContextId string `jazz:"contributorContextId"`
}

// ReadAccessType contains the reflection type of ReadAccess
var goReadAccessType = reflect.TypeOf(ReadAccess{})

// Spec returns the specification object for ReadAccess
func (o *ReadAccess) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "readAccess",
		TypeID:     "",
		Type:       goReadAccessType,
	}
}

// Load ReadAccess object
func (o *ReadAccess) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of ReadAccess object
func (o *ReadAccess) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// Role (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_process_Role)
type Role struct {
	BaseObject

	// The role Id
	Id string `jazz:"id"`

	// The role name
	Name string `jazz:"name"`

	// The role description
	Description string `jazz:"description"`
}

// RoleType contains the reflection type of Role
var goRoleType = reflect.TypeOf(Role{})

// Spec returns the specification object for Role
func (o *Role) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.process.Role",
		Type:       goRoleType,
	}
}

// LoadAllFields of Role object
func (o *Role) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// RoleAssignment (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_process_RoleAssignm)
type RoleAssignment struct {
	BaseObject

	// The contributor with assigned roles
	Contributor *Contributor `jazz:"contributor"`

	// The roles assigned to the contributor
	ContributorRoles []*Role `jazz:"contributorRoles"`
}

// RoleAssignmentType contains the reflection type of RoleAssignment
var goRoleAssignmentType = reflect.TypeOf(RoleAssignment{})

// Spec returns the specification object for RoleAssignment
func (o *RoleAssignment) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.process.RoleAssignment",
		Type:       goRoleAssignmentType,
	}
}

// LoadAllFields of RoleAssignment object
func (o *RoleAssignment) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Contributor,
		o.ContributorRoles,
	)
}

// Workspace (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#workspace_type_com_ibm_team_scm)
// This element represents an SCM Workspace or Stream
type Workspace struct {
	BaseObject

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
	Properties []*Property `jazz:"properties"`

	// The owner of the workspace or stream
	Contributor *Contributor `jazz:"contributor"`
}

// WorkspaceType contains the reflection type of Workspace
var goWorkspaceType = reflect.TypeOf(Workspace{})

// Spec returns the specification object for Workspace
func (o *Workspace) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "scm",
		ElementID:  "workspace",
		TypeID:     "com.ibm.team.scm.Workspace",
		Type:       goWorkspaceType,
	}
}

// Load Workspace object
func (o *Workspace) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of Workspace object
func (o *Workspace) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Properties,
		o.Contributor,
	)
}

// Property (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_scm_Property)
// This element only occurs in a workspace, and represents a property of a
// Workspace or Stream
type Property struct {
	BaseObject

	// The property key
	Key string `jazz:"key"`
}

// PropertyType contains the reflection type of Property
var goPropertyType = reflect.TypeOf(Property{})

// Spec returns the specification object for Property
func (o *Property) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "scm",
		ElementID:  "",
		TypeID:     "com.ibm.team.scm.Property",
		Type:       goPropertyType,
	}
}

// LoadAllFields of Property object
func (o *Property) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// Component (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#component_type_com_ibm_team_scm)
// This element represents an SCM Component
type Component struct {
	BaseObject

	// The name of the component
	Name string `jazz:"name"`
}

// ComponentType contains the reflection type of Component
var goComponentType = reflect.TypeOf(Component{})

// Spec returns the specification object for Component
func (o *Component) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "scm",
		ElementID:  "component",
		TypeID:     "com.ibm.team.scm.Component",
		Type:       goComponentType,
	}
}

// Load Component object
func (o *Component) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of Component object
func (o *Component) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// ChangeSet (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#changeSet_type_com_ibm_team_scm)
// This element represents an SCM Change Set
type ChangeSet struct {
	BaseObject

	// The comment on the change set
	Comment string `jazz:"comment"`

	// The owner of the change set
	Owner *Contributor `jazz:"owner"`
}

// ChangeSetType contains the reflection type of ChangeSet
var goChangeSetType = reflect.TypeOf(ChangeSet{})

// Spec returns the specification object for ChangeSet
func (o *ChangeSet) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "scm",
		ElementID:  "changeSet",
		TypeID:     "com.ibm.team.scm.ChangeSet",
		Type:       goChangeSetType,
	}
}

// Load ChangeSet object
func (o *ChangeSet) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of ChangeSet object
func (o *ChangeSet) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Owner,
	)
}

// BuildDefinition (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#buildDefinition_type_com_ibm_tea)
// This element represents a Build Definition.
type BuildDefinition struct {
	BaseObject

	// The id of the build definition
	Id string `jazz:"id"`

	// The description of the build definition
	Description string `jazz:"description"`

	// The project area containing the build definition
	ProjectArea *ProjectArea `jazz:"projectArea"`

	// The team area containing the build definition
	TeamArea *TeamArea `jazz:"teamArea"`
}

// BuildDefinitionType contains the reflection type of BuildDefinition
var goBuildDefinitionType = reflect.TypeOf(BuildDefinition{})

// Spec returns the specification object for BuildDefinition
func (o *BuildDefinition) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "build",
		ElementID:  "buildDefinition",
		TypeID:     "com.ibm.team.build.BuildDefinition",
		Type:       goBuildDefinitionType,
	}
}

// Load BuildDefinition object
func (o *BuildDefinition) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of BuildDefinition object
func (o *BuildDefinition) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.ProjectArea,
		o.TeamArea,
	)
}

// BuildResult (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#buildResult_type_com_ibm_team_bu)
// This element represents a Build Result.
type BuildResult struct {
	BaseObject

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
	BuildDefinition *BuildDefinition `jazz:"buildDefinition"`

	// The contributor who requested the build
	Creator *Contributor `jazz:"creator"`

	// The engine the build ran on
	BuildEngine *BuildEngine `jazz:"buildEngine"`

	// Code compilation results
	CompilationResults []*CompilationResult `jazz:"compilationResults"`

	// Unit test results
	UnitTestResults []*UnitTestResult `jazz:"unitTestResults"`

	// Unit test changes from the previous build
	UnitTestEvents []*UnitTestEvent `jazz:"unitTestEvents"`
}

// BuildResultType contains the reflection type of BuildResult
var goBuildResultType = reflect.TypeOf(BuildResult{})

// Spec returns the specification object for BuildResult
func (o *BuildResult) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "build",
		ElementID:  "buildResult",
		TypeID:     "com.ibm.team.build.BuildResult",
		Type:       goBuildResultType,
	}
}

// Load BuildResult object
func (o *BuildResult) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of BuildResult object
func (o *BuildResult) LoadAllFields() error {
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

// CompilationResult (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_build_CompilationRe)
// This element only occurs in a buildResult. The number of errors and
// warnings for a particular component in the containing build result
type CompilationResult struct {
	BaseObject

	// The component for which the errors and warnings are being reported
	Component string `jazz:"component"`

	// The number of compilation errors for the component in the containing build
	// result
	Errors int64 `jazz:"errors"`

	// The umber of compilation warnings for the component in the containing build
	// result
	Warnings int64 `jazz:"warnings"`
}

// CompilationResultType contains the reflection type of CompilationResult
var goCompilationResultType = reflect.TypeOf(CompilationResult{})

// Spec returns the specification object for CompilationResult
func (o *CompilationResult) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "build",
		ElementID:  "",
		TypeID:     "com.ibm.team.build.CompilationResult",
		Type:       goCompilationResultType,
	}
}

// LoadAllFields of CompilationResult object
func (o *CompilationResult) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// UnitTestResult (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_build_UnitTestResul)
// This element only occurs in a buildResult. The number of unit tests run,
// along with number of failures and errors, for a particular component in the
// containing build result
type UnitTestResult struct {
	BaseObject

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

// UnitTestResultType contains the reflection type of UnitTestResult
var goUnitTestResultType = reflect.TypeOf(UnitTestResult{})

// Spec returns the specification object for UnitTestResult
func (o *UnitTestResult) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "build",
		ElementID:  "",
		TypeID:     "com.ibm.team.build.UnitTestResult",
		Type:       goUnitTestResultType,
	}
}

// LoadAllFields of UnitTestResult object
func (o *UnitTestResult) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// UnitTestEvent (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_build_UnitTestEvent)
// This element only occurs in a buildResult. It represents a single unit test
// execution, along with a pass, fail or regression label
type UnitTestEvent struct {
	BaseObject

	// The component for which the test and event is being reported
	Component string `jazz:"component"`

	// The name of the unit test run
	Test string `jazz:"test"`

	// Indication of test passing, failing or regressing. James: To do, provide
	// the literals here.
	Event string `jazz:"event"`
}

// UnitTestEventType contains the reflection type of UnitTestEvent
var goUnitTestEventType = reflect.TypeOf(UnitTestEvent{})

// Spec returns the specification object for UnitTestEvent
func (o *UnitTestEvent) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "build",
		ElementID:  "",
		TypeID:     "com.ibm.team.build.UnitTestEvent",
		Type:       goUnitTestEventType,
	}
}

// LoadAllFields of UnitTestEvent object
func (o *UnitTestEvent) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// BuildEngine (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#buildEngine_type_com_ibm_team_bu)
// This element represents a build engine.
type BuildEngine struct {
	BaseObject

	// The id of this build engine
	Id string `jazz:"id"`
}

// BuildEngineType contains the reflection type of BuildEngine
var goBuildEngineType = reflect.TypeOf(BuildEngine{})

// Spec returns the specification object for BuildEngine
func (o *BuildEngine) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "build",
		ElementID:  "buildEngine",
		TypeID:     "com.ibm.team.build.BuildEngine",
		Type:       goBuildEngineType,
	}
}

// Load BuildEngine object
func (o *BuildEngine) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of BuildEngine object
func (o *BuildEngine) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// WorkItem (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#workItem_type_com_ibm_team_worki)
// This element represents a Work Item.
type WorkItem struct {
	BaseObject

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
	Creator *Contributor `jazz:"creator"`

	// The contributor who owns the work item
	Owner *Contributor `jazz:"owner"`

	// The category to which the work item is assigned. In the UI, this is called
	// "Filed Against".
	Category *Category `jazz:"category"`

	// A collection of zero or more comments appended to the work item
	Comments []*Comment `jazz:"comments"`

	// A collection of zero or more "custom attributes" attached to the work item.
	// These are user-defined attributes (as opposed to the built-in attributes
	// elsewhere in this list).
	CustomAttributes []*Attribute `jazz:"customAttributes"`

	// A collection of zero or more Contributors who are subscribed to the work
	// item
	Subscriptions []*Contributor `jazz:"subscriptions"`

	// The project area to which the work item belongs
	ProjectArea *ProjectArea `jazz:"projectArea"`

	// The Contributor who resolved the work item, or null if the work item has
	// not been resolved
	Resolver *Contributor `jazz:"resolver"`

	// A collection of zero or more Approvals attached to the work item
	Approvals []*Approval `jazz:"approvals"`

	// A collection of zero or more Approval Descriptors attached to the work item
	ApprovalDescriptors []*ApprovalDescriptor `jazz:"approvalDescriptors"`

	// The iteration that the work item is "Planned For"
	Target *Iteration `jazz:"target"`

	// The deliverable that the work item is "Found In"
	FoundIn *Deliverable `jazz:"foundIn"`

	// A collection of zero or more WorkItem elements, representing the entire
	// history of the work item. Each state the work item has ever been in is
	// reflected in this history list.
	ItemHistory []*WorkItem `jazz:"itemHistory"`

	// The team area to which the work item belongs
	TeamArea *TeamArea `jazz:"teamArea"`

	// The state of the work item (e.g. "Resolved", "In Progress", "New"). The
	// states are user-defined as part of the project area process.
	State *State `jazz:"state"`

	// The resolution of the work item (e.g. "Duplicate", "Invalid", "Fixed"). The
	// resolutions are user-defined as part of the project area process.
	Resolution *Resolution `jazz:"resolution"`

	// The type of the work item (e.g. "Defect", "Task", "Story"). The work item
	// types are user-defined as part of the project area process.
	Type *WorkItemType `jazz:"type"`

	// The severity of the work item (e.g. "Critical", "Normal", "Blocker"). The
	// work item severities are user-defined as part of the project area process.
	Severity *Literal `jazz:"severity"`

	// The priority of the work item (e.g. "High", "Medium", "Low"). The work item
	// priorities are user-defined as part of the project area process.
	Priority *Literal `jazz:"priority"`

	// The parent work item of this work item, if one exists
	Parent *WorkItem `jazz:"parent"`

	// A collection of zero or more child work items
	Children []*WorkItem `jazz:"children"`

	// A collection of zero or more work items which this work item blocks
	Blocks []*WorkItem `jazz:"blocks"`

	// A collection of zero or more work items which block this work item
	DependsOn []*WorkItem `jazz:"dependsOn"`

	// A collection of zero or more work items which are closed as duplicates of
	// this work item
	DuplicatedBy []*WorkItem `jazz:"duplicatedBy"`

	// A collection of zero or more work items which this work item is a duplicate
	// of
	DuplicateOf []*WorkItem `jazz:"duplicateOf"`

	// A collection of zero of more work items which this work item is related to
	Related []*WorkItem `jazz:"related"`

	// A collection of zero or more items linked to the work item as custom
	// attributes
	ItemExtensions []*ItemExtensionEntry `jazz:"itemExtensions"`

	// A collection of zero or more lists of items linked to the work item as
	// custom attributes
	MultiItemExtensions []*MultiItemExtensionEntry `jazz:"multiItemExtensions"`

	// A collection of zero or more custom attributes of type medium string
	MediumStringExtensions []*MediumStringExtensionEntry `jazz:"mediumStringExtensions"`

	// A collection of zero or more custom attributes of type boolean
	BooleanExtensions []*BooleanExtensionEntry `jazz:"booleanExtensions"`

	// A collection of zero or more custom attributes of type timestamp
	TimestampExtensions []*TimestampExtensionEntry `jazz:"timestampExtensions"`

	// A collection of zero or more custom attributes of type long
	LongExtensions []*LongExtensionEntry `jazz:"longExtensions"`

	// A collection of zero or more custom attributes of type integer
	IntExtensions []*IntExtensionEntry `jazz:"intExtensions"`

	// A collection of zero or more custom attributes of type big decimal
	BigDecimalExtensions []*BigDecimalExtensionEntry `jazz:"bigDecimalExtensions"`

	// A collection of zero or more custom attributes of type large string
	LargeStringExtensions []*LargeStringExtensionEntry `jazz:"largeStringExtensions"`

	// A collection of zero or more custom attributes of type string
	StringExtensions []*StringExtensionEntry `jazz:"stringExtensions"`

	// A collection of zero or more custom attributes of all types
	AllExtensions []*ExtensionEntry `jazz:"allExtensions"`

	// A collection of zero or more timesheet entries linked to the work item
	TimeSheetEntries []*TimeSheetEntry `jazz:"timeSheetEntries"`

	// The work item's planned start date as specified in the plan.
	PlannedStartDate *time.Time `jazz:"plannedStartDate"`

	// The work item's planned end date as specified in the plan.
	PlannedEndDate *time.Time `jazz:"plannedEndDate"`
}

// WorkItemType contains the reflection type of WorkItem
var goWorkItemType = reflect.TypeOf(WorkItem{})

// Spec returns the specification object for WorkItem
func (o *WorkItem) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "workItem",
		TypeID:     "com.ibm.team.workitem.WorkItem",
		Type:       goWorkItemType,
	}
}

// Load WorkItem object
func (o *WorkItem) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of WorkItem object
func (o *WorkItem) LoadAllFields() error {
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

// Comment (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_Comment)
// This element represents a single work item comment.
type Comment struct {
	BaseObject

	// The date/time that the comment was saved in the work item
	CreationDate *time.Time `jazz:"creationDate"`

	// The string content of the comment
	Content string `jazz:"content"`

	// Whether or not the comment has been edited
	Edited bool `jazz:"edited"`

	// The contributor who created the comment
	Creator *Contributor `jazz:"creator"`
}

// CommentType contains the reflection type of Comment
var goCommentType = reflect.TypeOf(Comment{})

// Spec returns the specification object for Comment
func (o *Comment) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.Comment",
		Type:       goCommentType,
	}
}

// LoadAllFields of Comment object
func (o *Comment) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Creator,
	)
}

// Attribute (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_Attribute)
// This element represents information about a custom attribute declaration.
// Custom attribute declarations are process-specific.
type Attribute struct {
	BaseObject

	// An identifier for the custom attribute, unique within a project area
	Identifier string `jazz:"identifier"`

	// The data type of the attribute value
	AttributeType string `jazz:"attributeType"`

	// Whether or not the attribute is built-in
	BuiltIn bool `jazz:"builtIn"`

	// The project in which the attribute is defined
	ProjectArea *ProjectArea `jazz:"projectArea"`
}

// AttributeType contains the reflection type of Attribute
var goAttributeType = reflect.TypeOf(Attribute{})

// Spec returns the specification object for Attribute
func (o *Attribute) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.Attribute",
		Type:       goAttributeType,
	}
}

// LoadAllFields of Attribute object
func (o *Attribute) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.ProjectArea,
	)
}

// Approval (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_Approval)
// This element represents an approval from a single contributor with a
// particular state.
type Approval struct {
	BaseObject

	// The state of the approval
	StateIdentifier string `jazz:"stateIdentifier"`

	// The date the state was assigned
	StateDate *time.Time `jazz:"stateDate"`

	// The name of the state
	StateName string `jazz:"stateName"`

	// The contributor who is asked for approval
	Approver *Contributor `jazz:"approver"`
}

// ApprovalType contains the reflection type of Approval
var goApprovalType = reflect.TypeOf(Approval{})

// Spec returns the specification object for Approval
func (o *Approval) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.Approval",
		Type:       goApprovalType,
	}
}

// LoadAllFields of Approval object
func (o *Approval) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Approver,
	)
}

// ApprovalDescriptor (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_ApprovalDe)
// This element represents an approval descriptor aggregates approvals from
// contributors.
type ApprovalDescriptor struct {
	BaseObject

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
	Approvals []*Approval `jazz:"approvals"`
}

// ApprovalDescriptorType contains the reflection type of ApprovalDescriptor
var goApprovalDescriptorType = reflect.TypeOf(ApprovalDescriptor{})

// Spec returns the specification object for ApprovalDescriptor
func (o *ApprovalDescriptor) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.ApprovalDescriptor",
		Type:       goApprovalDescriptorType,
	}
}

// LoadAllFields of ApprovalDescriptor object
func (o *ApprovalDescriptor) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Approvals,
	)
}

// State (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_State)
// This element represents the state of a work item. States are defined by the
// user in the process specification for a project area.
type State struct {
	BaseObject

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

// StateType contains the reflection type of State
var goStateType = reflect.TypeOf(State{})

// Spec returns the specification object for State
func (o *State) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.State",
		Type:       goStateType,
	}
}

// LoadAllFields of State object
func (o *State) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// Resolution (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_Resolution)
// This element represents the resolution of a work item. This indicates how
// or why a work item was resolved; for example, "Fixed", "Invalid", "Won't
// Fix". Resolutions are process-dependent.
type Resolution struct {
	BaseObject

	// The id of the resolution (e.g. "com.ibm.team.workitem.defect.fixed"),
	// unique in a repository.
	Id string `jazz:"id"`

	// The name of the resolution (e.g. "Fixed"). Not necessarily unique.
	Name string `jazz:"name"`
}

// ResolutionType contains the reflection type of Resolution
var goResolutionType = reflect.TypeOf(Resolution{})

// Spec returns the specification object for Resolution
func (o *Resolution) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.Resolution",
		Type:       goResolutionType,
	}
}

// LoadAllFields of Resolution object
func (o *Resolution) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// WorkItemType (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_WorkItemTy)
// This element represents the type of a work item. Work item types are
// process-dependent.
type WorkItemType struct {
	BaseObject

	// The id of the type (e.g. "com.ibm.team.workitem.defect"), unique in a
	// repository.
	Id string `jazz:"id"`

	// The name of the type (e.g. "Defect"). Not necessarily unique.
	Name string `jazz:"name"`
}

// WorkItemTypeType contains the reflection type of WorkItemType
var goWorkItemTypeType = reflect.TypeOf(WorkItemType{})

// Spec returns the specification object for WorkItemType
func (o *WorkItemType) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.WorkItemType",
		Type:       goWorkItemTypeType,
	}
}

// LoadAllFields of WorkItemType object
func (o *WorkItemType) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// Literal (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_workitem_Literal)
// This element represents a user-defined literal value, used for priority and
// severity in a work item. Work item severities and priorities are
// process-dependent.
type Literal struct {
	BaseObject

	// The id of the literal (e.g. "com.ibm.team.workitem.blocking"), unique in a
	// repository.
	Id string `jazz:"id"`

	// The name of the literal (e.g. "Blocking"). Not necessarily unique.
	Name string `jazz:"name"`
}

// LiteralType contains the reflection type of Literal
var goLiteralType = reflect.TypeOf(Literal{})

// Spec returns the specification object for Literal
func (o *Literal) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.Literal",
		Type:       goLiteralType,
	}
}

// LoadAllFields of Literal object
func (o *Literal) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// Category (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#category_type_com_ibm_team_worki)
// This element represents a work item Category. Work item categories are
// process-dependent.
type Category struct {
	BaseObject

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

// CategoryType contains the reflection type of Category
var goCategoryType = reflect.TypeOf(Category{})

// Spec returns the specification object for Category
func (o *Category) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "category",
		TypeID:     "com.ibm.team.workitem.Category",
		Type:       goCategoryType,
	}
}

// Load Category object
func (o *Category) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of Category object
func (o *Category) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

// Deliverable (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#deliverable_type_com_ibm_team_wo)
// This element represents a deliverable, often used in Work Items to identify
// in which deliverable a work item was found ("Found In"). Deliverables are
// process-dependent.
type Deliverable struct {
	BaseObject

	// The name of the deliverable (e.g. "RTC 3.0")
	Name string `jazz:"name"`

	// A textual description of the deliverable
	Description string `jazz:"description"`

	// The creation date of the deliverable
	CreationDate *time.Time `jazz:"creationDate"`

	// The project area associated with the deliverable
	ProjectArea *ProjectArea `jazz:"projectArea"`

	// An optional link to a repository item associated with the deliverable. This
	// field should be treated as internal.
	Artifact *Item `jazz:"artifact"`
}

// DeliverableType contains the reflection type of Deliverable
var goDeliverableType = reflect.TypeOf(Deliverable{})

// Spec returns the specification object for Deliverable
func (o *Deliverable) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "deliverable",
		TypeID:     "com.ibm.team.workitem.Deliverable",
		Type:       goDeliverableType,
	}
}

// Load Deliverable object
func (o *Deliverable) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of Deliverable object
func (o *Deliverable) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.ProjectArea,
		o.Artifact,
	)
}

// ExtensionEntry (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#allExtensions_type_com_ibm_team)
// This element represents the value of a custom attribute.
type ExtensionEntry struct {
	BaseObject

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
	ItemValue *Item `jazz:"itemValue"`

	// A collection of zero of more items assigned as the value of the custom
	// attribute if the type is itemList, else null
	ItemList []*Item `jazz:"itemList"`
}

// ExtensionEntryType contains the reflection type of ExtensionEntry
var goExtensionEntryType = reflect.TypeOf(ExtensionEntry{})

// Spec returns the specification object for ExtensionEntry
func (o *ExtensionEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.workitem.ExtensionEntry",
		Type:       goExtensionEntryType,
	}
}

// LoadAllFields of ExtensionEntry object
func (o *ExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.ItemValue,
		o.ItemList,
	)
}

// TimeSheetEntry (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#time_SheetEntry_type_com_ibm_tea)
// This element represents a time sheet entry, each of the cells seen in the
// Time Tracking tab of a work item.
type TimeSheetEntry struct {
	BaseObject

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
	WorkItem *WorkItem `jazz:"workItem"`
}

// TimeSheetEntryType contains the reflection type of TimeSheetEntry
var goTimeSheetEntryType = reflect.TypeOf(TimeSheetEntry{})

// Spec returns the specification object for TimeSheetEntry
func (o *TimeSheetEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "timeSheetEntry",
		TypeID:     "com.ibm.team.workitem.TimeSheetEntry",
		Type:       goTimeSheetEntryType,
	}
}

// Load TimeSheetEntry object
func (o *TimeSheetEntry) Load() (err error) {
	o.init.Do(func() {
		if o.ReportableUrl == "" {
			err = o.ccm.get(o.Spec(), reflect.ValueOf(o), o.ItemId)
		}
	})
	return
}

// LoadAllFields of TimeSheetEntry object
func (o *TimeSheetEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.WorkItem,
	)
}

// Item (see https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#com_ibm_team_repository_Item)
// Item The only time you're likely to see a raw Item is when using the referencedItem
// field of a Reference. Most of the time you'll want to fetch whichever concrete item
// type is represented by this artifact (e.g. a Work Item). The only standard field here
// likely to be useful is itemId, which can be used to look up the concrete element.
// This element is always contained in a com.ibm.team.links.Reference, and represents
// whether the reference is by uri or by itemId.
type Item struct {
	BaseObject

	// Type of item
	ItemType string `jazz:"itemType"`

	// The UUID representing the item in storage
	ItemId string `jazz:"itemId"`
}

// ItemType contains the reflection type of Item
var goItemType = reflect.TypeOf(Item{})

// Spec returns the specification object for Item
func (o *Item) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "foundation",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.Item",
		Type:       goItemType,
	}
}

// LoadAllFields of Item object
func (o *Item) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type BooleanExtensionEntry struct {
	BaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value bool `jazz:"value"`
}

// BooleanExtensionEntryType contains the reflection type of BooleanExtensionEntry
var goBooleanExtensionEntryType = reflect.TypeOf(BooleanExtensionEntry{})

// Spec returns the specification object for BooleanExtensionEntry
func (o *BooleanExtensionEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.BooleanExtensionEntry",
		Type:       goBooleanExtensionEntryType,
	}
}

// LoadAllFields of BooleanExtensionEntry object
func (o *BooleanExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type IntExtensionEntry struct {
	BaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value int `jazz:"value"`
}

// IntExtensionEntryType contains the reflection type of IntExtensionEntry
var goIntExtensionEntryType = reflect.TypeOf(IntExtensionEntry{})

// Spec returns the specification object for IntExtensionEntry
func (o *IntExtensionEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.IntExtensionEntry",
		Type:       goIntExtensionEntryType,
	}
}

// LoadAllFields of IntExtensionEntry object
func (o *IntExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type LongExtensionEntry struct {
	BaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value int64 `jazz:"value"`
}

// LongExtensionEntryType contains the reflection type of LongExtensionEntry
var goLongExtensionEntryType = reflect.TypeOf(LongExtensionEntry{})

// Spec returns the specification object for LongExtensionEntry
func (o *LongExtensionEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.LongExtensionEntry",
		Type:       goLongExtensionEntryType,
	}
}

// LoadAllFields of LongExtensionEntry object
func (o *LongExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type StringExtensionEntry struct {
	BaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value string `jazz:"value"`
}

// StringExtensionEntryType contains the reflection type of StringExtensionEntry
var goStringExtensionEntryType = reflect.TypeOf(StringExtensionEntry{})

// Spec returns the specification object for StringExtensionEntry
func (o *StringExtensionEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.StringExtensionEntry",
		Type:       goStringExtensionEntryType,
	}
}

// LoadAllFields of StringExtensionEntry object
func (o *StringExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type MediumStringExtensionEntry struct {
	BaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value string `jazz:"value"`
}

// MediumStringExtensionEntryType contains the reflection type of MediumStringExtensionEntry
var goMediumStringExtensionEntryType = reflect.TypeOf(MediumStringExtensionEntry{})

// Spec returns the specification object for MediumStringExtensionEntry
func (o *MediumStringExtensionEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.MediumStringExtensionEntry",
		Type:       goMediumStringExtensionEntryType,
	}
}

// LoadAllFields of MediumStringExtensionEntry object
func (o *MediumStringExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type LargeStringExtensionEntry struct {
	BaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value string `jazz:"value"`
}

// LargeStringExtensionEntryType contains the reflection type of LargeStringExtensionEntry
var goLargeStringExtensionEntryType = reflect.TypeOf(LargeStringExtensionEntry{})

// Spec returns the specification object for LargeStringExtensionEntry
func (o *LargeStringExtensionEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.LargeStringExtensionEntry",
		Type:       goLargeStringExtensionEntryType,
	}
}

// LoadAllFields of LargeStringExtensionEntry object
func (o *LargeStringExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type TimestampExtensionEntry struct {
	BaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value *time.Time `jazz:"value"`
}

// TimestampExtensionEntryType contains the reflection type of TimestampExtensionEntry
var goTimestampExtensionEntryType = reflect.TypeOf(TimestampExtensionEntry{})

// Spec returns the specification object for TimestampExtensionEntry
func (o *TimestampExtensionEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.TimestampExtensionEntry",
		Type:       goTimestampExtensionEntryType,
	}
}

// LoadAllFields of TimestampExtensionEntry object
func (o *TimestampExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type BigDecimalExtensionEntry struct {
	BaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value float64 `jazz:"value"`
}

// BigDecimalExtensionEntryType contains the reflection type of BigDecimalExtensionEntry
var goBigDecimalExtensionEntryType = reflect.TypeOf(BigDecimalExtensionEntry{})

// Spec returns the specification object for BigDecimalExtensionEntry
func (o *BigDecimalExtensionEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.BigDecimalExtensionEntry",
		Type:       goBigDecimalExtensionEntryType,
	}
}

// LoadAllFields of BigDecimalExtensionEntry object
func (o *BigDecimalExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
	)
}

type ItemExtensionEntry struct {
	BaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value *Item `jazz:"value"`
}

// ItemExtensionEntryType contains the reflection type of ItemExtensionEntry
var goItemExtensionEntryType = reflect.TypeOf(ItemExtensionEntry{})

// Spec returns the specification object for ItemExtensionEntry
func (o *ItemExtensionEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.ItemExtensionEntry",
		Type:       goItemExtensionEntryType,
	}
}

// LoadAllFields of ItemExtensionEntry object
func (o *ItemExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Value,
	)
}

type MultiItemExtensionEntry struct {
	BaseObject

	// Key of the custom attribute
	Key string `jazz:"key"`

	// Value of the custom attribute
	Value []*Item `jazz:"value"`
}

// MultiItemExtensionEntryType contains the reflection type of MultiItemExtensionEntry
var goMultiItemExtensionEntryType = reflect.TypeOf(MultiItemExtensionEntry{})

// Spec returns the specification object for MultiItemExtensionEntry
func (o *MultiItemExtensionEntry) Spec() *ObjectSpec {
	return &ObjectSpec{
		ResourceID: "workitem",
		ElementID:  "",
		TypeID:     "com.ibm.team.repository.MultiItemExtensionEntry",
		Type:       goMultiItemExtensionEntryType,
	}
}

// LoadAllFields of MultiItemExtensionEntry object
func (o *MultiItemExtensionEntry) LoadAllFields() error {
	return o.loadFields(
		o.ModifiedBy,
		o.Value,
	)
}
