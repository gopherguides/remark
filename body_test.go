package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Body_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var body Body
	r.Equal("body", body.Name())
}

func Test_Body_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	body := Body{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	body.Set("k1", "one")
	body.Set("k2", "two")
	body.Set("k3", "three")

	exp := template.HTML(`<body k1="one" k2="two" k3="three">hello</body>`)
	act := body.HTML()
	r.Equal(exp, act)
}
