package main

import (
	greeter "github.com/alehechka/grpc-playground/go-server/grpc/helloworld"
	"github.com/alehechka/grpc-playground/go-server/utils"
	"github.com/friendsofgo/graphiql"
	"github.com/gin-gonic/gin"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
)

func main() {
	mux, err := createGraphQLHandler()
	utils.Check(err)

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	utils.Check(err)

	engine := gin.Default()
	engine.POST("/graphql", gin.WrapH(mux))
	engine.GET("/graphiql", gin.WrapH(graphiqlHandler))
	utils.Check(engine.Run())
}

func createGraphQLHandler() (mux *runtime.ServeMux, err error) {
	mux = runtime.NewServeMux()

	if err = greeter.RegisterGreeterGraphql(mux); err != nil {
		return nil, err
	}

	return
}
