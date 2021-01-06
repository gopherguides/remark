package remark

import (
	"strings"

	"github.com/gobuffalo/syncx"
	"github.com/gobuffalo/tags/v3"
)

type Attributes struct {
	data *syncx.StringMap
}

func (at *Attributes) smap() *syncx.StringMap {
	if at.data == nil {
		at.data = &syncx.StringMap{}
	}
	return at.data
}

func (at *Attributes) Set(key string, value string) {
	at.smap().Store(key, value)
}

func (at *Attributes) Add(key string, value string) {
	x, ok := at.smap().Load(key)
	if !ok || len(x) == 0 {
		at.Set(key, value)
		return
	}
	vals := []string{x, value}
	at.Set(key, strings.Join(vals, " "))
}

func (at *Attributes) Get(key string) (string, bool) {
	return at.smap().Load(key)
}

func (at *Attributes) Options() tags.Options {
	opts := tags.Options{}

	at.smap().Range(func(key string, value string) bool {
		opts[key] = value
		return true
	})

	return opts
}
