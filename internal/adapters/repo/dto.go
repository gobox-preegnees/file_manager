package repo

type FileDTO struct {
	Client      string
	FileName    string
	HashSum     string
	SizeFile    int64
	ModTime     int64
	VirtualName string
	State       int
	Removed     bool
}

type File struct {
	FileName string
	HashSum  string
	SizeFile int64
	ModTime  int64
}

type Identifier struct {
	Username string
	Folder   string
}

type SaveFileReqDTO struct {
	Identifier
	File
	Client string
}

type SetStateReqDTO struct {
	Identifier
	File
	VirtualName string
	State       int
}

type RenameFileReqDTO struct {
	Identifier
	Client  string
	OldName string
	NewName string
}

type DeleteFileReqDTO struct {
	Identifier
	Client   string
	FileName string
}

type RestoreFileReqDTO struct {
	Identifier
	FileName string
}
