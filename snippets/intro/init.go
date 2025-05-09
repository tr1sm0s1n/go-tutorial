package main

import (
	"fmt"
	"os"
)

var (
	// Parsed markdown contents.
	data []string
	// Markdown files.
	slides = []string{
		"./slides/what_go.md",
		"./slides/why_go.md",
		"./slides/how_go.md",
	}
)

func init() {
	for _, s := range slides {
		d, err := os.ReadFile(s)
		if err != nil {
			panic(err)
		}

		data = append(data, fmt.Sprintf("\n%s\n", string(d)))
	}
}
