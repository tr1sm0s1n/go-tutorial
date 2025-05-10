package main

import (
	"embed"
	"fmt"
)

var (
	//go:embed slides/*
	dir embed.FS
	// Parsed markdown contents.
	data []string
	// Markdown files.
	slides = []string{
		"slides/what_go.md",
		"slides/why_go.md",
		"slides/how_go.md",
	}
)

func init() {
	for _, s := range slides {
		d, err := dir.ReadFile(s)
		if err != nil {
			panic(err)
		}

		data = append(data, fmt.Sprintf("\n%s\n", string(d)))
	}
}
