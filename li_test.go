package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ListItem_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var li ListItem
	r.Equal("li", li.Name())
}

func Test_ListItem_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	li := ListItem{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	li.Set("k1", "one")
	li.Set("k2", "two")
	li.Set("k3", "three")

	exp := template.HTML(`<li k1="one" k2="two" k3="three">hello</li>`)
	act := li.HTML()
	r.Equal(exp, act)
}
