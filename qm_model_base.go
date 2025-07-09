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
	"encoding/json"
	"encoding/xml"
	"time"
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
	Ref() QMRef
	SetRef(url string)
}

// QMBaseObject for RQM resources
type QMBaseObject struct {
	// ResourceUrl of object (used as "identifier")
	ResourceUrl string `json:"identifier" xml:"identifier"`

	// QMProject instance used for interactions with the server
	proj *QMProject
}

// SetProj of object
func (o *QMBaseObject) SetProj(proj *QMProject) {
	o.proj = proj
}

// Ref returns QMRef of object
func (o QMBaseObject) Ref() QMRef {
	return QMRef{
		Href: o.ResourceUrl,
	}
}

// SetRef URL of object
func (o *QMBaseObject) SetRef(url string) {
	o.ResourceUrl = url
}

// QMRef reference to object
type QMRef struct {
	Href string `json:"href" xml:"href,attr"`
}

func (s QMRef) String() string {
	return s.Href
}

// QMRefList list of object references
type QMRefList []QMRef

func (s QMRefList) IDList() []string {
	ids := make([]string, len(s))
	for i, ref := range s {
		ids[i] = ref.Href
	}
	return ids
}

// QMCategory entry of test case
type QMCategory struct {
	Name  string `xml:"term,attr"`
	Value string `xml:"value,attr"`
}

// QMDuration used in QM objects (stored as milliseconds)
type QMDuration time.Duration

func (d *QMDuration) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var s string
	err := decoder.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	duration, err := time.ParseDuration(s + "ms")
	if err != nil {
		return err
	}
	*d = QMDuration(duration)
	return nil
}

// QMVariableMap contains list of variables
type QMVariableMap map[string]string

func (m *QMVariableMap) UnmarshalJSON(b []byte) error {
	var buffer struct {
		Variables []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"variable"`
	}
	err := json.Unmarshal(b, &buffer)
	if err == nil {
		*m = make(map[string]string, len(buffer.Variables))
		for _, v := range buffer.Variables {
			(*m)[v.Name] = v.Value
		}
	}
	return err
}

func (m *QMVariableMap) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var buffer struct {
		Variables []struct {
			Name  string `xml:"name"`
			Value string `xml:"value"`
		} `xml:"variable"`
	}
	err := decoder.DecodeElement(&buffer, &start)
	if err != nil {
		return err
	}

	*m = make(map[string]string, len(buffer.Variables))
	for _, variable := range buffer.Variables {
		(*m)[variable.Name] = variable.Value
	}

	return nil
}

// QMXmlText is a custom type to handle XML text content
type QMXmlText string

func (t *QMXmlText) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var buffer struct {
		Data string `xml:",innerxml"`
	}
	err := decoder.DecodeElement(&buffer, &start)
	if err != nil {
		return err
	}

	*t = QMXmlText(buffer.Data)

	return nil
}
