package entity

type Identifier struct {
	Username string
	FolderID string
	ClientID string
}

type File struct {
	Path        string
	Hash        string
	Size        int64
	ModTime     int64
	VirtualName string
}