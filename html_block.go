package remark

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gomarkdown/markdown/ast"
	"golang.org/x/net/html"
)

type HTMLBlock struct {
	*ast.HTMLBlock
	Elements Elements
}

func (hb HTMLBlock) String() string {
	return hb.Elements.String()
}

func (hb HTMLBlock) HTML() template.HTML {
	return hb.Elements.HTML()
}

func NewHTMLBlock(node *ast.HTMLBlock) (*HTMLBlock, error) {
	hb := &HTMLBlock{
		HTMLBlock: node,
	}

	doc, err := html.Parse(bytes.NewReader(node.Leaf.Literal))
	if err != nil {
		return nil, err
	}

	fmt.Printf(">>>TODO html_block.go:35: doc (%T): %v \n", doc, doc)
	var f func(*html.Node)
	f = func(n *html.Node) {
		fmt.Printf(">>>TODO html_block.go:38: n (%T): %v \n", n, n)
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return hb, nil
}
