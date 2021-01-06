package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Head_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var head Head
	r.Equal("head", head.Name())
}

func Test_Head_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	head := Head{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	head.Set("k1", "one")
	head.Set("k2", "two")
	head.Set("k3", "three")

	exp := template.HTML(`<head k1="one" k2="two" k3="three">hello</head>`)
	act := head.HTML()
	r.Equal(exp, act)
}
