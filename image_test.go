package remark

import "testing"

func Test_Image_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "image")
}
