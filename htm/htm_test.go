package htm

import (
	"io/fs"
	"testing"

	"github.com/gopherguides/remark/px"
)

func NewModule(t *testing.T, cab fs.FS) *px.Module {
	t.Helper()
	return &px.Module{
		FS:      cab,
		ModName: "module.md",
	}
}
