package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TableRow_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	th := TableRow{
		Elements: Elements{
			TableCell{
				Elements: Elements{String("hooligan")},
			},
			TableCell{
				Elements: Elements{String("hippo")},
			},
		},
	}

	exp := template.HTML(`<tr><td>hooligan</td><td>hippo</td></tr>`)
	r.Equal(exp, th.HTML())
}
