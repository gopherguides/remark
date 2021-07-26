package rpm

import (
	"testing"
	"testing/fstest"
)

func ModuleFS(t *testing.T) fstest.MapFS {
	t.Helper()
	cab := fstest.MapFS{
		"module.md": File(t, "module.md"),
	}
	return cab
}
