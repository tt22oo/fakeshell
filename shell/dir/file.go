package dir

func (e *Entry) NewFile(name string, data string) {
	e.Children[name] = &Entry{
		Type: TypeFile,
		Data: data,
	}
}
