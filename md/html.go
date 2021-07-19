package md

import (
	"bytes"
	"fmt"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/htm"
	"golang.org/x/net/html"
)

func (p *Parser) parseHTMLNode(node *html.Node) (remark.Tag, error) {
	if node == nil {
		return nil, fmt.Errorf("nil node can not be parsed")
	}
	// fmt.Println("TODO >> node", htm.NodeString(node))

	body := htm.HTMLBody(node)
	// fmt.Println("TODO >> ok", ok)
	if c := body.FirstChild; c != nil {
		switch c.Data {
		case "include":
			return p.NewInclude(c)
		}
	}

	// switch node.Data {
	// case "", "html":
	// 	return p.parseHTMLNode(node.FirstChild)
	// case "head":
	// 	return p.parseHTMLNode(node.NextSibling)
	// case "body":
	// 	return p.parseHTMLNode(node.FirstChild)
	// }

	hp := htm.NewParser(p.Root)
	return hp.ParseNode(node)
}

func (p *Parser) parseHTML(src []byte) (remark.Tag, error) {
	node, err := html.Parse(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}
	return p.parseHTMLNode(node)
	// fmt.Println("TODO >> pp.Data", pp.Data)
	// fmt.Println("TODO >> pp.Namespace", pp.Namespace)
	// if pp.Namespace == "html" {
	// 	panic(pp)
	// }
	// // body, err := htm.HTMLBody(pp)
	// // if err != nil {
	// // 	return nil, err
	// // }
	// // fmt.Println("TODO >> body", body)

	// if c := pp.FirstChild; c != nil {
	// 	fmt.Println("TODO >> c", c.Data)
	// 	switch c.Data {
	// 	case "include":
	// 		return p.NewInclude(c)
	// 	}
	// }
	// // s := strings.ToLower(string(src))
	// // if strings.HasPrefix(s, "<include") {
	// // 	doc, err := html.Parse(strings.NewReader(s))
	// // 	if err != nil {
	// // 		return nil, err
	// // 	}
	// // 	body, err := htm.HTMLBody(doc)
	// // 	if err != nil {
	// // 		return nil, err
	// // 	}
	// // 	inc := body.FirstChild
	// // 	panic(inc)
	// // }

	// hp := htm.NewParser(p.Root)
	// tag, err := hp.Parse(src)
	// if err != nil {
	// 	return nil, err
	// }

	// return tag, nil
}
