package meta

import (
	"net/http"
	"path/filepath"
)

type Informer struct {
	Dir    string
	Client *http.Client
	Info   *Info
}

func NewInformer(path string, c *http.Client) (*Informer, error) {
	dir := path
	if len(filepath.Ext(path)) > 0 {
		dir = filepath.Dir(path)
	}

	if c == nil {
		c = http.DefaultClient
	}

	snow := &Informer{
		Dir:    dir,
		Client: c,
	}

	info, err := snow.NewInfo(path)
	if err != nil {
		return nil, err
	}
	snow.Info = info

	return snow, nil
}
