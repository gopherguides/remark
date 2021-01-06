package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BlockQuote_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var blockquote BlockQuote
	r.Equal("blockquote", blockquote.Name())
}

func Test_BlockQuote_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	blockquote := BlockQuote{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	blockquote.Set("k1", "one")
	blockquote.Set("k2", "two")
	blockquote.Set("k3", "three")

	exp := template.HTML(`<blockquote k1="one" k2="two" k3="three">hello</blockquote>`)
	act := blockquote.HTML()
	r.Equal(exp, act)
}
