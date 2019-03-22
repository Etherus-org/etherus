package version

// Major version component of the current release
const Major = 0

// Minor version component of the current release
const Minor = 9

// Fix version component of the current release
const Fix = 2

var (
	// Version is the full version string
	Version = "0.9.3"

	// GitCommit is set with --ldflags "-X main.gitCommit=$(git rev-parse --short HEAD)"
	GitCommit = "0000000"
)

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}
