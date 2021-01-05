package md

import (
	"html/template"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
)

func Test_ListItem_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	li := &remark.ListItem{
		Elements: Elements{String("Flowers")},
	}

	exp := template.HTML(`<li>Flowers</li>`)
	r.Equal(exp, li.HTML())
}
