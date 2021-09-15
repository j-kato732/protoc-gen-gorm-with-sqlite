# grpc_postgres_sample
```
$ git clone git@github.com:j-kato732/grpc_gateway_sample.git
$ cd grpc_gateway_sample
$ docker-compose up -d
$ docker exec -ti grpc_gateway_sample bash
```
# proto変更した場合
```
// serverの生成
$ protoc -I .:${GOPATH}/src --go_out ./ --go_opt paths=source_relative     --go-grpc_out ./ --go-grpc_opt paths=source_relative proto/aimo.proto
```
```
// gatewayの生成
$ protoc -I .:${GOPATH}/src --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative proto/aimo.proto
```

```
// gormモデルの生成
$ protoc -I .:${GOPATH}/src --gorm_out . --gorm_opt paths=source_relative proto/aimo.proto
```

// happy set
```
protoc -I .:${GOPATH}/src --go_out ./ --go_opt paths=source_relative     --go-grpc_out ./ --go-grpc_opt paths=source_relative --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --gorm_out . --gorm_opt paths=source_relative proto/aimo.proto
```

# connect database in container
```
$ sqlite3 db/testdb
```
# show table
```
$ .table
```
# sqlite web(コンテナ用ポートを開けていないのでホスト側から共有しているDBファイルを指定する)
pythonが入っていない場合はpythonを入れる
```
brew install python3
```
```
$ pip3 install sqlite-web
$ sqlite_web /path/to/database.db
```

# Create Table
```
# create table table_name
# (columm1 data_type const,
# columm2 data_type const,
# columm3 data_type const,
# primary key(columm1));
```
```
create table periods
(id integer not null,
period char(6) not null,
primary key(id));
```
