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
	"fmt"
	"io"
	"time"
)

// resources: https://jazz.net/wiki/bin/view/Main/RqmApi#Resources_and_their_Supported_Op
// fields: https://jazz.net/wiki/bin/view/Main/RqmApi#fields

// QMAttachment implements the RQM "attachment" resource
type QMAttachment struct {
	QMBaseObject

	// Title of object
	Title string `xml:"title"`

	// Numeric identifier shown in webinterface
	WebId int `xml:"webId"`

	// FileSize of attachment
	FileSize float64 `xml:"fileSize"`
}

// Spec returns the specification object for QMAttachment
func (o *QMAttachment) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "attachment",
	}
}

// Download content of attachment
func (o *QMAttachment) Download(ctx context.Context, w io.Writer) error {
	// copy attachment content
	response, err := o.proj.qm.client.get(ctx, o.ResourceUrl, "application/octet-stream", false)
	if err != nil {
		return fmt.Errorf("failed to get attachment: %w", err)
	}
	defer response.Body.Close()

	// copy attachment content
	_, err = io.Copy(w, response.Body)
	if err != nil {
		return fmt.Errorf("failed to get attachment: %w", err)
	}
	return nil
}

// QMTestEnvironment implements the RQM "configuration" resource
// (WebUI Name: "Test Environment")
type QMTestEnvironment struct {
	QMBaseObject

	// Title of object
	Title string `xml:"title"`

	// Summary of configuration
	Summary string `xml:"summary"`
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
	Title string `xml:"title"`

	// Numeric identifier shown in webinterface
	WebId int `xml:"webId"`

	// Description of object
	Description string `xml:"description"`

	// TODO state

	// Owner of test case
	Owner string `xml:"owner"`

	// Creator of test case
	Creator string `xml:"creator"`

	// Updated contains last update time
	Updated time.Time `xml:"updated"`

	// estimated execution time
	Estimate QMDuration `xml:"estimate"`

	// Categories of test case
	Categories []QMCategory `xml:"category"`

	// AutomaticTestScriptRefs contains list of resource URLs for QMAutomaticTestScript
	AutomaticTestScriptRefs QMRefList `xml:"remotescript"`

	// ManualTestScriptRefs contains list of resource URLs for QMManualTestScript
	ManualTestScriptRefs QMRefList `xml:"testscript"`
}

// Spec returns the specification object for QMTestEnvironment
func (o *QMTestCase) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "testcase",
	}
}

// AutomaticTestScripts that are part of this QMTestCase
func (o *QMTestCase) AutomaticTestScripts(ctx context.Context) ([]*QMAutomaticTestScript, error) {
	return qmGetList[*QMAutomaticTestScript](ctx, o.proj, o.AutomaticTestScriptRefs.IDList())
}

// ManualTestScripts that are part of this QMTestCase
func (o *QMTestCase) ManualTestScripts(ctx context.Context) ([]*QMManualTestScript, error) {
	return qmGetList[*QMManualTestScript](ctx, o.proj, o.ManualTestScriptRefs.IDList())
}

// QMManualTestScript implements the RQM "testscript" resource
type QMManualTestScript struct {
	QMBaseObject

	// Title of object
	Title string `xml:"title"`

	// Numeric identifier shown in webinterface
	WebId int `xml:"webId"`

	// Description of object
	Description string `xml:"description"`

	// TODO state

	// Owner of test script
	Owner string `xml:"owner"`

	// Creator of test script
	Creator string `xml:"creator"`

	// Updated contains last update time
	Updated time.Time `xml:"updated"`
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
	Title string `xml:"title"`

	// Numeric identifier shown in webinterface
	WebId int `xml:"webId"`

	// Description of object
	Description string `xml:"description"`

	// TODO state

	// Owner of test case
	Owner string `xml:"owner"`

	// Creator of test case
	Creator string `xml:"creator"`

	// Updated contains last update time
	Updated time.Time `xml:"updated"`

	// Command for automatic test script
	Command string `xml:"command"`

	// Arguments for automatic test script
	Arguments string `xml:"arguments"`
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
	Title string `xml:"title"`

	// Numeric identifier shown in webinterface
	WebId int `xml:"webId"`

	// Description of object
	Description string `xml:"description"`

	// TODO state

	// estimated execution time
	Estimate QMDuration `xml:"estimate"`

	// Owner of test case
	Owner string `xml:"owner"`

	// Creator of test case
	Creator string `xml:"creator"`

	// Updated contains last update time
	Updated time.Time `xml:"updated"`

	// TestCaseRef contains reference to last execution QMTestCase
	TestCaseRef QMRef `xml:"testcase"`

	// TestEnvironmentRef contains reference to last execution QMTestEnvironment
	TestEnvironmentRef QMRef `xml:"configuration"`

	// LastExecutionResultRef contains reference to last execution QMTestExecutionResult
	LastExecutionResultRef QMRef `xml:"currentexecutionresult"`

	// TestExecutionResults contains list of resource URLs for QMTestExecutionResult
	TestExecutionResults QMRefList `xml:"executionresult"`
}

// Spec returns the specification object for QMManualTestScript
func (o *QMTestExecutionRecord) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "executionworkitem",
	}
}

// TestCase of this QMTestExecutionRecord
func (o *QMTestExecutionRecord) TestCase(ctx context.Context) (*QMTestCase, error) {
	return QMGet[*QMTestCase](ctx, o.proj, o.TestCaseRef.Href)
}

