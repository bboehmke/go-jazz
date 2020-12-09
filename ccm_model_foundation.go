package jazz

type Contributor struct {
	BaseObject

	// https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#contributor
	_resource bool `jazz:"foundation"`
	_element  bool `jazz:"contributor"`
	_type     bool `jazz:"com.ibm.team.repository.Contributor"`

	// The human-readable name of the contributor (e.g. "James Moody")
	Name string `jazz:"name"`

	// The email address of the contributor
	EmailAddress string `jazz:"emailAddress"`

	// The userId of the contributor, unique in this application (e.g. "jmoody")
	UserId string `jazz:"userId"`
}
