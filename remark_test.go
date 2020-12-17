package remark

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func parseT(t *testing.T, input string) *Document {
	t.Helper()

	r := require.New(t)

	m, err := Parse([]byte(input))
	r.NoError(err)
	r.NotNil(m)
	return m
}

func td(t *testing.T, name string) string {
	t.Helper()
	r := require.New(t)

	if filepath.Ext(name) == "" {
		name += ".md"
	}

	fp := filepath.Join("internal", "testdata", name)
	b, err := ioutil.ReadFile(fp)
	r.NoError(err)

	b = bytes.TrimSpace(b)
	return string(b)
}

func compareStringer(t *testing.T, name string) {
	t.Helper()
	r := require.New(t)

	s := td(t, name)
	m, err := Parse([]byte(s))
	r.NoError(err)

	exp := strings.TrimSpace(s)
	act := strings.TrimSpace(m.String())
	fmt.Println(act)

	r.Equal(exp, act)
}
