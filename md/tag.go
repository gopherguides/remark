package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/htm"
)

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
		if md, ok := t.(*htm.Document); ok {
			tag.Data = md.Metadata()
		}
		tag.Append(t)
	}

	return tag, nil
}

func (p *Parser) NewDocument(node ast.Node) (remark.Tags, error) {
	var tags remark.Tags

	sections := [][]ast.Node{[]ast.Node{}}
	var ind int

	for _, n := range node.GetChildren() {
		if _, ok := n.(*ast.HorizontalRule); ok {
			sections = append(sections, []ast.Node{})
			ind++
			continue
		}
		sections[ind] = append(sections[ind], n)
	}

	for _, nodes := range sections {
		tag, err := p.NewSection(nodes)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
