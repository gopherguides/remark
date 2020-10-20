package remark

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
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

	b, err := ioutil.ReadFile(filepath.Join("internal", "testdata", name+".md"))
	r.NoError(err)

	b = bytes.TrimSpace(b)

	m, err := Parse(b)
	r.NoError(err)
	fmt.Println(m.String())
	r.Equal(string(b), m.String())
}
