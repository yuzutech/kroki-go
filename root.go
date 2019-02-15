package kroki

// Client is a Kroki client
type Client struct {
	Config Configuration
}

// New creates a new Client
func New(config Configuration) Client {
	client := Client{
		Config: config,
	}
	return client
}
