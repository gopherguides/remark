package htm

import (
	"testing"
	"testing/fstest"

	"github.com/gopherguides/remark"
	"github.com/markbates/fsx"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func testParser(t *testing.T, cab fstest.MapFS) *Parser {
	t.Helper()
	return NewParser(NewModule(t, cab))
}

func NewModule(t *testing.T, cab fstest.MapFS) *remark.Module {
	t.Helper()
	return &remark.Module{
		FS:   fsx.NewFS(cab),
		Name: "module.md",
	}
}

func DocNode(t *testing.T) *html.Node {
	t.Helper()
	return &html.Node{
		Type: html.DocumentNode,
	}
}

func ElementNode(t *testing.T, name string) *html.Node {
	t.Helper()

	return &html.Node{
		Data:     name,
		DataAtom: atom.Lookup([]byte(name)),
		Type:     html.ElementNode,
	}
}

func AttrNode(t *testing.T, name string, ats map[string]string) *html.Node {
	t.Helper()
	node := ElementNode(t, name)
	node.Attr = MapToAttr(t, ats)
	return node
}

func TextNode(t *testing.T, text string) *html.Node {
	t.Helper()
	return &html.Node{
		Data: text,
		Type: html.TextNode,
	}
}

func MapToAttr(t *testing.T, m map[string]string) []html.Attribute {
	t.Helper()

	ats := make([]html.Attribute, 0, len(m))
	for k, v := range m {
		ats = append(ats, html.Attribute{
			Key: k,
			Val: v,
		})
	}

	return ats
}
