package grpc_dto

type SaveFileReqDTO struct {
	Identifier
	File
}

type RenameFileReqDTO struct {
	Identifier
	RenameInfo
}

type DeleteFileReqDTO struct {
	Identifier
	Path
}

type RecoverFileReqDTO struct {
	Identifier
	Path
}

type GetFileReqDTO struct {
	Identifier
}

type GetFileRespDTO struct {
	StandardResponse
	File
}