// TestEnvironment of this QMTestExecutionRecord
func (o *QMTestExecutionRecord) TestEnvironment(ctx context.Context) (*QMTestEnvironment, error) {
	return QMGet[*QMTestEnvironment](ctx, o.proj, o.TestEnvironmentRef.Href)
}

// QMTestExecutionResult implements the RQM "executionresult" resource
type QMTestExecutionResult struct {
	QMBaseObject

	// Title of object
	Title string `xml:"title"`

	// Numeric identifier shown in webinterface
	WebId int `xml:"webId" jazz:"qm:webId"`

	// State of test execution
	State string `xml:"state" jazz:"alm:state"`

	// Creator of entry
	Creator string `xml:"creator"`

	// Updated contains last update time
	Updated time.Time `xml:"updated"`

	// Machine of where test was executed
	Machine string `xml:"machine" jazz:"qmresult:machine"`

	// StartTime of test execution
	StartTime time.Time `xml:"starttime" jazz:"qmresult:starttime"`

	// EndTime of test execution
	EndTime time.Time `xml:"endtime" jazz:"qmresult:endtime"`

	// Variables of test execution result
	Variables QMVariableMap `xml:"variables" jazz:"qm:variables"`

	// TestCaseRef contains reference to last execution QMTestCase
	TestCaseRef QMRef `xml:"testcase" jazz:"qm:testcase"`

	// TestEnvironmentRef contains reference to last execution QMTestEnvironment
	TestEnvironmentRef QMRef `xml:"configuration" jazz:"qm:configuration"`

	// TestExecutionRecordRef contains reference to last execution QMTestExecutionRecord
	TestExecutionRecordRef QMRef `xml:"executionworkitem" jazz:"qm:executionworkitem"`

	// AutomaticTestScriptRef contains reference to last execution QMAutomaticTestScript
	AutomaticTestScriptRef QMRef `xml:"remotescript" jazz:"qm:remotescript"`

	// ManualTestScriptRef contains reference to last execution QMManualTestScript
	ManualTestScriptRef QMRef `xml:"testscript" jazz:"qm:testscript"`
}

// Spec returns the specification object for QMManualTestScript
func (o *QMTestExecutionResult) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "executionresult",
	}
}

// TestCase of this QMTestExecutionResult
func (o *QMTestExecutionResult) TestCase(ctx context.Context) (*QMTestCase, error) {
	return QMGet[*QMTestCase](ctx, o.proj, o.TestCaseRef.Href)
}

// TestEnvironment of this QMTestExecutionResult
func (o *QMTestExecutionResult) TestEnvironment(ctx context.Context) (*QMTestEnvironment, error) {
	return QMGet[*QMTestEnvironment](ctx, o.proj, o.TestEnvironmentRef.Href)
}

// TestExecutionRecord of this QMTestExecutionResult
func (o *QMTestExecutionResult) TestExecutionRecord(ctx context.Context) (*QMTestExecutionRecord, error) {
	return QMGet[*QMTestExecutionRecord](ctx, o.proj, o.TestExecutionRecordRef.Href)
}

// AutomaticTestScript of this QMTestExecutionResult
func (o *QMTestExecutionResult) AutomaticTestScript(ctx context.Context) (*QMAutomaticTestScript, error) {
	return QMGet[*QMAutomaticTestScript](ctx, o.proj, o.AutomaticTestScriptRef.Href)
}

// ManualTestScript of this QMTestExecutionResult
func (o *QMTestExecutionResult) ManualTestScript(ctx context.Context) (*QMManualTestScript, error) {
	return QMGet[*QMManualTestScript](ctx, o.proj, o.ManualTestScriptRef.Href)
}

// QMTestPlan implements the RQM "testplan" resource
type QMTestPlan struct {
	QMBaseObject

	// Title of object
	Title string `xml:"title"`

	// Alias of object (used in resource URL)
	Alias string `xml:"alias"`

	// Numeric identifier shown in webinterface
	WebId int `xml:"webId"`

	// Description of object
	Description string `xml:"description"`

	// TestEnvironmentRefs contains list of resource URLs for QMTestEnvironment
	TestEnvironmentRefs QMRefList `xml:"configuration"`

	// TestCaseRefs contains list of resource URLs for QMTestCase
	TestCaseRefs QMRefList `xml:"testcase"`
}

// Spec returns the specification object for QMTestPlan
func (o *QMTestPlan) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "testplan",
	}
}

// TestEnvironments that are part of this QMTestPlan
func (o *QMTestPlan) TestEnvironments(ctx context.Context) ([]*QMTestEnvironment, error) {
	return qmGetList[*QMTestEnvironment](ctx, o.proj, o.TestEnvironmentRefs.IDList())
}

// TestExecutionRecords that are part of this QMTestPlan
func (o *QMTestPlan) TestExecutionRecords(ctx context.Context) ([]*QMTestExecutionRecord, error) {
	return QMList[*QMTestExecutionRecord](ctx, o.proj, map[string]string{
		"testplan/@href": o.ResourceUrl,
	})
}

// TestExecutionResults that are part of this QMTestPlan
func (o *QMTestPlan) TestExecutionResults(ctx context.Context) ([]*QMTestExecutionResult, error) {
	return QMList[*QMTestExecutionResult](ctx, o.proj, map[string]string{
		"testplan/@href": o.ResourceUrl,
	})
}
