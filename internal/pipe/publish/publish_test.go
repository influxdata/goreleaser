package publish

import (
	"testing"

	"github.com/influxdata/goreleaser/internal/pipe"
	"github.com/influxdata/goreleaser/pkg/config"
	"github.com/influxdata/goreleaser/pkg/context"
	"github.com/stretchr/testify/require"
)

func TestDescription(t *testing.T) {
	require.NotEmpty(t, Pipe{}.String())
}

func TestPublishDisable(t *testing.T) {
	var ctx = context.New(config.Project{})
	ctx.SkipPublish = true
	require.EqualError(t, Pipe{}.Run(ctx), pipe.ErrSkipPublishEnabled.Error())
}

func TestPublish(t *testing.T) {
	var ctx = context.New(config.Project{})
	ctx.Config.Release.Disable = true
	for i := range ctx.Config.Dockers {
		ctx.Config.Dockers[i].SkipPush = true
	}
	require.NoError(t, Pipe{}.Run(ctx))
}
