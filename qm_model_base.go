package jazz

import (
	"strings"
)

const (
	QMResultStatePaused       = "com.ibm.rqm.execution.common.state.paused"
	QMResultStateInProgress   = "com.ibm.rqm.execution.common.state.inprogress"
	QMResultStateNotRun       = "com.ibm.rqm.execution.common.state.notrun"
	QMResultStatePassed       = "com.ibm.rqm.execution.common.state.passed"
	QMResultStatePermFailed   = "com.ibm.rqm.execution.common.state.perm_failed"
	QMResultStateIncomplete   = "com.ibm.rqm.execution.common.state.incomplete"
	QMResultStateInconclusive = "com.ibm.rqm.execution.common.state.inconclusive"
	QMResultStatePartBlocked  = "com.ibm.rqm.execution.common.state.part_blocked"
	QMResultStateDeferred     = "com.ibm.rqm.execution.common.state.deferred"
	QMResultStateFailed       = "com.ibm.rqm.execution.common.state.failed"
	QMResultStateError        = "com.ibm.rqm.execution.common.state.error"
	QMResultStateBlocked      = "com.ibm.rqm.execution.common.state.blocked"
)

// QMObject describes a QM object implementation
type QMObject interface {
	Spec() *QMObjectSpec
	SetProj(proj *QMProject)
}

// QMBaseObject for RQM resources
type QMBaseObject struct {
	// ResourceUrl of object (used as "identifier")
	ResourceUrl string `json:"identifier"`

	// Title of object
	Title string `json:"title"`

	// QMProject instance used for interactions with the server
	proj *QMProject
}

// SetProj of object
func (o *QMBaseObject) SetProj(proj *QMProject) {
	o.proj = proj
}

// QMString handles json marshalling of broken values
type QMString string

func (s *QMString) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")
	if str != "true" {
		*s = QMString(str)
	}
	return nil
}

func (s QMString) String() string {
	return string(s)
}

// QMRef reference to object
type QMRef struct {
	Href string `json:"href"`
}

func (s QMRef) String() string {
	return s.Href
}

// QMRefList list of object references
type QMRefList []QMRef

func (s *QMRefList) UnmarshalJSON(b []byte) error {
	entries, err := UnmarshalJSONOptionalList[QMRef](b)
	if err == nil {
		*s = entries
	}
	return err
}

func (s QMRefList) IDList() []string {
	ids := make([]string, len(s))
	for i, ref := range s {
		ids[i] = ref.Href
	}
	return ids
}

// QMUser provides name and ID of user
type QMUser struct {
	Id   string `json:"content"`
	Name string `json:"name"`
}

// QMCategory entry of test case
type QMCategory struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// QMCategoryList contains list of categories
type QMCategoryList []QMCategory

func (l *QMCategoryList) UnmarshalJSON(b []byte) error {
	entries, err := UnmarshalJSONOptionalList[QMCategory](b)
	if err == nil {
		*l = entries
	}
	return err
}
