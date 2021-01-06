package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Strike_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var strike Strike
	r.Equal("strike", strike.Name())
}

func Test_Strike_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	strike := Strike{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	strike.Set("k1", "one")
	strike.Set("k2", "two")
	strike.Set("k3", "three")

	exp := template.HTML(`<strike k1="one" k2="two" k3="three">hello</strike>`)
	act := strike.HTML()
	r.Equal(exp, act)
}
