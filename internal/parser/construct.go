package parser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gopherguides/remark/ast"
)

func newHR(c *current) (*ast.HR, error) {
	s := string(c.text)
	s = strings.TrimSpace(s)
	h := &ast.HR{
		Text: s,
	}
	return h, nil
}

func newParagraph(c *current, i interface{}) (*ast.Paragraph, error) {
	ps, err := ast.NewNodes(i)
	if err != nil {
		return nil, err
	}

	p := &ast.Paragraph{
		Nodes: ps,
	}
	fmt.Printf(">>>TODO internal/parser/construct.go:29: p (%T): %v \n", p, p)
	return p, nil
}

func newString(c *current) ast.String {
	s := string(c.text)
	s = strings.TrimSpace(s)
	fmt.Printf(">>>TODO internal/parser/construct.go:36: s (%T): %v \n", s, s)
	return ast.String(s)
}

func newNodes(c *current, i interface{}) (ast.Nodes, error) {
	return ast.NewNodes(i)
}

func iiToString(i interface{}) (string, error) {
	ti, ok := i.([]interface{})
	if !ok {
		return "", fmt.Errorf("expected i to be []interface{} %T", i)
	}

	bb := &bytes.Buffer{}
	for _, t := range ti {
		b, ok := t.([]byte)
		if !ok {
			return "", fmt.Errorf("expected t to be []byte %T", t)
		}
		bb.Write(b)
	}
	s := bb.String()

	return strings.TrimSpace(s), nil
}

func newHeader(c *current, level interface{}, text interface{}, nodes interface{}) (*ast.Header, error) {
	l, ok := level.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected level to be []interface{} %T", level)
	}

	ts, ok := text.(*ast.Paragraph)
	if !ok {
		return nil, fmt.Errorf("expected text to be *ast.Paragraph %T", text)
	}

	h := &ast.Header{
		Text:  ts,
		Level: len(l),
	}

	ni, ok := nodes.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected nodes to be []interface{} %T", nodes)
	}
	for _, n := range ni {
		if node, ok := n.(ast.Node); ok {
			h.Nodes = append(h.Nodes, node)
		}
	}
	fmt.Printf(">>>TODO internal/parser/construct.go:81: h (%T): %v \n", h, h)
	return h, nil
}
