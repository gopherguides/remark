package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HTML_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var html HTML
	r.Equal("html", html.Name())
}

func Test_HTML_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	html := HTML{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	html.Set("k1", "one")
	html.Set("k2", "two")
	html.Set("k3", "three")

	exp := template.HTML(`<html k1="one" k2="two" k3="three">hello</html>`)
	act := html.HTML()
	r.Equal(exp, act)
}
