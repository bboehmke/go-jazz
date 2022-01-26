package jazz

// Contributor: https://jazz.net/wiki/bin/view/Main/ReportsRESTAPI#contributor
type Contributor struct {
	BaseObject `jazz:"foundation,com.ibm.team.repository.Contributor,contributor"`

	// The human-readable name of the contributor (e.g. "James Moody")
	Name string `jazz:"name"`

	// The email address of the contributor
	EmailAddress string `jazz:"emailAddress"`

	// The userId of the contributor, unique in this application (e.g. "jmoody")
	UserId string `jazz:"userId"`
}
