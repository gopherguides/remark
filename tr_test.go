package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TR_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var tr TR
	r.Equal("tr", tr.Name())
}

func Test_TR_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tr := TR{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	tr.Set("k1", "one")
	tr.Set("k2", "two")
	tr.Set("k3", "three")

	exp := template.HTML(`<tr k1="one" k2="two" k3="three">hello</tr>`)
	act := tr.HTML()
	r.Equal(exp, act)
}
