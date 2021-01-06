package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Paragraph_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var p Paragraph
	r.Equal("p", p.Name())
}

func Test_Paragraph_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	p := Paragraph{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	p.Set("k1", "one")
	p.Set("k2", "two")
	p.Set("k3", "three")

	exp := template.HTML(`<p k1="one" k2="two" k3="three">hello</p>`)
	act := p.HTML()
	r.Equal(exp, act)
}
