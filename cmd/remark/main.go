package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gopherguides/remark/md"
	"github.com/gopherguides/remark/px"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// b, err := ioutil.ReadAll(os.Stdin)
	// if err != nil {
	// 	return err
	// }

	pwd := os.Getenv("MARKED_ORIGIN")
	// panic(pwd)
	// if len(pwd) == 0 {
	// 	pwd, _ = os.Getwd()
	// }

	cab := os.DirFS(pwd)

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
