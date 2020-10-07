package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func parseT(t *testing.T, input string) *Module {
	t.Helper()

	r := require.New(t)

	m, err := Parse([]byte(input))
	r.NoError(err)
	r.NotNil(m)
	return m
}
