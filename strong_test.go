package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Strong_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var strong Strong
	r.Equal("strong", strong.Name())
}

func Test_Strong_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	strong := Strong{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	strong.Set("k1", "one")
	strong.Set("k2", "two")
	strong.Set("k3", "three")

	exp := template.HTML(`<strong k1="one" k2="two" k3="three">hello</strong>`)
	act := strong.HTML()
	r.Equal(exp, act)
}
