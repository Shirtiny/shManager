package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"shManager/model"
	"shManager/server"

	"github.com/gin-gonic/gin"
	serverlessplus "github.com/serverlessplus/go"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

const (
	dev = false
)

var handler *serverlessplus.Handler

func serverlessInit(r *gin.Engine, port int) {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverlessplus.Host, port))
	if err != nil {
		fmt.Printf("failed to listen on port %d: %v\n", port, err)
		// panic to force the runtime to restart
		panic(err)
	}
	go http.Serve(l, r)

	// setup handler
	types := []string{"image/png"}
	handler = serverlessplus.NewHandler(port).WithBinaryMIMETypes(types)
}

func entry(ctx context.Context, req *serverlessplus.APIGatewayRequest) (*serverlessplus.APIGatewayResponse, error) {
	return handler.Handle(ctx, req)
}

func main() {
	model.ConnectDatabase()
	router := server.CreateRouter("/shManager")
	if dev {
		router.Run(":" + "2021")
	} else {
		// start your server
		serverlessInit(router, 2021)
		cloudfunction.Start(entry)
	}

}

// GOOS=linux GOARCH=amd64 go build -o main main.go
