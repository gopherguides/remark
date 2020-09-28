package parser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gopherguides/remark/ast"
)

func newParagraph(c *current, i interface{}) (*ast.Paragraph, error) {
	ps, err := iiToString(i)
	if err != nil {
		return nil, err
	}

	p := &ast.Paragraph{
		Text: ps,
	}
	fmt.Printf(">>>TODO internal/parser/construct.go:19: p (%T): %v \n", p, p)
	return p, nil
}

func newNodes(c *current, i interface{}) (ast.Nodes, error) {
	ii, ok := i.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected a slice of nodes")
	}
	fmt.Printf(">>>TODO internal/parser/construct.go:14: len(ii) (%T): %v \n", len(ii), len(ii))
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

func newHeader(c *current, level interface{}, text interface{}) (*ast.Header, error) {
	l, ok := level.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected level to be []interface{} %T", level)
	}

	ts, err := iiToString(text)
	if err != nil {
		return nil, err
	}

	h := &ast.Header{
		Text:  ts,
		Level: len(l),
	}

	fmt.Printf(">>>TODO internal/parser/construct.go:48: h (%T): %v \n", h, h)
	return h, nil
}
