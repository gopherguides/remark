package meta

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strconv"
	"time"
)

func (i *Informer) NewHTTPInfo(path string) (*Info, error) {
	res, err := i.Client.Head(path)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("unsucessful status code %q %d", path, res.StatusCode)
	}

	hi := &httpInfo{
		name: filepath.Base(path),
	}

	mt, err := time.Parse("", res.Header.Get("Last-Modified"))
	if err != nil {
		return nil, err
	}
	hi.modTime = mt

	size, err := strconv.Atoi(res.Header.Get("Content-Length"))
	if err != nil {
		return nil, err
	}

	hi.size = int64(size)

	info := &Info{
		FileInfo: hi,
		Dir:      filepath.Dir(path),
		Path:     path,
	}

	return info, nil
}

var _ fs.FileInfo = &httpInfo{}

type httpInfo struct {
	name    string      // base name of the file
	size    int64       // length in bytes for regular files; system-dependent for others
	mode    fs.FileMode // file mode bits
	modTime time.Time   // modification time
}

func (h httpInfo) Name() string {
	return h.name
}

func (h httpInfo) Size() int64 {
	return h.size
}

func (httpInfo) IsDir() bool {
	return false
}

func (h httpInfo) Mode() fs.FileMode {
	return h.mode
}

func (h httpInfo) ModTime() time.Time {
	return h.modTime
}

func (httpInfo) Sys() interface{} {
	return nil
}
