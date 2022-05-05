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
	"time"
)

// resources: https://jazz.net/wiki/bin/view/Main/RqmApi#Resources_and_their_Supported_Op
// fields: https://jazz.net/wiki/bin/view/Main/RqmApi#fields

// QMAttachment implements the RQM "attachment" resource
type QMAttachment struct {
	QMBaseObject

	// Title of object
	Title string `json:"title"`

	// Numeric identifier shown in webinterface
	WebId int `json:"webId,string"`

	// FileSize of attachment
	FileSize int `json:"fileSize"`
}

// Spec returns the specification object for QMAttachment
func (o *QMAttachment) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "attachment",
	}
}

// QMTestEnvironment implements the RQM "configuration" resource
// (WebUI Name: "Test Environment")
type QMTestEnvironment struct {
	QMBaseObject

	// Summary of configuration
	Summary string `json:"summary"`
}

// Spec returns the specification object for QMTestEnvironment
func (o *QMTestEnvironment) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "configuration",
	}
}

// QMTestCase implements the RQM "testcase" resource
type QMTestCase struct {
	QMBaseObject

	// Title of object
	Title string `json:"title"`

	// Numeric identifier shown in webinterface
	WebId int `json:"webId,string"`

	// Description of object
	Description QMString `json:"description"`

	// TODO state

	// Owner of test case
	Owner QMUser `json:"owner"`

	// Creator of test case
	Creator QMUser `json:"creator"`

	// Updated contains last update time
	Updated time.Time `json:"updated"`

	// estimated execution time
	Estimate QMDuration `json:"estimate,string"`

	// Categories of test case
	Categories QMCategoryList `json:"category"`

	// AutomaticTestScriptRefs contains list of resource URLs for QMAutomaticTestScript
	AutomaticTestScriptRefs QMRefList `json:"remotescript"`

	// ManualTestScriptRefs contains list of resource URLs for QMManualTestScript
	ManualTestScriptRefs QMRefList `json:"testscript"`
}

// Spec returns the specification object for QMTestEnvironment
func (o *QMTestCase) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "testcase",
	}
}

// AutomaticTestScripts that are part of this QMTestCase
func (o *QMTestCase) AutomaticTestScripts() ([]*QMAutomaticTestScript, error) {
	return qmGetList[*QMAutomaticTestScript](o.proj, o.AutomaticTestScriptRefs.IDList())
}

// ManualTestScripts that are part of this QMTestCase
func (o *QMTestCase) ManualTestScripts() ([]*QMManualTestScript, error) {
	return qmGetList[*QMManualTestScript](o.proj, o.ManualTestScriptRefs.IDList())
}

// QMManualTestScript implements the RQM "testscript" resource
type QMManualTestScript struct {
	QMBaseObject

	// Title of object
	Title string `json:"title"`

	// Numeric identifier shown in webinterface
	WebId int `json:"webId,string"`

	// Description of object
	Description QMString `json:"description"`

	// TODO state

	// Owner of test script
	Owner QMUser `json:"owner"`

	// Creator of test script
	Creator QMUser `json:"creator"`

	// Updated contains last update time
	Updated time.Time `json:"updated"`
}

// Spec returns the specification object for QMManualTestScript
func (o *QMManualTestScript) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "testscript",
	}
}

// QMAutomaticTestScript implements the RQM "remotescript" resource
type QMAutomaticTestScript struct {
	QMBaseObject

	// Title of object
	Title string `json:"title"`

	// Numeric identifier shown in webinterface
	WebId int `json:"webId,string"`

	// Description of object
	Description QMString `json:"description"`

	// TODO state

	// Owner of test case
	Owner QMUser `json:"owner"`

	// Creator of test case
	Creator QMUser `json:"creator"`

	// Updated contains last update time
	Updated time.Time `json:"updated"`

	// Command for automatic test script
	Command string `json:"command"`

	// Arguments for automatic test script
	Arguments string `json:"arguments"`
}

// Spec returns the specification object for QMManualTestScript
func (o *QMAutomaticTestScript) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "remotescript",
	}
}

// QMTestExecutionRecord implements the RQM "executionworkitem" resource
type QMTestExecutionRecord struct {
	QMBaseObject

	// Title of object
	Title string `json:"title"`

	// Numeric identifier shown in webinterface
	WebId int `json:"webId,string"`

	// Description of object
	Description QMString `json:"description"`

	// TODO state

	// estimated execution time
	Estimate QMDuration `json:"estimate,string"`

	// Owner of test case
	Owner QMUser `json:"owner"`

	// Creator of test case
	Creator QMUser `json:"creator"`

	// Updated contains last update time
	Updated time.Time `json:"updated"`

	// TestCaseRef contains reference to last execution QMTestCase
	TestCaseRef QMRef `json:"testcase"`

	// TestEnvironmentRef contains reference to last execution QMTestEnvironment
	TestEnvironmentRef QMRef `json:"configuration"`

	// LastExecutionResultRef contains reference to last execution QMTestExecutionResult
	LastExecutionResultRef QMRef `json:"currentexecutionresult"`

	// TestExecutionResults contains list of resource URLs for QMTestExecutionResult
	TestExecutionResults QMRefList `json:"executionresult"`
}

// Spec returns the specification object for QMManualTestScript
func (o *QMTestExecutionRecord) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "executionworkitem",
	}
}

