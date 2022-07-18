package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "github.com/alehechka/grpc-playground/go-server/grpc"
	"github.com/alehechka/grpc-playground/go-server/utils"
)

var (
	// gRPC server endpoint
	grpcServerEndpoint = utils.GetEnv("GRCP-SERVER-ENDPOINT", "localhost:80")
)

func main() {
	mux, err := createRESTHandler()
	utils.Check(err)

	engine := gin.Default()
	engine.Any("/*path", gin.WrapH(mux))
	utils.Check(engine.Run())
}

func createRESTHandler() (mux *runtime.ServeMux, err error) {
	ctx := context.Background()

	mux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err = gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		return nil, err
	}

	return
}
