.SILENT:

run:
	go run cmd/main.go

test:
	go test ./... -v

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/grpc/service.proto