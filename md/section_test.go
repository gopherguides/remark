package md

import (
	"testing"
	"testing/fstest"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/internal/rpm"

	"github.com/markbates/fsx"
	"github.com/stretchr/testify/require"
)

func Test_Parse_NewSection(t *testing.T) {
	exp := `Intro to Go: Getting Started`

	t.Parallel()
	r := require.New(t)

	cab := rpm.ModuleFS(t)
	mod := &remark.Module{
		FS:   fsx.NewFS(cab),
		Name: "module.md",
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
