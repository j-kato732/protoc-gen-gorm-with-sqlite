FROM golang:1.17.0-buster

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64


RUN apt update && apt install -y git curl protobuf-compiler tree sqlite3
# protoc-gen-go @ v1.27.1
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
# v1.1.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

RUN go install github.com/infobloxopen/protoc-gen-gorm@latest

WORKDIR /grpc_gateway_sample
RUN go mod init grpc_gateway_sample
# grpc @ v1.40.0
RUN go get google.golang.org/grpc@latest

RUN go get github.com/grpc-ecosystem/grpc-gateway/v2

RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/sqlite

RUN git clone https://github.com/googleapis/googleapis.git ${GOPATH}/src

RUN git clone https://github.com/infobloxopen/protoc-gen-gorm.git ${GOPATH}/src/protoc-gen-gorm

RUN go mod tidy
