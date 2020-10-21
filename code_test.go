package remark

import (
	"testing"
)

func Test_Code_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "code")
}
