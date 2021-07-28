package px

import (
	"encoding/json"
	"testing"

	"github.com/gopherguides/remark/internal/rpm"
	"github.com/stretchr/testify/require"
)

func Test_FS_Exists(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mfs := rpm.ModuleFS(t)

	cab := NewFS(mfs)
	r.NotNil(cab)

	r.True(cab.Exists("module.md"))
	r.False(cab.Exists("404"))
}

func Test_FS_ReadFile(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mfs := rpm.ModuleFS(t)
	cab := NewFS(mfs)
	r.NotNil(cab)

	b, err := cab.ReadFile("module.md")
	r.NoError(err)
	r.Equal(`MODULE.MD`, string(b))
}

func Test_FS_Sub(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mfs := rpm.ModuleFS(t)
	rpm.AppendFile(t, mfs, "assets/foo.png")

	cab := NewFS(mfs)
	r.NotNil(cab)
	r.False(cab.Exists("foo.png"))

	kid, err := cab.Sub("assets")
	r.NoError(err)
	r.True(kid.Exists("foo.png"))
}

func Test_FS_MarshalJSON(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mfs := rpm.ModuleFS(t)
	rpm.AppendFile(t, mfs, ".DS_Store")
	rpm.AppendFile(t, mfs, "assets/foo.png")

	cab := NewFS(mfs)
	r.NotNil(cab)

	exp, err := json.Marshal([]string{
		"assets",
		"assets/foo.png",
		"module.md",
	})
	r.NoError(err)

	act, err := cab.MarshalJSON()
	r.NoError(err)
	r.Equal(string(exp), string(act))
}

func Test_FS_Stat(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mfs := rpm.ModuleFS(t)
	cab := NewFS(mfs)
	r.NotNil(cab)

	info, err := cab.Stat("module.md")
	r.NoError(err)
	r.NotNil(info)
}

func Test_Paths(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mfs := rpm.ModuleFS(t)
	rpm.AppendFile(t, mfs, ".DS_Store")
	rpm.AppendFile(t, mfs, "assets/foo.png")

	act, err := Paths(mfs)
	r.NoError(err)

	exp := []string{
		"assets",
		"assets/foo.png",
		"module.md",
	}
	r.Equal(exp, act)
}

func Test_DirFS(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	cab, err := DirFS(".")
	r.NoError(err)
	r.NotNil(cab)
	r.True(cab.Exists("fs_test.go"))
}
