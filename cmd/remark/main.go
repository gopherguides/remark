package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gopherguides/remark/md"
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

	path := os.Getenv("MARKED_PATH")
	if len(path) == 0 {
		args := os.Args[1:]
		if len(args) == 0 {
			return fmt.Errorf("file name missing")
		}
		path = args[0]
	}

	tags, err := md.Parse(path, b)
	if err != nil {
		return err
	}

	fmt.Println(tags)
	return nil
}
