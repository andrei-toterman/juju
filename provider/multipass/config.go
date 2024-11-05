package multipass

import "github.com/juju/juju/environs/config"

type environConfig struct {
	*config.Config
	attrs map[string]interface{}
}
