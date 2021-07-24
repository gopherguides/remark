package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing/fstest"

	"github.com/gopherguides/remark/md"
	"github.com/gopherguides/remark/px"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	// pwd := os.Getenv("MARKED_ORIGIN")
	// if len(pwd) == 0 {
	// 	pwd, _ = os.Getwd()
	// }

	cab := fstest.MapFS{
		"module.md": {
			Data: b,
		},
	}

	mod := &px.Module{
		FS:      cab,
		ModName: "module.md",
	}

	p := md.NewParser(mod)

	tags, err := p.Parse()
	if err != nil {
		return err
	}

	fmt.Println(tags)
	return nil
}
