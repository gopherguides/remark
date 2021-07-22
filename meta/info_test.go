package meta

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewInfo_String(t *testing.T) {
	r := require.New(t)

	fp := "./info.go"

	i, err := NewInfo(fp)
	r.NoError(err)
	r.NotNil(i)

	act := i.String()
	// exp := `{"dir":"/Users/markbates/Dropbox/dev/remark/meta","file_info":{"is_dir":false,"mod_time":"2021-07-21T17:22:35.428582509-04:00","mode":420,"name":"info.go","size":947},"path":"/Users/markbates/Dropbox/dev/remark/meta/info.go"}`
	// r.Equal(exp, act)
	r.Contains(act, fmt.Sprintf("%q:%q", "dir", i.Dir))
	r.Contains(act, fmt.Sprintf("%q:%d", "size", i.Size()))
}

func Test_NewInfo(t *testing.T) {
	r := require.New(t)

	fp := "./info.go"

	i, err := NewInfo(fp)
	r.NoError(err)
	r.NotNil(i)

	r.True(strings.HasSuffix(i.Dir, filepath.Join("remark", "meta")))
	r.True(strings.HasSuffix(i.Path, filepath.Join("remark", "meta", "info.go")))
	r.NotNil(i.FileInfo)
	r.Equal("info.go", i.Name())

	i, err = NewInfo("badpath")
	r.Error(err)
	r.Nil(i)
}
