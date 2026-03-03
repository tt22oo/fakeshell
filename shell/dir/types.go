package dir

import "sync"

type Entry struct {
	Mu       sync.Mutex
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Data     string            `json:"data"`
	Children map[string]*Entry `json:"children"`
}

type Directory struct {
	Mu       sync.Mutex
	HomePath string
	Entry    *Entry
	Entries  map[string]*Entry
}

const (
	TypeFile      string = "file"
	TypeDirectory string = "directory"
)
