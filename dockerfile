FROM golang:1.19.4

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod tidy

RUN go build -o file_manager ./cmd/main.go

CMD [ "./file_manager" ]

