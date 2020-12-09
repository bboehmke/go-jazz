package jazz

type CCMApplication struct {
	client *Client
}

func (a *CCMApplication) Name() string {
	return "Change and Configuration Management"
}

func (a *CCMApplication) ID() string {
	return "ccm"
}

func (a *CCMApplication) Client() *Client {
	return a.Client()
}
