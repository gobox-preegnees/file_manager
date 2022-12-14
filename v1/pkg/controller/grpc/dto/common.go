package grpc_dto

type Status int8
type Description string
type StandardResponse struct {
	Status
	Description
}

type Username string
type FolderID string
type ClientID string
type Identifier struct {
	Username
	FolderID
	ClientID
}

type Path string
type Hash string
type ModTime int64
type Size int64
type File struct {
	Identifier
	Path
	Hash
	ModTime
	Size
}

type Files []File

type Folder struct {
	Identifier
	Path
	ModTime
}

type Folders []Folder

type NewPath Path
type RenameInfo struct {
	Path
	NewPath
}
