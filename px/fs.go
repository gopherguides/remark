package px

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	_ fmt.Stringer   = &FS{}
	_ fs.FS          = &FS{}
	_ fs.StatFS      = &FS{}
	_ fs.ReadFileFS  = &FS{}
	_ json.Marshaler = &FS{}
)

type FS struct {
	fs.FS
}

func (cab FS) Exists(name string) bool {
	if _, err := cab.Stat(name); err != nil {
		return false
	}
	return true
}

func (cab FS) ReadFile(name string) ([]byte, error) {
	if rfs, ok := cab.FS.(fs.ReadFileFS); ok {
		return rfs.ReadFile(name)
	}
	return fs.ReadFile(cab, name)
}

func (cab FS) Sub(path string) (*FS, error) {
	kid, err := fs.Sub(cab.FS, path)
	if err != nil {
		return nil, err
	}
	return &FS{
		FS: kid,
	}, nil
}

func (cab FS) MarshalJSON() ([]byte, error) {
	paths, err := Paths(cab)
	if err != nil {
		return nil, err
	}
	return json.Marshal(paths)
}

func Paths(cab fs.FS) ([]string, error) {
	var paths []string
	err := fs.WalkDir(cab, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		base := filepath.Base(path)
		if strings.HasPrefix(base, ".") {
			return nil
		}
		paths = append(paths, path)
		return nil
	})
	return paths, err
}

func (cab FS) Stat(path string) (fs.FileInfo, error) {
	if sfs, ok := cab.FS.(fs.StatFS); ok {
		return sfs.Stat(path)
	}

	f, err := cab.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Stat()
}

func (cab FS) String() string {
	b, err := json.Marshal(cab)
	if err != nil {
		return fmt.Sprintf("can not marshal cab %v %s", cab.FS, err)
	}
	return string(b)
}

func DirFS(path string) (*FS, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	cab := os.DirFS(path)
	return &FS{
		FS: cab,
	}, nil
}

func NewFS(cab fs.FS) *FS {
	return &FS{
		FS: cab,
	}
}
