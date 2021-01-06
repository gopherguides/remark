package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Row_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var row Row
	r.Equal("row", row.Name())
}

func Test_Row_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	row := Row{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	row.Set("k1", "one")
	row.Set("k2", "two")
	row.Set("k3", "three")

	exp := template.HTML(`<row k1="one" k2="two" k3="three">hello</row>`)
	act := row.HTML()
	r.Equal(exp, act)
}
