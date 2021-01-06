package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Sup_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var sup Sup
	r.Equal("sup", sup.Name())
}

func Test_Sup_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	sup := Sup{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	sup.Set("k1", "one")
	sup.Set("k2", "two")
	sup.Set("k3", "three")

	exp := template.HTML(`<sup k1="one" k2="two" k3="three">hello</sup>`)
	act := sup.HTML()
	r.Equal(exp, act)
}
