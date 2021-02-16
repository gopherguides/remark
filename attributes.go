package remark

import (
	"encoding/json"
	"strings"

	"github.com/gobuffalo/tags/v3"
)

type Attributable interface {
	Attrs() Attributes
}

type Attributes map[string]string

func (a Attributes) String() string {
	b, _ := json.Marshal(a)
	return string(b)
}

func (at Attributes) Attrs() Attributes {
	return at
}

func (at Attributes) Set(key string, value string) {
	at[key] = value
}

func (at Attributes) Add(key string, value string) {
	if len(at) == 0 {
		at.Set(key, value)
		return
	}

	vals := []string{at[key], value}
	at.Set(key, strings.Join(vals, " "))
}

func (at Attributes) Get(key string) (string, bool) {
	s, ok := at[key]
	return s, ok
}

func (at Attributes) Options() tags.Options {
	opts := tags.Options{}

	for k, v := range at {
		opts[k] = v
	}
	return opts
}
