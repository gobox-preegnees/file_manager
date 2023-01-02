package entity

type Identifier struct {
	Username string
	Folder string
}

type File struct {
	Client string
	FileName    string
	HashSum     string
	SizeFile    int64
	ModTime     int64
}
