.SILENT:

run:
	go run cmd/main.go

protov1:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative v1/pkg/contract/service.proto