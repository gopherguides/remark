package md

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/internal/md"
)

type Parser struct {
	*parser.Parser
}

func (p *Parser) Parse(src []byte) (*remark.HTML, error) {
	if p.Parser == nil {
		p.Parser = md.NewParser()
	}

	node := p.Parser.Parse(src)

	doc, ok := node.(*ast.Document)
	if !ok {
		return nil, fmt.Errorf("expected a document, got %T", node)
	}

	m, err := NewDocument(doc)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (p *Parser) ParseFile(filename string) (*remark.HTML, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return p.Parse(b)
}

func (p *Parser) ParseReader(r io.Reader) (*remark.HTML, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return p.Parse(b)
}

func Parse(src []byte) (*remark.HTML, error) {
	p := &Parser{}
	return p.Parse(src)
}

func ParseFile(filename string) (*remark.HTML, error) {
	p := &Parser{}
	return p.ParseFile(filename)
}

func ParseReader(r io.Reader) (*remark.HTML, error) {
	p := &Parser{}
	return p.ParseReader(r)
}
