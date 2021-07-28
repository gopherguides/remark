package htm

import (
	"testing"
	"testing/fstest"

	"github.com/gopherguides/remark/px"
)

func NewModule(t *testing.T, cab fstest.MapFS) *px.Module {
	t.Helper()
	return &px.Module{
		FS:   px.NewFS(cab),
		Name: "module.md",
	}
}
