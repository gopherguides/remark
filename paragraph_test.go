package remark

import (
	"testing"
)

func Test_Paragraph_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "paragraph")
}
