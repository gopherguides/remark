package md

import (
	"testing"
	"testing/fstest"

	"github.com/gopherguides/remark/internal/rpm"
	"github.com/gopherguides/remark/px"
	"github.com/stretchr/testify/require"
)

func Test_Parse_NewSection(t *testing.T) {
	exp := `Intro to Go: Getting Started`

	t.Parallel()
	r := require.New(t)

	cab := rpm.ModuleFS(t)
	mod := &px.Module{
		FS:      cab,
		ModName: "module.md",
	}
	cab["module.md"] = &fstest.MapFile{
		Data: []byte(`# Intro to Go: Getting Started`),
	}

	p := NewParser(mod)
	tags, err := p.Parse()
	r.NoError(err)
	r.NotZero(tags)

	r.Contains(tags.String(), exp)
}
