package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Link_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var a Link
	r.Equal("a", a.Name())
}

func Test_Link_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	a := Link{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	a.Set("k1", "one")
	a.Set("k2", "two")
	a.Set("k3", "three")

	exp := template.HTML(`<a k1="one" k2="two" k3="three">hello</a>`)
	act := a.HTML()
	r.Equal(exp, act)
}
