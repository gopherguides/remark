package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Code_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const input = "`foo += bar`"

	m := parseT(t, input)
	r.Equal(input, m.String())
}
