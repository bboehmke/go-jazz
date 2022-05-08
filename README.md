[![Go Reference](https://pkg.go.dev/badge/gitlab.com/bboehmke/go-jazz.svg)](https://pkg.go.dev/gitlab.com/bboehmke/go-jazz)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/bboehmke/go-jazz)](https://goreportcard.com/report/gitlab.com/bboehmke/go-jazz)

# go-jazz

This package provides access to the API of the 
[IBM Engineering Lifecycle Management](https://www.ibm.com/products/engineering-lifecycle-management).

## Features

* generic client that handles multiple [authentication methods](https://jazz.net/wiki/bin/view/Main/NativeClientAuthentication)
  * Form Challenge
  * Basic Auth
* support for requests under a global configuration
* interface to RTC SCM credentials (see [Credential helper](#credential-helper))
* support of multiple jazz applications:
  * CCM:
    * generated code with all available object types
    * read only
  * QM:
    * only some objects implemented
    * modification of some object implemented
    * upload of attachments
  * GC:
    * only minimal list and get implementation


## Usage

> Note: At least [Go 1.18](https://tip.golang.org/doc/go1.18) is required as 
> this package is using generics.

```go
import "gitlab.com/bboehmke/go-jazz"
```

Construct a new Jazz client, then access one of the applications. 
For example, to list all work items available in CCM:

```go
client, err := jazz.NewClient("https://jazz.example.com/", "user", "password")
if err != nil {
    panic(err)
}

workItems, err := jazz.CCMList[*jazz.CCMWorkItem](context.TODO(), client.CCM, nil)
```
> Note: Use the generic URL without the `ccm`, `qm` or `gc` suffix!

### CCM Application

The CCM interface is build based on the description of the 
[Reportable REST API](https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI).

The objects are directly generated from the documentation (see `cmd/ccm_model_generator`).

There are two request types:
1. `CCMList`, `CCMListChan`: returns a list of objects
2. `CCMGet`, `CCMGetFilter`: returns only one object

For example to get all work items created by a specific user we can do the following:
```go
// first get the user with the user ID "user"
user, err := jazz.CCMGetFilter[*jazz.CCMContributor](context.TODO(), client.CCM, jazz.CCMFilter{
    "UserId": []interface{}{"user"},
})
if err != nil {
    panic(err)
}

// then query the work items
workItems, err := jazz.CCMList[*jazz.CCMWorkItem](context.TODO(), client.CCM, jazz.CCMFilter{
    "Creator": []interface{}{user},
})
if err != nil {
    panic(err)
}
```

### QM Application

The QM interface is build based on the description of the
[Reportable  REST API](https://jazz.net/wiki/bin/view/Main/RqmApi).

The interface only implements a small subset of available objects and values
provided by the API.

There are 3 request types:
1. `QMList`, `QMListChan`: returns a list of objects
2. `QMGet`, `QMGetFilter`: returns only one object
3. `QMSave`: is used to modify an object (only supported for some objects)

> Note: currently only some fields and objects are supported for write operations

Each action against the QM API are related to a project so this has to be 
get first. All following action are executed against this project.

For example to get all execution records of a test plan:
```go
// get the project
project, err := client.QM.GetProject(context.TODO(), "Project Title")
if err != nil {
    panic(err)
}

// get a test plan by its name/title
testPlan, err := jazz.QMGetFilter[*jazz.QMTestPlan](context.TODO(), project, jazz.QMFilter{
    "title": "TestPlan Title",
})
if err != nil {
    panic(err)
}

// get execution records from test plan
executionRecords, err := QMList[*QMTestExecutionRecord](context.TODO(), project, map[string]string{
    "testplan/@href": testPlan.ResourceUrl,
})
if err != nil {
    panic(err)
}
```

For some relation the objects contains methods to query for related objects:
```go
// get execution records from test plan
executionRecords, err := testPlan.TestExecutionRecords(context.TODO())
if err != nil {
    panic(err)
}
```

### Credential Helper

Instead of providing the password directly it is also possible to reuse the 
password stored in the eclipse password store (which is used by teh SCM desktop client):

```go
password, err := jazz.ReadEclipsePassword("https://jazz.example.com/", "user")
if err != nil {
    panic(err)
}
```

> Note: this is currently only possible on Windows.

# API Reference

This implementation is mainly based on these resources:

* Auth: https://jazz.net/wiki/bin/view/Main/NativeClientAuthentication
* CCM: https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI
* QM: https://jazz.net/wiki/bin/view/Main/RqmApi
* GC: https://jazz.net/sandbox02-gc/doc/scenarios

Also, some information were collected based on API responses and other information
available in the jazz.net forums:
* missing CCM type (see [ccm_model_generator](cmd/ccm_model_generator/missing.go))
