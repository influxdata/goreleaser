// Package defaults make the list of Defaulter implementations available
// so projects extending GoReleaser are able to use it, namely, GoDownloader.
package defaults

import (
	"fmt"

	"github.com/influxdata/goreleaser/internal/pipe/archive"
	"github.com/influxdata/goreleaser/internal/pipe/artifactory"
	"github.com/influxdata/goreleaser/internal/pipe/brew"
	"github.com/influxdata/goreleaser/internal/pipe/build"
	"github.com/influxdata/goreleaser/internal/pipe/checksums"
	"github.com/influxdata/goreleaser/internal/pipe/docker"
	"github.com/influxdata/goreleaser/internal/pipe/env"
	"github.com/influxdata/goreleaser/internal/pipe/nfpm"
	"github.com/influxdata/goreleaser/internal/pipe/project"
	"github.com/influxdata/goreleaser/internal/pipe/release"
	"github.com/influxdata/goreleaser/internal/pipe/s3"
	"github.com/influxdata/goreleaser/internal/pipe/scoop"
	"github.com/influxdata/goreleaser/internal/pipe/sign"
	"github.com/influxdata/goreleaser/internal/pipe/snapcraft"
	"github.com/influxdata/goreleaser/internal/pipe/snapshot"
	"github.com/influxdata/goreleaser/pkg/context"
)

// Defaulter can be implemented by a Piper to set default values for its
// configuration.
type Defaulter interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Default(ctx *context.Context) error
}

// Defaulters is the list of defaulters
// nolint: gochecknoglobals
var Defaulters = []Defaulter{
	env.Pipe{},
	snapshot.Pipe{},
	release.Pipe{},
	project.Pipe{},
	archive.Pipe{},
	build.Pipe{},
	nfpm.Pipe{},
	snapcraft.Pipe{},
	checksums.Pipe{},
	sign.Pipe{},
	docker.Pipe{},
	artifactory.Pipe{},
	s3.Pipe{},
	brew.Pipe{},
	scoop.Pipe{},
}
