package px

import (
	"testing"

	"github.com/gopherguides/remark/internal/rpm"
	"github.com/stretchr/testify/require"
)

func Test_FindModules(t *testing.T) {
	r := require.New(t)

	cab := rpm.ModuleFS(t)
	r.NotNil(cab)

	mods, err := FindModules(cab)
	r.NoError(err)
	r.Len(mods, 1)
}

func Test_FindModules_Sub(t *testing.T) {
	r := require.New(t)

	cab := rpm.ModuleFS(t)

	cab["assignment/module.md"] = rpm.File(t, "/assignment/module.md")

	mods, err := FindModules(cab)
	r.NoError(err)
	r.Len(mods, 2)
}
