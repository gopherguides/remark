package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Attributes_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	attrs := Attributes{
		"class": "test",
		"id":    "myid",
	}

	result := attrs.String()
	// JSON marshaling may order keys differently, so we check it's valid JSON
	r.Contains(result, `"class":"test"`)
	r.Contains(result, `"id":"myid"`)
	r.Contains(result, `{`)
	r.Contains(result, `}`)
}

func Test_Attributes_Attrs(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	attrs := Attributes{
		"class": "test",
		"id":    "myid",
	}

	result := attrs.Attrs()
	r.Equal(attrs, result)
}

func Test_Attributes_Add(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	t.Run("add to empty attributes", func(t *testing.T) {
		attrs := Attributes{}
		attrs.Add("class", "test")
		r.Equal("test", attrs["class"])
	})

	t.Run("add to existing attribute", func(t *testing.T) {
		attrs := Attributes{
			"class": "existing",
		}
		attrs.Add("class", "new")
		r.Equal("existing new", attrs["class"])
	})

	t.Run("add to non-existing attribute", func(t *testing.T) {
		attrs := Attributes{
			"id": "test",
		}
		attrs.Add("class", "new")
		r.Equal(" new", attrs["class"]) // Add() joins with space, so empty + "new" = " new"
		r.Equal("test", attrs["id"])
	})
}

func Test_Attributes_Get(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	attrs := Attributes{
		"class": "test",
		"id":    "myid",
	}

	t.Run("get existing attribute", func(t *testing.T) {
		value, ok := attrs.Get("class")
		r.True(ok)
		r.Equal("test", value)
	})

	t.Run("get non-existing attribute", func(t *testing.T) {
		value, ok := attrs.Get("nonexistent")
		r.False(ok)
		r.Equal("", value)
	})
}

func Test_Attributes_Options(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	attrs := Attributes{
		"class": "test",
		"id":    "myid",
	}

	opts := attrs.Options()
	r.Equal("test", opts["class"])
	r.Equal("myid", opts["id"])
	r.Len(opts, 2)
}
