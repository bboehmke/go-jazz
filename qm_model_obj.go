package jazz

import "reflect"

// resources: https://jazz.net/wiki/bin/view/Main/RqmApi#Resources_and_their_Supported_Op
// fields: https://jazz.net/wiki/bin/view/Main/RqmApi#fields

// QMTestEnvironment implements the RQM "configuration" resource
// (WebUI Name: "Test Environment")
type QMTestEnvironment struct {
	QMBaseObject

	// Summary of configuration
	Summary string `json:"summary"`
}

// goQMTestEnvironment contains the reflection type of QMTestEnvironment
var goQMTestEnvironment = reflect.TypeOf(QMTestEnvironment{})

// Spec returns the specification object for QMTestEnvironment
func (o *QMTestEnvironment) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "configuration",
		Type:       goQMTestEnvironment,
	}
}

// QMTestCase implements the RQM "testcase" resource
type QMTestCase struct {
	QMBaseObject

	// Alias of object (used in resource URL)
	Alias string `json:"alias"`

	// Numeric identifier shown in webinterface
	WebId int `json:"webId,string"`

	// Description of object
	Description QMString `json:"description"`
}

// goQMTestCaseType contains the reflection type of QMTestCase
var goQMTestCaseType = reflect.TypeOf(QMTestCase{})

// Spec returns the specification object for QMTestEnvironment
func (o *QMTestCase) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "testcase",
		Type:       goQMTestCaseType,
	}
}

// QMTestPlan implements the RQM "testplan" resource
type QMTestPlan struct {
	QMBaseObject

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

// goQMTestPlanType contains the reflection type of QMTestPlan
var goQMTestPlanType = reflect.TypeOf(QMTestPlan{})

// Spec returns the specification object for QMTestPlan
func (o *QMTestPlan) Spec() *QMObjectSpec {
	return &QMObjectSpec{
		ResourceID: "testplan",
		Type:       goQMTestPlanType,
	}
}

// TestEnvironments that are part of this QMTestPlan
func (o *QMTestPlan) TestEnvironments() ([]*QMTestEnvironment, error) {
	return QMGetList[*QMTestEnvironment](o.proj, o.TestEnvironmentRefs)
}

// TestCases that are part of this QMTestPlan
func (o *QMTestPlan) TestCases() ([]*QMTestCase, error) {
	return QMGetList[*QMTestCase](o.proj, o.TestCaseRefs)
}
