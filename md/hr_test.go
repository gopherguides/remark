package md

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HR_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "hr")
}

func Test_HR_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	hr := HR{}

	r.Equal(template.HTML("<hr>"), hr.HTML())
}
