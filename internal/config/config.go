package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Settings holds onto the settings that our application uses. Each field maps
// to an environmental variable, and when I get around to it a flag.
type Settings struct {
	Network Network
}

// Network holds onto the network settings
type Network struct {
	TLSCert string
	TLSKey  string
	Enabled *bool
	Host    string
	Port    int
}

var defaultConfig Settings

// Get returns the default settings used by the application
func Get() Settings { return defaultConfig }

func init() {
	if err := envconfig.Process("cubed", &defaultConfig); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read in the env settings: %v\n", err)
		os.Exit(1)
	}
}
