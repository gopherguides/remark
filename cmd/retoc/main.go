package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/md"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	root := "/Users/markbates/dev/canmark/_courses/intro-to-go/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		base := filepath.Base(path)
		if base != "module.md" {
			return nil
		}

		if err := print(path); err != nil {
			return err
		}

		return nil
	})

	return err
}

func print(path string) error {
	fmt.Println()
	tags, err := md.ParseFile(path)
	if err != nil {
		return err
	}

	for i, t := range tags {
		s, ok := t.(*md.Section)
		if !ok {
			continue
		}

		if i == 0 {
			fmt.Println(s.Title)
		}

		kids := s.GetChildren()
		for ii, k := range kids {
			h, ok := k.(*remark.Heading)
			if !ok {
				continue
			}

			if ii == 0 && i == 0 {
				continue
			}

			title := h.Children.String()

			sb := strings.Builder{}
			for i := 0; i < h.Level; i++ {
				sb.WriteString("\t")
			}
			sb.WriteString(title)
			fmt.Println(sb.String())
		}

	}
	return nil
}
