package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewDocument(node ast.Node) (remark.Tags, error) {
	var tags remark.Tags

	sections := [][]ast.Node{{}}
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
