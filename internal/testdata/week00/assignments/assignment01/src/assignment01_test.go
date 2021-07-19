package assignment01

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_System(t *testing.T) {
	// section: one
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		t.Fatal("Your GOPATH has not been set!")
	}
	// section: one

	// section: two
	path := os.Getenv("PATH")
	gobin := filepath.Join(gopath, "bin")
	if !strings.Contains(path, gobin) {
		t.Fatalf("Your PATH does not contain %s", gobin)
	}

	log.Println("Success!")
	// section: two
}
