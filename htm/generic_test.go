package htm

import (
	"testing"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/internal/rpm"
	"github.com/stretchr/testify/require"
)

func Test_Parser_NewGeneric(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	div := AttrNode(t, "div", map[string]string{
		"class": "school",
		"id":    "42",
	})

	divp := ElementNode(t, "p")
	divp.FirstChild = TextNode(t, "Hola!")
	div.FirstChild = divp

	mod := NewModule(t, rpm.ModuleFS(t))

	p := NewParser(mod)
	g, err := p.NewGeneric(div)
	r.NoError(err)
	r.NotNil(g)

	r.Equal("div", g.Name)
	ats := remark.Attributes{
		"class": "school",
		"id":    "42",
	}
	r.Equal(ats, g.Attributes)
	r.Len(g.Children, 1)

	pg, ok := g.Children[0].(*remark.Generic)
	r.True(ok)
	r.Equal("p", pg.TagName())
	r.Len(pg.Children, 1)

	txt := pg.Children[0]
	r.Equal("Hola!", txt.String())
}
