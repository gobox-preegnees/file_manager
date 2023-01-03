package grpc

import (
	"context"

	pb "github.com/gobox-preegnees/file_manager/api/grpc"
	"github.com/gobox-preegnees/file_manager/internal/domain"
	"github.com/gobox-preegnees/file_manager/internal/domain/entity"
)

func (s *server) GetFiles(ctx context.Context, in *pb.GetFilesReq) (*pb.GetFilesResp, error) {

	files, err := s.fileService.GetFiles(ctx, domain.GetFilesReqDTO{
		Identifier: entity.Identifier{
			Username: in.Identifier.Username,
			Folder:   in.Identifier.Folder,
		},
		OwnerId: int(in.OwnerId),
		FileId:  int(in.FileId),
	})
	if err != nil {
		return nil, err
	}

	fullFiles := make([]*pb.FullFile, len(files))
	for _, file := range files {
		fullFile := pb.FullFile{
			FileId:      int64(file.FileId),
			Removed:     file.Removed,
			State:       int32(file.State),
			VirtualName: file.VirtualName,
			OwnerId:     int64(file.OwnerId),
			File: &pb.File{
				HashSum:  file.HashSum,
				FileName: file.FileName,
				SizeFile: file.SizeFile,
				ModTime:  file.ModTime,
				Client:   file.Client,
			},
		}
		fullFiles = append(fullFiles, &fullFile)
	}
	return &pb.GetFilesResp{
		FullFiles: fullFiles,
	}, nil
}

func (s *server) SaveFile(ctx context.Context, in *pb.SaveFileReq) (*pb.SaveFileResp, error) {

	id, err := s.fileService.SaveFile(ctx, domain.SaveFileReqDTO{
		Identifier: entity.Identifier{
			Username: in.Identifier.Username,
			Folder:   in.Identifier.Folder,
		},
		File: entity.File{
			Client:   in.File.Client,
			FileName: in.File.FileName,
			HashSum:  in.File.HashSum,
			SizeFile: in.File.SizeFile,
			ModTime:  in.File.ModTime,
		},
		Client: in.File.Client,
	})
	return &pb.SaveFileResp{
		Id: int64(id),
	}, err
}

func (s *server) DeleteFile(ctx context.Context, in *pb.DeleteFileReq) (*pb.DeleteFileResp, error) {

	err := s.fileService.DeleteFile(ctx, domain.DeleteFileReqDTO{
		Identifier: entity.Identifier{
			Username: in.Identifier.Username,
			Folder:   in.Identifier.Folder,
		},
		Client:   in.FileName,
		FileName: in.Client,
	})
	return nil, err
}

func (s *server) RenameFile(ctx context.Context, in *pb.RenameFileReq) (*pb.RenameFileResp, error) {

	err := s.fileService.RenameFile(ctx, domain.RenameFileReqDTO{
		Identifier: entity.Identifier{
			Username: in.Identifier.Username,
			Folder:   in.Identifier.Folder,
		},
		OldFilName:  in.OldFileName,
		NewFileName: in.NewFileName,
		Client:      in.Client,
	})
	return nil, err
}
