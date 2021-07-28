package htm

import (
	"testing"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/internal/rpm"
	"github.com/gopherguides/remark/px"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Test_Parser_NewGeneric(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mod := &px.Module{
		FS: px.NewFS(rpm.ModuleFS(t)),
	}

	p := NewParser(mod)

	node := &html.Node{
		Data:     "div",
		DataAtom: atom.Div,
		Type:     html.ElementNode,
		Attr: []html.Attribute{
			{Key: "class", Val: "school"},
			{Key: "id", Val: "1"},
		},
		FirstChild: &html.Node{
			Data:     "p",
			DataAtom: atom.P,
			Type:     html.ElementNode,
			FirstChild: &html.Node{
				Data: "hello!",
				Type: html.TextNode,
			},
		},
	}

	g, err := p.NewGeneric(node)
	r.NoError(err)
	r.NotNil(g)

	r.Equal("div", g.Name)
	ats := remark.Attributes{
		"class": "school",
		"id":    "1",
	}
	r.Equal(ats, g.Attributes)
	r.Len(g.Children, 1)

	pg, ok := g.Children[0].(*remark.Generic)
	r.True(ok)
	r.Equal("p", pg.TagName())
	r.Len(pg.Children, 1)

	txt := pg.Children[0]
	r.Equal("hello!", txt.String())
}
