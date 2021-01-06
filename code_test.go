package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Code_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var code Code
	r.Equal("code", code.Name())
}

func Test_Code_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	code := Code{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	code.Set("k1", "one")
	code.Set("k2", "two")
	code.Set("k3", "three")

	exp := template.HTML(`<code k1="one" k2="two" k3="three">hello</code>`)
	act := code.HTML()
	r.Equal(exp, act)
}
