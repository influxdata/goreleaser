package project

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/influxdata/goreleaser/pkg/config"
	"github.com/influxdata/goreleaser/pkg/context"
)

func TestCustomProjectName(t *testing.T) {
	var ctx = context.New(config.Project{
		ProjectName: "foo",
		Release: config.Release{
			GitHub: config.Repo{
				Owner: "bar",
				Name:  "bar",
			},
		},
	})
	require.NoError(t, Pipe{}.Default(ctx))
	require.Equal(t, "foo", ctx.Config.ProjectName)
}

func TestEmptyProjectName(t *testing.T) {
	var ctx = context.New(config.Project{
		Release: config.Release{
			GitHub: config.Repo{
				Owner: "bar",
				Name:  "bar",
			},
		},
	})
	require.NoError(t, Pipe{}.Default(ctx))
	require.Equal(t, "bar", ctx.Config.ProjectName)
}
