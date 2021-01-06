package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Div_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var div Div
	r.Equal("div", div.Name())
}

func Test_Div_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	div := Div{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	div.Set("k1", "one")
	div.Set("k2", "two")
	div.Set("k3", "three")

	exp := template.HTML(`<div k1="one" k2="two" k3="three">hello</div>`)
	act := div.HTML()
	r.Equal(exp, act)
}
