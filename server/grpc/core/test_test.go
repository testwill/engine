package core

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/docker/docker/pkg/archive"
	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/database"
	"github.com/mesg-foundation/core/sdk"
	"github.com/stretchr/testify/require"
)

var (
	taskServicePath = filepath.Join("..", "..", "..", "service-test", "task")
)

const (
	servicedbname  = "service.db.test"
	instancedbname = "instance.db.test"
	execdbname     = "exec.db.test"
)

func newServerWithContainer(t *testing.T, c container.Container) (*Server, func()) {
	db, err := database.NewServiceDB(servicedbname)
	require.NoError(t, err)

	instanceDB, err := database.NewInstanceDB(instancedbname)
	require.NoError(t, err)

	execDB, err := database.NewExecutionDB(execdbname)
	require.NoError(t, err)

	a := sdk.New(c, db, instanceDB, execDB)

	server := NewServer(a)

	closer := func() {
		db.Close()
		execDB.Close()
		os.RemoveAll(servicedbname)
		os.RemoveAll(instancedbname)
		os.RemoveAll(execdbname)
	}
	return server, closer
}

func newServer(t *testing.T) (*Server, func()) {
	c, err := container.New()
	require.NoError(t, err)
	return newServerWithContainer(t, c)
}

func serviceTar(t *testing.T, path string) io.Reader {
	reader, err := archive.TarWithOptions(path, &archive.TarOptions{
		Compression: archive.Gzip,
	})
	require.NoError(t, err)
	return reader
}
