package jackett

// Client is the base client struct
type Client struct {
	Apikey   string
	Endpoint string
}

// New returns a pointer to jackett client
func New(apikey, endpoint string) *Client {
	return &Client{Apikey: apikey, Endpoint: endpoint}
}
