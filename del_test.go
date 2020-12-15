package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Del_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "del")
}

func Test_Del_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	del := Del{
		Elements: Elements{
			String("Hello"),
		},
	}

	exp := `<strike>Hello</strike>`
	r.Equal(template.HTML(exp), del.HTML())
}
