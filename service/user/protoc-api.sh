
# make api 指令会自动执行此脚本
protoc --proto_path=./api \
       --proto_path=./third_party \
               --go_out=paths=source_relative:./api \
               --go-http_out=paths=source_relative:./api \
               --go-grpc_out=paths=source_relative:./api \
       --openapi_out=fq_schema_naming=true,default_response=false:. \
       api/user/v1/user.proto


# make config 指令会自动执行此脚本
protoc --proto_path=./internal \
       --proto_path=./third_party \
               --go_out=paths=source_relative:./internal \
       internal/conf/conf.proto