package md

import (
	"html/template"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
)

func Test_HR_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	hr := remark.HR{}

	r.Equal(template.HTML("<hr>"), hr.HTML())
}
