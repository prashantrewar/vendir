// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package version

import (
	"runtime/debug"
)

var (
	// Version can be set via:
	// -ldflags="-X 'carvel.dev/vendir/pkg/vendir/version.Version=$TAG'"
	defaultVersion = "0.0.0+develop"
	Version        = ""
	moduleName     = "carvel.dev/vendir"
)

func init() {
	Version = version()
}

func version() string {
	if Version != "" {
		// Version was set via ldflags, just return it.
		return Version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return defaultVersion
	}

	// Anything else.
	for _, dep := range info.Deps {
		if dep.Path == moduleName {
			return dep.Version
		}
	}

	return defaultVersion
}
