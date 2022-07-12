generate_go:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		grpc/helloworld.proto

	mkdir -p go-server/helloworld
	mv grpc/helloworld_grpc.pb.go go-server/helloworld/helloworld_grpc.pb.go
	mv grpc/helloworld.pb.go go-server/helloworld/helloworld.pb.go

generate_ts:
	protoc --grpc-gateway-ts_out=. \
		grpc/helloworld.proto

	mkdir -p client/src/api/helloworld
	mv fetch.pb.ts client/src/api/fetch.pb.ts
	mv grpc/helloworld.pb.ts client/src/api/helloworld/helloworld.pb.ts