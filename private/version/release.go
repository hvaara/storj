// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package version

import _ "unsafe" // needed for go:linkname

//go:linkname buildTimestamp storj.io/private/version.buildTimestamp
var buildTimestamp string = "1588775602"

//go:linkname buildCommitHash storj.io/private/version.buildCommitHash
var buildCommitHash string = "80e17fc545bd5ac65c823ea724c5ab6f882ce32b"

//go:linkname buildVersion storj.io/private/version.buildVersion
var buildVersion string = "v1.4.1"

//go:linkname buildRelease storj.io/private/version.buildRelease
var buildRelease string = "true"

// ensure that linter understands that the variables are being used.
func init() { use(buildTimestamp, buildCommitHash, buildVersion, buildRelease) }

func use(...interface{}) {}
