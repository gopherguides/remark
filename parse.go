package remark

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gopherguides/remark/internal/md"
)

type Parser struct {
	*parser.Parser
}

func (p *Parser) Parse(src []byte) (*Module, error) {
	if p.Parser == nil {
		p.Parser = md.NewParser()
	}

	node := markdown.Parse(src, p.Parser)

	doc, ok := node.(*ast.Document)
	if !ok {
		return nil, fmt.Errorf("expected a document, got %T", node)
	}

	m, err := NewModule(doc)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (p *Parser) ParseFile(filename string) (*Module, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return p.Parse(b)
}

func (p *Parser) ParseReader(r io.Reader) (*Module, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return p.Parse(b)
}

func Parse(src []byte) (*Module, error) {
	p := &Parser{}
	return p.Parse(src)
}

func ParseFile(filename string) (*Module, error) {
	p := &Parser{}
	return p.ParseFile(filename)
}

func ParseReader(r io.Reader) (*Module, error) {
	p := &Parser{}
	return p.ParseReader(r)
}
