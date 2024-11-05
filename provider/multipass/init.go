// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package multipass

import (
	"github.com/juju/juju/environs"
)

const (
	providerType = "multipass"
)

func init() {
	environs.RegisterProvider(providerType, NewProvider())
}
