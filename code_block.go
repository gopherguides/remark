package remark

import (
	"bytes"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type CodeBlock struct {
	*ast.CodeBlock
}

func (cb CodeBlock) Language() string {
	s := string(cb.Info)
	if len(s) == 0 {
		s = "plain"
	}
	return s
}

func (cb CodeBlock) String() string {
	if !cb.IsFenced {
		return string(cb.Literal)
	}
	bb := &bytes.Buffer{}
	fmt.Fprintf(bb, "```%s\n", cb.Language())
	bb.Write(bytes.TrimSpace(cb.Literal))
	fmt.Fprintf(bb, "\n```")
	return bb.String()
}

type CodeBlocks []CodeBlocks

func NewCodeBlock(node *ast.CodeBlock) (CodeBlock, error) {
	cb := CodeBlock{
		CodeBlock: node,
	}
	return cb, nil
}
