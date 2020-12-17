package remark

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/gobuffalo/tags/v3"
)

type CodeBlock struct {
	Language string
	Text     string
}

func (cb CodeBlock) String() string {
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
