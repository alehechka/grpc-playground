generate_go:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--graphql_out=./go-server/grpc \
		--grpc-gateway_out=. \
		--grpc-gateway_opt=logtostderr=true \
    	--grpc-gateway_opt=paths=source_relative \
    	--grpc-gateway_opt=generate_unbound_methods=true \
		--openapiv2_out=. \
    	--openapiv2_opt=logtostderr=true \
		go-server/grpc/helloworld.proto

	mv go-server/grpc/helloworld/helloworld.graphql.go go-server/grpc/
	rmdir go-server/grpc/helloworld