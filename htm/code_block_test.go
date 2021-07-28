package htm

import (
	"testing"
	"testing/fstest"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/internal/rpm"
	"github.com/stretchr/testify/require"
)

func Test_NewCodeBlock(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const gof = `package remark

import (
	"fmt"
)

// section: one
type CodeBlock struct {
	*Generic
	Language string
	Snippets Snippets
}
// section: one

// section: two
func (code CodeBlock) Name() string {
	return "code"
}
// section: two`

	const top = `
# Top Level

<code src="src/foo.go"></code>`

	cab := rpm.ModuleFS(t)

	cab["module.md"] = &fstest.MapFile{
		Data: []byte(top),
	}

	cab["src/foo.go"] = &fstest.MapFile{
		Data: []byte(gof),
	}

	p := testParser(t, cab)
	tags, err := p.Parse([]byte(top))
	r.NoError(err)
	r.NotEmpty(tags)

	cbs := tags.Children.FindAll("code")
	r.NotEmpty(cbs)
	r.Len(cbs, 1)

	code, ok := cbs[0].(*remark.CodeBlock)
	r.True(ok)
	r.Len(code.GetChildren(), 1)
	r.Len(code.Snippets, 2)

	table := []struct {
		Name  string
		Start int
		End   int
		Exp   string
	}{
		{Name: "one", Start: 7, End: 13, Exp: "type CodeBlock struct"},
		{Name: "two", Start: 15, End: 19, Exp: "(code CodeBlock) Name()"},
	}

	for _, tt := range table {
		tt := tt
		t.Run(tt.Name, func(st *testing.T) {
			r := require.New(st)
			snip := code.Snippets[tt.Name]
			r.NotEqual(remark.Snippet{}, snip)
			r.Equal("go", snip.Language)
			r.Equal(tt.Name, snip.Name)
			r.Equal("src/foo.go", snip.File)
			r.Equal(tt.Start, snip.Start)
			r.Equal(tt.End, snip.End)
			r.Contains(snip.Content, tt.Exp)
		})
	}

}
