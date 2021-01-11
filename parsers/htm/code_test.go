package htm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func Test_NewCode(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const fp = "../internal/testdata/buffalo/src/initial/main.go"

	node := &html.Node{
		Data: "code",
		Attr: []html.Attribute{
			{Key: "src", Val: fp},
		},
	}

	p := &Parser{}
	code, err := p.NewCode(node)
	r.NoError(err)

	r.Len(code.Elements(), 1)
	r.Len(code.Snippets, 3)

	snip := code.Snippets["snipa"]
	r.Equal("go", snip.Language)
	r.Equal("snipa", snip.Name)
	r.Equal(fp, snip.File)
	r.Equal(9, snip.Start)
	r.Equal(25, snip.End)

	exp := `All we ask is that`
	r.Contains(snip.Content, exp)
	r.NotContains(snip.Content, "snipa")

	snip = code.Snippets["snipb"]
	r.Equal("go", snip.Language)
	r.Equal("snipb", snip.Name)
	r.Equal(fp, snip.File)
	r.Equal(27, snip.Start)
	r.Equal(45, snip.End)

	exp = `Nginx`
	r.Contains(snip.Content, exp)
	r.NotContains(snip.Content, "snipb")

	snip = code.Snippets["snipc"]
	r.Equal("go", snip.Language)
	r.Equal("snipc", snip.Name)
	r.Equal(fp, snip.File)
	r.Equal(17, snip.Start)
	r.Equal(19, snip.End)

	exp = `app := actions.App()`
	r.Contains(snip.Content, exp)
	r.NotContains(snip.Content, "app.Serve")
}
