install_tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 

	go install github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql@v0.22.0 

	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.10.3 
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.10.3

	# go install github.com/gogo/protobuf/proto@v1.3.2
	# go install github.com/gogo/protobuf/jsonpb@v1.3.2
	# go install github.com/gogo/protobuf/protoc-gen-gogo@v1.3.2
	# go install github.com/gogo/protobuf/gogoproto@v1.3.2

generate_go:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--graphql_out=./go-server/grpc/helloworld \
		--grpc-gateway_out=. \
		--grpc-gateway_opt=logtostderr=true \
    	--grpc-gateway_opt=paths=source_relative \
    	--grpc-gateway_opt=generate_unbound_methods=true \
		--openapiv2_out=. \
    	--openapiv2_opt=logtostderr=true \
		go-server/grpc/helloworld/helloworld.proto

run_grpc:
	go run go-server/greeter_server/main.go

run_graphql:
	go run go-server/graphql/main.go

run_rest:
	go run go-server/rest/main.go