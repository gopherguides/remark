package md

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"golang.org/x/net/html"
)

type HTMLParser struct {
	doc *ast.Document
}

func (p *HTMLParser) Parse(b []byte) (*ast.Document, error) {
	node, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	if err := p.node(node); err != nil {
		return nil, err
	}
	return p.doc, nil
}

func (p *HTMLParser) node(n *html.Node) error {
	switch n.Type {
	case html.CommentNode:
		fmt.Println("CommentNode", n)
		// fmt.Printf(">>>TODO internal/md/html.go:23: n (%T): %v \n", n, n)
	case html.DoctypeNode:
		fmt.Println("DoctypeNode", n)
		// fmt.Printf(">>>TODO internal/md/html.go:25: n (%T): %v \n", n, n)
	case html.DocumentNode:
		fmt.Println("DocumentNode", n)
		// fmt.Printf(">>>TODO internal/md/html.go:27: n (%T): %v \n", n, n)
	case html.ElementNode:
		fmt.Println("ElementNode", n.Data)
		// fmt.Printf(">>>TODO internal/md/html.go:28: n.Data (%T): %v \n", n.Data, n.Data)
		// fmt.Printf(">>>TODO internal/md/html.go:29: n.Attr (%T): %v \n", n.Attr, n.Attr)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if err := p.node(c); err != nil {
				return err
			}
		}
		return nil
	case html.RawNode:
		fmt.Println("RawNode", n)
		// fmt.Printf(">>>TODO internal/md/html.go:31: n (%T): %v \n", n, n)
	case html.TextNode:
		d := strings.TrimSpace(n.Data)
		if len(d) == 0 {
			return nil
		}
		fmt.Println("TextNode", d)
	default:
		return fmt.Errorf("unexpected node type %v", n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if err := p.node(c); err != nil {
			return err
		}
	}
	return nil
}

func HTML(b []byte) (*ast.Document, error) {
	hp := HTMLParser{}
	return hp.Parse(b)
}
