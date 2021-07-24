package px

import (
	"io/fs"
	"os"
	"path/filepath"
)

type FS = fs.FS

func DirFS(path string) (FS, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	cab := os.DirFS(path)
	return cab, nil
}
