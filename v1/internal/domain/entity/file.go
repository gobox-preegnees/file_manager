package entity

type Hash string
type Size int64
type VirtualName string
type File struct {
	Identifier
	Path        Path
	Hash        Hash
	Size        Size
	ModTime     ModTime
	VirtualName VirtualName
}

func (f File) ToString() string { return "" }
