package src

import (
	"log"
	src_controllers "vazir_hessab/src/controllers"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
)

var SocketioServer *socketio.Server

func ApplicationInit() {
	log.Println("Application Init")

	defer log.Println("Application Init Done")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func RouterInit(router *gin.Engine) {
	router.GET("/version", src_controllers.Get_Version)

	router.GET("/v0/*any", gin.WrapH(SocketioServer))
	router.POST("/v0/*any", gin.WrapH(SocketioServer))
	hessabSocket()
}

func hessabSocket() {
	SocketioServer.OnConnect("/hessab", func(c socketio.Conn) error {
		log.Println("connected:", c.ID())
		return nil
	})

	SocketioServer.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	SocketioServer.OnEvent("/hessab", "amin", func(s socketio.Conn, msg string) {

	})

	SocketioServer.OnDisconnect("/hessab", func(c socketio.Conn, s string) {
		log.Println("closed:", s)
	})
}
