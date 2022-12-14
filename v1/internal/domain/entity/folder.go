package entity

type Folder struct {
	Identifier
	Path    Path
	ModTIme ModTime
}

func (f Folder) ToString() string { return "" }
