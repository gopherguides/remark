package md

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gopherguides/remark"
	"github.com/markbates/markdownr/ast"
	"github.com/markbates/markdownr/parser"
)

type Parser struct {
	*parser.Parser
	Root string
}

func NewParser(root string) *Parser {
	if len(filepath.Ext(root)) != 0 {
		root = filepath.Dir(root)
	}
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

func (p *Parser) Parse(src []byte) (remark.Tags, error) {
	t := template.New(p.Root)
	// t = t.Funcs(template.FuncMap{
	// 	"include": p.include,
	// })

	t, err := t.Parse(string(src))
	if err != nil {
		return nil, err
	}

	bb := &bytes.Buffer{}
	if err := t.Execute(bb, nil); err != nil {
		return nil, err
	}
	node := p.Parser.Parse(bb.Bytes())

	if node == nil {
		return nil, fmt.Errorf("nil node")
	}

	return p.NewDocument(node)
}

func Parse(root string, src []byte) (remark.Tags, error) {
	p := NewParser(root)
	return p.Parse(src)
}

func (p *Parser) parseNode(node ast.Node) (remark.Tag, error) {
	var name string
	switch t := node.(type) {
	case *ast.Document:
	case *ast.Heading:
		return p.NewHeading(t)
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
		return p.NewCode(t)
	case *ast.CodeBlock:
		return p.NewCodeBlock(t)
	case *ast.Emph:
		name = "em"
	case *ast.Table:
		name = "table"
	case *ast.TableHeader:
		name = "th"
	case *ast.TableRow:
		name = "tr"
	case *ast.TableCell:
		name = "td"
	case *ast.TableBody:
		name = "tbody"
	case *ast.Strong:
		name = "strong"
	case *ast.BlockQuote:
		name = "blockquote"
	case *ast.Subscript:
		name = "sub"
	case *ast.Superscript:
		name = "sup"
	case *ast.Del:
		name = "del"
	case *ast.Hardbreak:
		return remark.NewGeneric("br"), nil
	case *ast.Image:
		return p.NewImage(t)
	default:
		return nil, fmt.Errorf("unexpected node %T (%q)", node, node)
	}

	if l := node.AsLeaf(); l != nil {
		return p.ParseLeaf(name, l)
	}
	if c := node.AsContainer(); c != nil {
		return p.ParseContainer(name, c)
	}

	return nil, fmt.Errorf("unexpected node %T (%q)", node, node)
}

func (p *Parser) ParseLeaf(name string, node *ast.Leaf) (remark.Tag, error) {
	l := string(node.Literal)
	l = strings.TrimSpace(l)
	if len(l) == 0 {
		return remark.Noop(), nil
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
	extensions := parser.NoIntraEmphasis | parser.Tables | parser.FencedCode | parser.Autolink | parser.Strikethrough | parser.SpaceHeadings | parser.HeadingIDs | parser.BackslashLineBreak | parser.DefinitionLists | parser.AutoHeadingIDs | parser.SuperSubscript | parser.OrderedListStart | parser.Footnotes | parser.LaxHTMLBlocks
	return parser.NewWithExtensions(extensions)
}
