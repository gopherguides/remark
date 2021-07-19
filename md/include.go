package md

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func (p *Parser) include(name string) (string, error) {
	fmt.Println("TODO >> include", name)
	base := filepath.Base(name)
	pn := "_" + base
	dir := filepath.Dir(name)
	fp := filepath.Join(dir, pn)
	fmt.Println("TODO >> include fp", fp)

	b, err := ioutil.ReadFile(fp)
	return string(b), err
}
