package meta

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Info struct {
	os.FileInfo
	Path string
	Dir  string
}

func (info Info) Join(paths ...string) string {
	paths = append([]string{info.Dir}, paths...)
	return filepath.Join(paths...)
}

func (info Info) MarshalJSON() ([]byte, error) {
	mm := map[string]interface{}{
		"path": info.Path,
		"dir":  info.Dir,
		"file_info": map[string]interface{}{
			"name":     info.Name(),
			"size":     info.Size(),
			"mode":     info.Mode(),
			"mod_time": info.ModTime(),
			"is_dir":   info.IsDir(),
		},
	}
	return json.MarshalIndent(mm, "", "  ")
}

func (info Info) String() string {
	b, _ := json.Marshal(info)
	return string(b)
}

func (i *Informer) NewInfo(path string) (*Info, error) {
	if strings.HasPrefix(path, "http") {
		return i.NewHTTPInfo(path)
	}

	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	inf, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if inf.IsDir() {
		return nil, fmt.Errorf("path must be file, not directory %q", path)
	}
	info := &Info{
		FileInfo: inf,
		Path:     path,
		Dir:      filepath.Dir(path),
	}

	return info, nil
}
