package htm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func Test_NewCodeBlock(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const fp = "../internal/testdata/buffalo/src/initial/actions/app.go"

	node := &html.Node{
		Data: "code",
		Attr: []html.Attribute{
			{Key: "src", Val: fp},
		},
	}

	p := &Parser{}
	code, err := p.NewCodeBlock(node)
	r.NoError(err)

	r.Len(code.Elements(), 1)
	r.Len(code.Snippets, 4)

	table := []struct {
		Name  string
		Start int
		End   int
		Exp   string
	}{
		{Name: "imports", Start: 3, End: 16, Exp: "gobuffalo/envy"},
		{Name: "func", Start: 24, End: 70, Exp: "App is where all"},
		{Name: "app", Start: 38, End: 69, Exp: `app.GET("/", HomeHandler)`},
		{Name: "home", Start: 59, End: 61, Exp: `app.GET("/", HomeHandler)`},
	}

	for _, tt := range table {
		tt := tt
		t.Run(tt.Name, func(st *testing.T) {
			r := require.New(st)
			snip := code.Snippets[tt.Name]
			r.Equal("go", snip.Language)
			r.Equal(tt.Name, snip.Name)
			r.Equal(fp, snip.File)
			r.Equal(tt.Start, snip.Start)
			r.Equal(tt.End, snip.End)
			r.Contains(snip.Content, tt.Exp)
		})
	}

}
