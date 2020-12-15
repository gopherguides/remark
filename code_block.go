package remark

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type CodeBlock struct {
	*ast.CodeBlock `json:"-"`
	Language       string
	Text           string
}

func (cb CodeBlock) String() string {
	if !cb.IsFenced {
		return cb.Text
	}
	bb := &bytes.Buffer{}
	fmt.Fprintf(bb, "```%s\n", cb.Language)
	fmt.Fprint(bb, strings.TrimSpace(cb.Text))
	fmt.Fprintf(bb, "\n```")
	return bb.String()
}

func (cb CodeBlock) HTML() template.HTML {
	pre := tags.New("pre", nil)
	code := tags.New("code", tags.Options{
		"class": cb.Language,
		"body":  cb.Text,
	})
	pre.Append(code)
	return pre.HTML()
}

type CodeBlocks []CodeBlocks

func NewCodeBlock(node *ast.CodeBlock) (*CodeBlock, error) {
	cb := &CodeBlock{
		CodeBlock: node,
		Language:  string(node.Info),
		Text:      string(node.Literal),
	}

	if len(cb.Language) == 0 {
		cb.Language = "plain"
	}

	return cb, nil
}
