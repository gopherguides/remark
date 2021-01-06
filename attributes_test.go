package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Attributes_Set(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	at := &Attributes{}

	key := "foo"
	exp := "bar"
	at.Set(key, exp)

	act, ok := at.smap().Load(key)
	r.True(ok)
	r.Equal(exp, act)
}

func Test_Attributes_Add(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	at := &Attributes{}

	key := "foo"
	exp := "bar"
	at.Add(key, exp)

	act, ok := at.smap().Load(key)
	r.True(ok)
	r.Equal(exp, act)

	exp2 := "baz"
	at.Add(key, exp2)

	act, ok = at.smap().Load(key)
	r.True(ok)
	r.Equal(exp+" "+exp2, act)
}

func Test_Attributes_Get(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	at := &Attributes{}

	key := "foo"
	_, ok := at.smap().Load(key)
	r.False(ok)

	exp := "bar"
	at.Add(key, exp)

	act, ok := at.smap().Load(key)
	r.True(ok)
	r.Equal(exp, act)
}

func Test_Attributes_Options(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	at := &Attributes{}

	key := "foo"
	exp := "bar"
	at.Set(key, exp)

	act, ok := at.smap().Load(key)
	r.True(ok)
	r.Equal(exp, act)

	opts := at.Options()
	r.Equal(exp, opts[key])
}
