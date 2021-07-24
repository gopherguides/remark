package rpm

import (
	"testing"
	"testing/fstest"
)

func ModuleFS(t *testing.T) fstest.MapFS {
	t.Helper()
	cab := fstest.MapFS{
		// "assets":    Dir(t, "assets"),
		"module.md": File(t, "module.md"),
		// "src":       Dir(t, "src"),
	}
	return cab
}
