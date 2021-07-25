package md

import (
	"fmt"
	"testing"
	"testing/fstest"

	"github.com/gopherguides/remark/internal/rpm"
	"github.com/gopherguides/remark/px"
	"github.com/stretchr/testify/require"
)

func Test_Parser_Include(t *testing.T) {
	r := require.New(t)

	cab := rpm.ModuleFS(t)
	mod := &px.Module{
		FS:      cab,
		ModName: "module.md",
	}

	cab["module.md"] = &fstest.MapFile{
		Data: []byte(`
# Top Level

<include src="assignments/01/assignment.md">
`),
	}

	const asmd = "assignments/01/assignment.md"

	cab[asmd] = &fstest.MapFile{
		Data: []byte(`
## Assignment

<img src="assets/foo.png">
`),
	}

	const png = "assignments/01/assets/foo.png"
	cab[png] = rpm.File(t, png)

	p := NewParser(mod)

	tags, err := p.Parse()
	r.NoError(err)

	s := tags.String()
	r.Contains(s, fmt.Sprintf("<img src=%q />", png))
}
