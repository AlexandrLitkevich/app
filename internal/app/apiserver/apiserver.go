package apiserver

// APIServer API Server ...
type APIServer struct {
}

func New() *APIServer {
	// Initial APIServer
	return &APIServer{}
}

// Start ...
func (s *APIServer) Start() error {
	return nil
}
