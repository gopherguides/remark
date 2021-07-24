package md

import (
	"testing"
	"testing/fstest"

	"github.com/gopherguides/remark/internal/rpm"
	"github.com/gopherguides/remark/px"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func Test_Parser_NewInclude(t *testing.T) {
	r := require.New(t)

	cab := rpm.ModuleFS(t)
	mod := &px.Module{
		FS:      cab,
		ModName: "module.md",
	}
	srcKey := "assignments/01/assignment.md"
	cab[srcKey] = &fstest.MapFile{
		Data: []byte(`
## Assignment

<img src="assets/foo.png">
		`),
	}
	cab["assignments/01/assets/foo.png"] = rpm.File(t, "assignments/01/assets/foo.png")

	p := NewParser(mod)

	n := &html.Node{
		Data: "include",
		Attr: []html.Attribute{
			{Key: "src", Val: srcKey},
		},
	}

	i, err := p.NewInclude(n)
	r.NoError(err)
	r.NotNil(i)

	r.Equal("include", i.Name)

	b := i.String()
	r.Equal(``, b)
}
