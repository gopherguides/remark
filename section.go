package remark

import (
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
)

type Section struct {
	*Heading
	Children []ast.Node
}

func (s Section) String() string {
	lines := make([]string, 0, len(s.Children))

	for _, c := range s.Children {
		lines = append(lines, md.String(c))
	}
	return strings.Join(lines, "\n\n")
}

type Sections []Section
