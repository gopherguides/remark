package htm

import (
	"fmt"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/internal/rpm"
	"github.com/gopherguides/remark/px"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func Test_NewCodeBlock(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const fp = "../internal/testdata/week00/assignments/assignment01/src/assignment01_test.go"

	node := &html.Node{
		Data: "code",
		Attr: []html.Attribute{
			{Key: "src", Val: fp},
		},
	}

	mod := &px.Module{
		FS: rpm.ModuleFS(t),
	}

	p := NewParser(mod)
	tag, err := p.NewCodeBlock(node)
	r.NoError(err)

	code, ok := tag.(*remark.CodeBlock)
	r.True(ok)

	r.Len(code.GetChildren(), 1)
	r.Len(code.Snippets, 2)
	fmt.Println(code.Snippets)
	table := []struct {
		Name  string
		Start int
		End   int
		Exp   string
	}{
		{Name: "one", Start: 12, End: 17, Exp: "Your GOPATH has not been set"},
		{Name: "two", Start: 19, End: 27, Exp: "Your PATH does not contain"},
	}

	for _, tt := range table {
		tt := tt
		t.Run(tt.Name, func(st *testing.T) {
			r := require.New(st)
			snip := code.Snippets[tt.Name]
			r.NotEqual(remark.Snippet{}, snip)
			r.Equal("go", snip.Language)
			r.Equal(tt.Name, snip.Name)
			r.Equal(fp, snip.File)
			r.Equal(tt.Start, snip.Start)
			r.Equal(tt.End, snip.End)
			r.Contains(snip.Content, tt.Exp)
		})
	}

}
