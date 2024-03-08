package main

import (
	"log"
	"os"
	"vazir_hessab/src"

	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	src.ApplicationInit()

	router := gin.New()
	src.SocketioServer = socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&websocket.Transport{Subprotocols: []string{"websocket"}},
		},
	})
	go func() {
		if err := src.SocketioServer.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	defer src.SocketioServer.Close()

	src.RouterInit(router)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3779"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
