package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Image_Structure(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	image := Image{
		Generic: NewGeneric("img"),
	}

	r.NotNil(image.Generic)
	r.Equal("img", image.Generic.Name)
}

func Test_Image_WithAttributes(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	image := Image{
		Generic: NewGeneric("img"),
	}

	image.Generic.Set("src", "/path/to/image.png")
	image.Generic.Set("alt", "Test image")
	image.Generic.Set("width", "100")
	image.Generic.Set("height", "200")

	src, exists := image.Generic.Get("src")
	r.True(exists)
	r.Equal("/path/to/image.png", src)

	alt, exists := image.Generic.Get("alt")
	r.True(exists)
	r.Equal("Test image", alt)

	width, exists := image.Generic.Get("width")
	r.True(exists)
	r.Equal("100", width)

	height, exists := image.Generic.Get("height")
	r.True(exists)
	r.Equal("200", height)
}

func Test_Image_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	image := Image{
		Generic: NewGeneric("img"),
	}

	image.Generic.Set("src", "test.jpg")
	image.Generic.Set("alt", "Test")

	result := image.Generic.String()
	r.Contains(result, "<img")
	r.Contains(result, `src="test.jpg"`)
	r.Contains(result, `alt="Test"`)
}

func Test_Image_AsTag(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	image := Image{
		Generic: NewGeneric("img"),
	}

	// Test that Image can be used as a Tag through its Generic field
	var tag Tag = image.Generic
	r.NotNil(tag)
	r.Equal("img", tag.TagName())
}
