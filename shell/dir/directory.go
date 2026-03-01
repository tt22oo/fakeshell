package dir

func (e *Entry) NewDirectory(name string) {
	e.Children[name] = &Entry{
		Type: Directory,
	}
}
