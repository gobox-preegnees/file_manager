package grpc

import (
	"context"

	pb "github.com/gobox-preegnees/file_manager/api/contract"
	"github.com/gobox-preegnees/file_manager/internal/domain/entity"
)

const CODE_OK = 100

func (s *server) DeleteFile(ctx context.Context, req *pb.DeleteFileReq) (*pb.StandardResponse, error) {

	err := s.fileUsecase.DeleteFile(
		ctx,
		entity.Identifier{
			Username: req.Identifier.Username,
			FolderID: req.Identifier.FolderID,
			ClientID: req.Identifier.ClientID,
		},
		req.Path,
		req.Hash,
	)
	if err != nil {
		return nil, err
	}
	return &pb.StandardResponse{
		Status: CODE_OK,
	}, nil
}

func (s *server) GetFiles(ctx context.Context, req *pb.GetFilesReq) (*pb.GetFilesResp, error) {

	files, err := s.fileUsecase.GetFiles(
		ctx,
		entity.Identifier{
			Username: req.Identifier.Username,
			FolderID: req.Identifier.FolderID,
			ClientID: req.Identifier.ClientID,
		},
	)
	if err != nil {
		return nil, err
	}

	pbFiles := make([]*pb.File, len(files))
	for _, file := range files {
		pbFiles = append(pbFiles, &pb.File{
			Path:        file.FileName,
			Hash:        file.HashSum,
			VirtualName: file.VirtualName,
			ModTime:     file.ModTime,
			Size:        file.SizeFile,
		})
	}

	return &pb.GetFilesResp{
		StandardResponse: &pb.StandardResponse{
			Status: CODE_OK,
		},
		File: pbFiles,
	}, nil
}

func (s *server) RenameFile(ctx context.Context, req *pb.RenameFileReq) (*pb.StandardResponse, error) {

	err := s.fileUsecase.RenameFile(ctx,
		entity.Identifier{
			Username: req.Identifier.Username,
			FolderID: req.Identifier.FolderID,
			ClientID: req.Identifier.ClientID,
		},
		req.RenameInfo.OldPath,
		req.RenameInfo.NewPath,
		req.RenameInfo.Hash,
	)
	if err != nil {
		return nil, err
	}

	return &pb.StandardResponse{
		Status: CODE_OK,
	}, nil
}

func (s *server) SaveFile(ctx context.Context, req *pb.SaveFileReq) (*pb.StandardResponse, error) {

	err := s.fileUsecase.SaveFile(
		ctx,
		entity.Identifier{
			Username: req.Identifier.Username,
			FolderID: req.Identifier.FolderID,
			ClientID: req.Identifier.ClientID,
		},
		entity.File{
			FileName:    req.File.Path,
			HashSum:     req.File.Hash,
			SizeFile:    req.File.Size,
			ModTime:     req.File.ModTime,
			VirtualName: req.File.VirtualName,
		},
	)
	if err != nil {
		return nil, err
	}

	return &pb.StandardResponse{
		Status: CODE_OK,
	}, nil
}
