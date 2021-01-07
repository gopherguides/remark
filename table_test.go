package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Table_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var table Table
	r.Equal("table", table.Name())
}

func Test_Table_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	table := Table{
		Attributes: Attributes{},
		Elems: Elements{
			String("hello"),
		},
	}
	table.Set("k1", "one")
	table.Set("k2", "two")
	table.Set("k3", "three")

	exp := template.HTML(`<table k1="one" k2="two" k3="three">hello</table>`)
	act := table.HTML()
	r.Equal(exp, act)
}
