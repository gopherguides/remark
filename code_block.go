package remark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

type CodeBlock struct {
	*ast.CodeBlock `json:"-"`
	Language       string
	Text           string
}

func (cb CodeBlock) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"content":   cb.Text,
		"is_fenced": cb.IsFenced,
		"language":  cb.Language,
		"literal":   cb.String(),
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", cb): m,
	}, "", "  ")
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