// TestCase of this QMTestExecutionRecord
func (o *QMTestExecutionRecord) TestCase() (*QMTestCase, error) {
	return QMGet[*QMTestCase](o.proj, o.TestCaseRef.Href)
}

// TestEnvironment of this QMTestExecutionRecord
func (o *QMTestExecutionRecord) TestEnvironment() (*QMTestEnvironment, error) {
	return QMGet[*QMTestEnvironment](o.proj, o.TestEnvironmentRef.Href)
}

// QMTestExecutionResult implements the RQM "executionresult" resource
type QMTestExecutionResult struct {
	QMBaseObject

	// Title of object
	Title string `json:"title"`

	// Numeric identifier shown in webinterface
	WebId int `json:"webId,string" jazz:"qm:webId"`

	// State of test execution
	State string `json:"state" jazz:"alm:state"`

	// Creator of entry
	Creator QMUser `json:"creator"`

	// Updated contains last update time
	Updated time.Time `json:"updated"`

	// Machine of where test was executed
	Machine string `json:"machine" jazz:"qmresult:machine"`

	// StartTime of test execution
	StartTime time.Time `json:"starttime" jazz:"qmresult:starttime"`

	// EndTime of test execution
	EndTime time.Time `json:"endtime" jazz:"qmresult:endtime"`

	// Variables of test execution result
	Variables QMVariableMap `json:"variables" jazz:"qm:variables"`

	// TestCaseRef contains reference to last execution QMTestCase
	TestCaseRef QMRef `json:"testcase" jazz:"qm:testcase"`

	// TestEnvironmentRef contains reference to last execution QMTestEnvironment
	TestEnvironmentRef QMRef `json:"configuration" jazz:"qm:configuration"`

	// TestExecutionRecordRef contains reference to last execution QMTestExecutionRecord
	TestExecutionRecordRef QMRef `json:"executionworkitem" jazz:"qm:executionworkitem"`

	// AutomaticTestScriptRef contains reference to last execution QMAutomaticTestScript
	AutomaticTestScriptRef QMRef `json:"remotescript" jazz:"qm:remotescript"`

	// ManualTestScriptRef contains reference to last execution QMManualTestScript
	ManualTestScriptRef QMRef `json:"testscript" jazz:"qm:testscript"`
}

// Spec returns the specification object for QMManualTestScript
func (o *QMTestExecutionResult) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "executionresult",
	}
}

// TestCase of this QMTestExecutionResult
func (o *QMTestExecutionResult) TestCase() (*QMTestCase, error) {
	return QMGet[*QMTestCase](o.proj, o.TestCaseRef.Href)
}

// TestEnvironment of this QMTestExecutionResult
func (o *QMTestExecutionResult) TestEnvironment() (*QMTestEnvironment, error) {
	return QMGet[*QMTestEnvironment](o.proj, o.TestEnvironmentRef.Href)
}

// TestExecutionRecord of this QMTestExecutionResult
func (o *QMTestExecutionResult) TestExecutionRecord() (*QMTestExecutionRecord, error) {
	return QMGet[*QMTestExecutionRecord](o.proj, o.TestExecutionRecordRef.Href)
}

// AutomaticTestScript of this QMTestExecutionResult
func (o *QMTestExecutionResult) AutomaticTestScript() (*QMAutomaticTestScript, error) {
	return QMGet[*QMAutomaticTestScript](o.proj, o.AutomaticTestScriptRef.Href)
}

// ManualTestScript of this QMTestExecutionResult
func (o *QMTestExecutionResult) ManualTestScript() (*QMManualTestScript, error) {
	return QMGet[*QMManualTestScript](o.proj, o.ManualTestScriptRef.Href)
}

// QMTestPlan implements the RQM "testplan" resource
type QMTestPlan struct {
	QMBaseObject

	// Title of object
	Title string `json:"title"`

	// Alias of object (used in resource URL)
	Alias string `json:"alias"`

	// Numeric identifier shown in webinterface
	WebId int `json:"webId,string"`

	// Description of object
	Description QMString `json:"description"`

	// TestEnvironmentRefs contains list of resource URLs for QMTestEnvironment
	TestEnvironmentRefs QMRefList `json:"configuration"`

	// TestCaseRefs contains list of resource URLs for QMTestCase
	TestCaseRefs QMRefList `json:"testcase"`
}

// Spec returns the specification object for QMTestPlan
func (o *QMTestPlan) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "testplan",
	}
}

// TestEnvironments that are part of this QMTestPlan
func (o *QMTestPlan) TestEnvironments() ([]*QMTestEnvironment, error) {
	return qmGetList[*QMTestEnvironment](o.proj, o.TestEnvironmentRefs.IDList())
}

// TestExecutionRecords that are part of this QMTestPlan
func (o *QMTestPlan) TestExecutionRecords() ([]*QMTestExecutionRecord, error) {
	return QMList[*QMTestExecutionRecord](o.proj, map[string]string{
		"testplan/@href": o.ResourceUrl,
	})
}

// TestExecutionResults that are part of this QMTestPlan
func (o *QMTestPlan) TestExecutionResults() ([]*QMTestExecutionResult, error) {
	return QMList[*QMTestExecutionResult](o.proj, map[string]string{
		"testplan/@href": o.ResourceUrl,
	})
}
