package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ListItem_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	li := ListItem{
		Elements: Elements{String("Flowers")},
	}

	exp := template.HTML(`<li>Flowers</li>`)
	r.Equal(exp, li.HTML())
}
