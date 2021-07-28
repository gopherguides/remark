package px

import (
	"testing"

	"github.com/gopherguides/remark/internal/rpm"

	"github.com/stretchr/testify/require"
)

func Test_Module_ReadFile(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mfs := rpm.ModuleFS(t)
	mod := &Module{
		FS: NewFS(mfs),
	}

	act, err := mod.ReadFile("module.md")
	r.NoError(err)
	r.Equal(`MODULE.MD`, string(act))
}

func Test_FindModules(t *testing.T) {
	r := require.New(t)

	cab := rpm.ModuleFS(t)
	r.NotNil(cab)

	mods, err := FindModules(NewFS(cab))
	r.NoError(err)
	r.Len(mods, 1)
}

func Test_FindModules_Sub(t *testing.T) {
	r := require.New(t)

	cab := rpm.ModuleFS(t)

	cab["assignment/module.md"] = rpm.File(t, "/assignment/module.md")

	mods, err := FindModules(NewFS(cab))
	r.NoError(err)
	r.Len(mods, 2)
}
