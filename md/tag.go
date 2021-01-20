package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewSection(nodes []ast.Node) (remark.Tag, error) {
	tag := remark.NewGeneric("section")

	for _, n := range nodes {
		t, err := p.parseNode(n)
		if err != nil {
			return nil, err
		}
		tag.Append(t)
	}

	tag.Data = tag.GetChildren().Metadata()

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
		if md := tag.GetChildren().Metadata(); md != nil {
			for k, v := range md {
				md[k] = v
			}
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
