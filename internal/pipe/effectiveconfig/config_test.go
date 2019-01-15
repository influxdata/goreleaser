package effectiveconfig

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/influxdata/goreleaser/internal/testlib"
	"github.com/influxdata/goreleaser/pkg/config"
	"github.com/influxdata/goreleaser/pkg/context"
	"github.com/stretchr/testify/assert"
)

func TestPipeDescription(t *testing.T) {
	assert.NotEmpty(t, Pipe{}.String())
}

func Test(t *testing.T) {
	folder, back := testlib.Mktmp(t)
	defer back()
	dist := filepath.Join(folder, "dist")
	assert.NoError(t, os.Mkdir(dist, 0755))
	var ctx = context.New(
		config.Project{
			Dist: dist,
		},
	)
	assert.NoError(t, Pipe{}.Run(ctx))
	bts, err := ioutil.ReadFile(filepath.Join(dist, "config.yaml"))
	assert.NoError(t, err)
	assert.NotEmpty(t, string(bts))
}
