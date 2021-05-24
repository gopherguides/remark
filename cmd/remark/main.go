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

	pwd := os.Getenv("MARKED_ORIGIN")
	if len(pwd) == 0 {
		pwd, _ = os.Getwd()
	}

	tags, err := md.Parse(pwd, b)
	if err != nil {
		return err
	}

	fmt.Println(tags)
	return nil
}
