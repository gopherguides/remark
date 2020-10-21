package remark

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
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

func compare(t *testing.T, name string) {
	t.Helper()
	r := require.New(t)

	fp := filepath.Join("internal", "testdata", name+".md")
	b, err := ioutil.ReadFile(fp)
	r.NoError(err)

	b = bytes.TrimSpace(b)

	m, err := Parse(b)
	r.NoError(err)

	exp := strings.TrimSpace(string(b))
	act := strings.TrimSpace(m.String())

	r.Equal(exp, act)
}
