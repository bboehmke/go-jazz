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

package main

// list of types missing in documentation
var missingTypes = map[string]string{
	"contributor": "com.ibm.team.repository.Contributor",
}

// mapping of invalid types to correct ones
var invalidTypes = map[string]string{
	"com.ibm.team.repository.Attribute": "com.ibm.team.workitem.Attribute",
	"com.ibm.team.workitem.Contributor": "com.ibm.team.repository.Contributor",
	"com.ibm.team.workitem.Approvals":   "com.ibm.team.workitem.Approval",
	"com.ibm.workitem.Deliverable":      "com.ibm.team.workitem.Deliverable",
}

// fields to ignore on objects
var skipFields = map[string]map[string]struct{}{
	"com.ibm.team.workitem.Approval": {
		"approvalDescriptor": {}, // causes infinite recursion
	},
	"com.ibm.team.build.BuildResult": {
		"compilationResults": {}, // invalid field?
	},
}

// mapping of invalid element IDs to correct ones
var invalidElementIDs = map[string]string{
	"com.ibm.team.workitem.ExtensionEntry": "",
}
