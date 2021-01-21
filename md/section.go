package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/htm"
)

type Section struct {
	*remark.Generic
	Title string
}

func (s *Section) Metadata() remark.Metadata {
	if s.Data == nil {
		s.Data = remark.Metadata{}
	}
	return s.Data
}

func (p *Parser) NewSection(nodes []ast.Node) (*Section, error) {
	g := remark.NewGeneric("section")
	tag := &Section{
		Generic: g,
	}

	for _, n := range nodes {
		t, err := p.parseNode(n)
		if err != nil {
			return nil, err
		}
		if h, ok := t.(*remark.Heading); ok {
			if h.Name == "h1" || len(tag.Title) == 0 {
				tag.Title = h.Children.String()
			}
		}
		if md, ok := t.(*htm.Document); ok {
			tag.Data = md.Metadata()
		}

		if _, ok := tag.Data["overview"]; !ok {
			tag.Append(t)
			continue
		}

		tag.Append(t)
	}

	return tag, nil
}
