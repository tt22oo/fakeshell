package dir

import "sync"

type Entry struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Data     string            `json:"data"`
	Children map[string]*Entry `json:"children"`
}

type Dir struct {
	Mu      sync.Mutex
	Entries map[string]*Entry
}

const (
	File      string = "file"
	Directory string = "directory"
)
