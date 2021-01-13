package htm

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func ParseFile(src string) (remark.Elements, error) {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return nil, err
	}
	return Parse(b)
}

func Parse(src []byte) (remark.Elements, error) {
	p := Parser{}
	return p.Parse(src)
}

type Parser struct {
	Root     string
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
		// el, err := p.NewElement(n)
		// if err != nil {
		// 	return err
		// }
		// p.elements = append(p.elements, el)
		// return nil
	case html.ElementNode:
		el, err := p.NewElement(n)
		if err != nil {
			return err
		}
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
