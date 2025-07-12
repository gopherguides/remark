package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Heading_Structure(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	heading := Heading{
		Generic: NewGeneric("h1"),
		Level:   1,
	}

	r.Equal(1, heading.Level)
	r.NotNil(heading.Generic)
	r.Equal("h1", heading.Generic.Name)
}

func Test_Heading_WithContent(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	heading := Heading{
		Generic: NewGeneric("h2"),
		Level:   2,
	}

	heading.Generic.Children = append(heading.Generic.Children, String("Chapter Title"))
	heading.Generic.Set("id", "chapter-1")

	r.Equal(2, heading.Level)
	r.Len(heading.Generic.Children, 1)
	r.Equal("Chapter Title", heading.Generic.Children[0].String())

	id, exists := heading.Generic.Get("id")
	r.True(exists)
	r.Equal("chapter-1", id)
}

func Test_Heading_DifferentLevels(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	testCases := []struct {
		tag   string
		level int
	}{
		{"h1", 1},
		{"h2", 2},
		{"h3", 3},
		{"h4", 4},
		{"h5", 5},
		{"h6", 6},
	}

	for _, tc := range testCases {
		heading := Heading{
			Generic: NewGeneric(tc.tag),
			Level:   tc.level,
		}

		r.Equal(tc.level, heading.Level)
		r.Equal(tc.tag, heading.Generic.Name)
	}
}
