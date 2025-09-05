package types

import (
	"fmt"
	"runtime"
)

type Data struct {
	Tasks  map[int]Task
	NextID int
}

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t Task) Display() {
	status := "not completed"
	if t.Completed {
		status = "completed"
	}
	fmt.Printf("Task %d: %s (%s)\n", t.ID, t.Title, status)
}

type Info struct {
	os, arch, version string
}

func InitializeInfo() Info {
	return Info{
		os:      runtime.GOOS,
		arch:    runtime.GOARCH,
		version: runtime.Version(),
	}
}

func (i Info) Display() {
	fmt.Printf("INFO (OS: %s, ARCH: %s, Version: %s)\n", i.os, i.arch, i.version)
}

type Displayer interface {
	Display()
}
