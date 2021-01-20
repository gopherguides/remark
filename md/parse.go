package md

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gopherguides/remark"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

type Parser struct {
	*parser.Parser
	Root string
}

func NewParser(root string) *Parser {
	return &Parser{
		Root:   root,
		Parser: newParser(),
	}
}

func ParseFile(name string) (remark.Tags, error) {
	root := filepath.Dir(name)
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return Parse(root, b)
}

func Parse(root string, src []byte) (remark.Tags, error) {
	p := NewParser(root)

	node := p.Parse(src)

	if node == nil {
		return nil, fmt.Errorf("nil node")
	}

	return p.NewDocument(node)
}

func (p *Parser) parseNode(node ast.Node) (remark.Tag, error) {
	var name string
	switch t := node.(type) {
	case *ast.Document:
	case *ast.Heading:
		name = fmt.Sprintf("h%d", t.Level)
	case *ast.Text:
		name = ""
	case *ast.Paragraph:
		name = "p"
	case *ast.HTMLSpan:
		return p.parseHTML(t.Literal)
	case *ast.HTMLBlock:
		return p.parseHTML(t.Literal)
	case *ast.List:
		name = "ul"
		if t.ListFlags == ast.ListTypeOrdered {
			name = "ol"
		}
	case *ast.Link:
		return p.parseLink(t)
	case *ast.ListItem:
		name = "li"
	case *ast.Code:
		name = "code"
	case *ast.CodeBlock:
		return p.NewCodeBlock(t)
	case *ast.Emph:
		name = "em"
	default:
		return nil, fmt.Errorf("unexpected node %T", node)
	}

	if l := node.AsLeaf(); l != nil {
		return p.ParseLeaf(name, l)
	}
	if c := node.AsContainer(); c != nil {
		return p.ParseContainer(name, c)
	}

	return nil, fmt.Errorf("expected node %v", node)
}

func (p *Parser) ParseLeaf(name string, node *ast.Leaf) (remark.Tag, error) {
	l := string(node.Literal)
	l = strings.TrimSpace(l)
	if len(l) == 0 {
		return nil, nil
	}
	return remark.String(l), nil
}

func (p *Parser) ParseContainer(name string, node *ast.Container) (remark.Tag, error) {
	t := remark.NewGeneric(name)
	for _, n := range node.Children {
		nt, err := p.parseNode(n)
		if err != nil {
			return nil, err
		}
		if nt != nil {
			t.Children = append(t.Children, nt)
		}
	}
	return t, nil
}

func newParser() *parser.Parser {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs |
		parser.SuperSubscript | parser.OrderedListStart | parser.Footnotes | parser.LaxHTMLBlocks
	return parser.NewWithExtensions(extensions)
}
