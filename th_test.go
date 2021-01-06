package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TH_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var th TH
	r.Equal("th", th.Name())
}

func Test_TH_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	th := TH{
		Attributes: Attributes{},
		Elements: Elements{
			String("hello"),
		},
	}
	th.Set("k1", "one")
	th.Set("k2", "two")
	th.Set("k3", "three")

	exp := template.HTML(`<th k1="one" k2="two" k3="three">hello</th>`)
	act := th.HTML()
	r.Equal(exp, act)
}
