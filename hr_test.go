package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HR_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var hr HR
	r.Equal("hr", hr.Name())
}

func Test_HR_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	hr := HR{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	hr.Set("k1", "one")
	hr.Set("k2", "two")
	hr.Set("k3", "three")

	exp := template.HTML(`<hr k1="one" k2="two" k3="three">hello</hr>`)
	act := hr.HTML()
	r.Equal(exp, act)
}
