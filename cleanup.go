// Script to remove AsciiDoctor comments from source code listings when brought over
// from the book.
package main

import (
	"go/format"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	taggedRegion = regexp.MustCompile(`\n\s*//\s?(tag|end)::.*`)
	callout      = regexp.MustCompile(`//\s*<\d+>`)
	nothing      = []byte("")
)

func main() {
	matches, err := files(".")
	failOnError(err)

	for _, name := range matches {
		src, err := os.ReadFile(name)
		failOnError(err)

		// remove comments used by AsciiDoctor
		src = taggedRegion.ReplaceAllLiteral(src, nothing)
		src = callout.ReplaceAllLiteral(src, nothing)

		// run gofmt to clean up the result
		src, err = format.Source(src)
		failOnError(err)

		// write it back out
		failOnError(os.WriteFile(name, src, 0644))
	}
}

// files recursively finds all *.go files.
func files(root string) ([]string, error) {
	matches := make([]string, 0, 100)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// skip hidden folders like .git
			if path != "." && strings.HasPrefix(path, ".") {
				return filepath.SkipDir
			}
		} else {
			if filepath.Ext(path) == ".go" {
				matches = append(matches, path)
			}
		}
		return nil
	})
	return matches, err
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
