package grpc_dto

type CreateFolderReqDTO struct {
	Identifier
	Folder
}

type RenameFolderReqDTO struct {
	Identifier
	RenameInfo
	ModTime
}

type DeleteFolderReqDTO struct {
	Identifier
	Path
}

type RecoverFolderReqDTO struct {
	Identifier
	Path
}

type GetDirSchemaReqDTO struct {
	Identifier
}

type GetDirSchemaRespDTO struct {
	StandardResponse
	Folder
	Files
}

type GetFolderReqDTO struct {
	Identifier
	Path
}

type GetFolderRespDTO struct {
	StandardResponse
	Files
}
