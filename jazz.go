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

// TODO GC -> https://jazz.net/sandbox02-gc/doc/scenarios
