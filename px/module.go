package px

import (
	"io/fs"
	"path/filepath"
)

type Modules []*Module

type Module struct {
	FS
	ModName string
}

func (m *Module) ReadFile(path string) ([]byte, error) {
	return fs.ReadFile(m.FS, path)
}

func (m *Module) Sub(path string) (FS, error) {
	return fs.Sub(m.FS, path)
}

func FindModules(cab FS) (Modules, error) {
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
			FS:      bac,
			ModName: "module.md",
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	return mods, nil
}
