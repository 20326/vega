package version

import (
	"github.com/coreos/go-semver/semver"
)

var (
	// GitCommit is the git commit that was compiled
	GitCommit string
	// Major is for an API incompatible changes.
	Major int64 = 0
	// Minor is for functionality in a backwards-compatible manner.
	Minor int64 = 1
	// Patch is for backwards-compatible bug fixes.
	Patch int64
	// Pre indicates prerelease.
	Pre = ""
	// Dev indicates development branch. Releases will be empty string.
	Dev string
)

// Version is the specification version that the package types support.
var Version = semver.Version{
	Major:      Major,
	Minor:      Minor,
	Patch:      Patch,
	PreRelease: semver.PreRelease(Pre),
	Metadata:   Dev,
}
