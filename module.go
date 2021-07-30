package remark

import (
	"io/fs"
	"path/filepath"

	"github.com/markbates/fsx"
)

type Modules []*Module

type Module struct {
	*fsx.FS
	Root string
	Name string
}

func (m *Module) ReadFile(path string) ([]byte, error) {
	return fs.ReadFile(m.FS, path)
}

func FindModules(cab fs.FS) (Modules, error) {
	var mods Modules
	err := fs.WalkDir(cab, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == "." {
			return nil
		}

		base := filepath.Base(path)
		if base != "module.md" {
			return nil
		}

		dir := filepath.Dir(path)
		bac, err := fs.Sub(cab, dir)
		if err != nil {
			return err
		}

		mods = append(mods, &Module{
			FS:   fsx.NewFS(bac),
			Name: "module.md",
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	return mods, nil
}
