package repo

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
type Owner struct {
	Identifier
	OwnerId int
}
type SaveFileReqDTO struct {
	Identifier
	File
	Client string
}
type SetStateReqDTO struct {
	Identifier
	FileName    string
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
	Client   string
	FileName string
}
type FindAllOwnersByUsernameDTO struct {
	Username string
}
type FindAllFilesByOwnerOrFileIdReqDTO struct {
	Owner
	FileId int
}
type FullFile struct {
	Owner
	File
	FileId      int
	Removed     bool
	VirtualName string
	State       int
	Client      string
}
type FindAllFilesByOwnerOrFileIdRespDTO struct {
	Files []FullFile
}
type SaveOwnerDTO struct {
	Identifier
}
type RenameOwnerDTO struct {
	OwnerID int
	NewName string
}
type DeleteOwnerDTO struct {
	OwnerID int
}
type FindAllOwnersReqDTO struct {
	Username string
}
type FindAllOwnersRespDTO struct {
	Owners []Owner
}
