// Package pipeline provides generic erros for pipes to use.
package pipeline

import (
	"fmt"

	"github.com/influxdata/goreleaser/internal/pipe/archive"
	"github.com/influxdata/goreleaser/internal/pipe/before"
	"github.com/influxdata/goreleaser/internal/pipe/build"
	"github.com/influxdata/goreleaser/internal/pipe/changelog"
	"github.com/influxdata/goreleaser/internal/pipe/checksums"
	"github.com/influxdata/goreleaser/internal/pipe/defaults"
	"github.com/influxdata/goreleaser/internal/pipe/dist"
	"github.com/influxdata/goreleaser/internal/pipe/docker"
	"github.com/influxdata/goreleaser/internal/pipe/effectiveconfig"
	"github.com/influxdata/goreleaser/internal/pipe/env"
	"github.com/influxdata/goreleaser/internal/pipe/git"
	"github.com/influxdata/goreleaser/internal/pipe/nfpm"
	"github.com/influxdata/goreleaser/internal/pipe/publish"
	"github.com/influxdata/goreleaser/internal/pipe/sign"
	"github.com/influxdata/goreleaser/internal/pipe/snapcraft"
	"github.com/influxdata/goreleaser/internal/pipe/snapshot"
	"github.com/influxdata/goreleaser/pkg/context"
)

// Piper defines a pipe, which can be part of a pipeline (a serie of pipes).
type Piper interface {
	fmt.Stringer

	// Run the pipe
	Run(ctx *context.Context) error
}

// Pipeline contains all pipe implementations in order
// nolint: gochecknoglobals
var Pipeline = []Piper{
	before.Pipe{},          // run global hooks before build
	git.Pipe{},             // get and validate git repo state
	defaults.Pipe{},        // load default configs
	snapshot.Pipe{},        // snapshot version handling
	dist.Pipe{},            // ensure ./dist is clean
	effectiveconfig.Pipe{}, // writes the actual config (with defaults et al set) to dist
	changelog.Pipe{},       // builds the release changelog
	env.Pipe{},             // load and validate environment variables
	build.Pipe{},           // build
	archive.Pipe{},         // archive in tar.gz, zip or binary (which does no archiving at all)
	nfpm.Pipe{},            // archive via fpm (deb, rpm) using "native" go impl
	snapcraft.Pipe{},       // archive via snapcraft (snap)
	checksums.Pipe{},       // checksums of the files
	sign.Pipe{},            // sign artifacts
	docker.Pipe{},          // create and push docker images
	publish.Pipe{},         // publishes artifacts
}
