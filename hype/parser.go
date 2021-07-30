package hype

import (
	"github.com/markbates/fsx"
)

type Parser struct {
	*fsx.FS
	Constructors
	Decorators
}

func NewParser(cab *fsx.FS) (*Parser, error) {
	p := &Parser{
		Constructors: Constructors{},
		Decorators:   Decorators{},
		FS:           cab,
	}
	return p, nil
}
