package entity

type Identifier struct {
	Username string
	FolderID string
	ClientID string
}

type File struct {
	FileName    string
	HashSum     string
	FileSize    int64
	ModTime     int64
	VirtualName string
}
