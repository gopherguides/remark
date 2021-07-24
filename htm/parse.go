package htm

import (
	"bytes"
	"fmt"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/px"
	"golang.org/x/net/html"
)

func Parse(src []byte) (*Document, error) {
	p := Parser{}
	return p.Parse(src)
}

type Parser struct {
	*px.Module
}

func NewParser(mod *px.Module) *Parser {
	return &Parser{
		Module: mod,
	}
}

func (p *Parser) ParseNode(node *html.Node) (*Document, error) {
	t, err := p.NewTag(node)
	if err != nil {
		return nil, err
	}

	doc, ok := t.(*Document)
	if !ok {
		return nil, fmt.Errorf("expected Document %T %v", t, t)
	}

	if md, ok := p.findMetadata(doc); ok {
		doc.Data = md
	}
	return doc, nil
}

func (p *Parser) Parse(src []byte) (*Document, error) {
	f := bytes.NewReader(src)

	node, err := html.Parse(f)
	if err != nil {
		return nil, err
	}
	return p.ParseNode(node)
}

func (p *Parser) findMetadata(tag remark.Tag) (remark.Metadata, bool) {
	if md, ok := tag.(remark.Metadata); ok {
		return md, true
	}
	for _, c := range tag.GetChildren() {
		if md, ok := p.findMetadata(c); ok {
			return md, ok
		}
	}
	return remark.Metadata{}, false
}
