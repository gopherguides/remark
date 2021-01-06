package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Sub_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var sub Sub
	r.Equal("sub", sub.Name())
}

func Test_Sub_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	sub := Sub{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	sub.Set("k1", "one")
	sub.Set("k2", "two")
	sub.Set("k3", "three")

	exp := template.HTML(`<sub k1="one" k2="two" k3="three">hello</sub>`)
	act := sub.HTML()
	r.Equal(exp, act)
}
