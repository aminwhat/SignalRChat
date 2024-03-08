package main

import (
	"log"
	"os"
	"vazir_hessab/src"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func init() {
	src.ApplicationInit()
}

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(c socketio.Conn) error {
		log.Println("connected:", c.ID())
		return nil
	})

	server.OnEvent("/", "amin", func(s socketio.Conn, msg string) {

	})

	server.OnDisconnect("/", func(c socketio.Conn, s string) {
		log.Println("closed:", s)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	router := gin.New()
	router.Any("/v0", gin.WrapH(server))
	src.RouterInit(router)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
