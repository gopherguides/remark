package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TableBody_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var tbody TableBody
	r.Equal("tbody", tbody.Name())
}

func Test_TableBody_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tbody := TableBody{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	tbody.Set("k1", "one")
	tbody.Set("k2", "two")
	tbody.Set("k3", "three")

	exp := template.HTML(`<tbody k1="one" k2="two" k3="three">hello</tbody>`)
	act := tbody.HTML()
	r.Equal(exp, act)
}
