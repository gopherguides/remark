package htm

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func Parse(src []byte) (remark.Elements, error) {
	p := Parser{}
	return p.Parse(src)
}

type Parser struct {
	elements remark.Elements
}

func (p *Parser) Parse(src []byte) (remark.Elements, error) {
	node, err := html.Parse(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}

	if err := p.node(node); err != nil {
		return nil, err
	}

	return p.elements, nil
}

func (p *Parser) node(n *html.Node) error {
	switch n.Type {
	case html.CommentNode:
		fmt.Println("CommentNode", n)
	case html.DoctypeNode:
		fmt.Println("DoctypeNode", n)
	case html.DocumentNode:
		fmt.Println("DocumentNode", n)
	case html.ElementNode:
		el, err := NewElement(n)
		if err != nil {
			return err
		}
		fmt.Printf(">>>TODO parsers/htm/parser.go:47: el (%T): %v \n", el, el)
		p.elements = append(p.elements, el)
		return nil
	case html.RawNode:
		fmt.Println("RawNode", n)
	case html.TextNode:
		d := strings.TrimSpace(n.Data)
		if len(d) == 0 {
			return nil
		}
		p.elements = append(p.elements, String(d))
		return nil
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

// fmt.Printf(">>>TODO html_block.go:35: doc (%T): %v \n", doc, doc)
// var f func(*html.Node)
// f = func(n *html.Node) {
// 	fmt.Printf(">>>TODO html_block.go:38: n (%T): %v \n", n, n)
// 	if n.Type == html.ElementNode && n.Data == "a" {
// 		for _, a := range n.Attr {
// 			if a.Key == "href" {
// 				fmt.Println(a.Val)
// 				break
// 			}
// 		}
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		f(c)
// 	}
// }
// f(doc)
