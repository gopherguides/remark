package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TableCell_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	td := TableCell{
		Elements: Elements{
			String("crazy"),
		},
	}

	exp := template.HTML(`<td>crazy</td>`)
	r.Equal(exp, td.HTML())
}
