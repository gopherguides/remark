package md

import "testing"

func Test_Text_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "text")
}
