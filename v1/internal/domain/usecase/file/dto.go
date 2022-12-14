package file_usecase

type Username string
type FolderID string
type ClientID string
type Identifier struct {
	Username
	FolderID
	ClientID
}

type VirtualName string
type Path string
type ModTime int64
type Size int64
type Hash string
type File struct {
	VirtualName
	Path
	ModTime
	Size
	Hash
}

type FullFile struct {
	Identifier
	File
}

type SaveFileDTO struct {
	FullFile
}

type GetFileByPathDTO struct {
	Identifier
	Path
}
