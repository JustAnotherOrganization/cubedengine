package minecraft

import (
	"context"
	"os"
	"os/exec"
)

// Server holds information about the running process and the configuration.
type Server struct {
	command *exec.Cmd
	process *os.Process
}

// Start will start the minecraft server
func (server *Server) Start() error {
	if err := server.command.Start(); err != nil {
		return err
	}
	server.process = server.command.Process

	return nil
}

// Stop will stop the minecraft server.
func (server *Server) Stop(ctx context.Context) {

}
