package jazz

import (
	"encoding/json"
	"strings"
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

// QMRefList list of object references
type QMRefList []string

func (s *QMRefList) UnmarshalJSON(b []byte) error {
	type ref struct {
		Href string `json:"href"`
	}

	var refs []ref
	err := json.Unmarshal(b, &refs)
	if err == nil {
		for _, r := range refs {
			*s = append(*s, r.Href)
		}
	} else if _, ok := err.(*json.UnmarshalTypeError); ok {
		var r ref
		err = json.Unmarshal(b, &r)
		if err == nil {
			*s = append(*s, r.Href)
		}
	}
	return err
}
