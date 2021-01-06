package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Heading_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var h Heading
	r.Equal("h", h.Name())
}

func Test_Heading_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	h := Heading{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	h.Set("k1", "one")
	h.Set("k2", "two")
	h.Set("k3", "three")

	exp := template.HTML(`<h k1="one" k2="two" k3="three">hello</h>`)
	act := h.HTML()
	r.Equal(exp, act)
}
