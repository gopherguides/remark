package remark

import (
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

type Module struct {
	Node     *ast.Document
	Sections Sections
}

func (m Module) String() string {
	lines := make([]string, 0, len(m.Sections))
	for _, s := range m.Sections {
		lines = append(lines, s.String())
	}
	s := strings.Join(lines, "\n\n---\n\n")
	return s
}

type Modules []Module
