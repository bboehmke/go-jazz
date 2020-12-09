package jazz

type Application interface {
	Name() string
	ID() string
	Client() *Client
}

type App struct {
	Application
}

func (a *App) RootServices() *RootService {
	return &RootService{
		client: a.Client(),
		base:   a.ID(),
	}
}
