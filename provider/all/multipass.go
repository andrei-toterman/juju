// Copyright Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

//go:build !minimal || provider_multipass

package all

import (
	// Register the provider.
	_ "github.com/juju/juju/provider/multipass"
)
