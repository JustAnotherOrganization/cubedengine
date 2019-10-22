package network

import "justanother.org/cubedengine/internal/config"

// StartHTTPServer starts the http server, but does not block. It does wait a
// few seconds to ensure that it is able to make the connections
func StartHTTPServer() error {
	config := config.Get()

	// Ensure that the network is not disabled.
	if config.Network.Enabled != nil && !*config.Network.Enabled {
		return nil
	}

	return nil
}
