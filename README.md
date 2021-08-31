# grpc_gateway_sample
grpc gateway sample

```
$ git clone git@github.com:j-kato732/grpc_gateway_sample.git
$ cd grpc_gateway_sample
$ docker-compose up -d
$ docker exec -ti grpc_gateway_sample bash
```
# proto変更した場合
```
protoc -I .:${GOPATH}/src --grpc-gateway_out .     --grpc-gateway_opt logtostderr=true     --grpc-gateway_opt paths=source_relative     api/proto/aimo.proto
```
