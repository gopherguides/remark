package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TD_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var td TD
	r.Equal("td", td.Name())
}

func Test_TD_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	td := TD{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	td.Set("k1", "one")
	td.Set("k2", "two")
	td.Set("k3", "three")

	exp := template.HTML(`<td k1="one" k2="two" k3="three">hello</td>`)
	act := td.HTML()
	r.Equal(exp, act)
}
