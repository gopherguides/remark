package htm

import (
	"bytes"
	"fmt"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func Parse(src []byte) (*Document, error) {
	p := Parser{}
	return p.Parse(src)
}

type Parser struct {
	Root string
}

func (p *Parser) Parse(src []byte) (*Document, error) {
	f := bytes.NewReader(src)

	node, err := html.Parse(f)
	if err != nil {
		return nil, err
	}

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
