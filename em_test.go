package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Em_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var em Em
	r.Equal("em", em.Name())
}

func Test_Em_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em := Em{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	em.Set("k1", "one")
	em.Set("k2", "two")
	em.Set("k3", "three")

	exp := template.HTML(`<em k1="one" k2="two" k3="three">hello</em>`)
	act := em.HTML()
	r.Equal(exp, act)
}
