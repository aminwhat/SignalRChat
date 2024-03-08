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
	rootSocket()

	router.GET("/version", src_controllers.Get_Version)

	router.GET("/v0/*any", gin.WrapH(SocketioServer))
	router.POST("/v0/*any", gin.WrapH(SocketioServer))

}

func rootSocket() {
	SocketioServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	SocketioServer.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	SocketioServer.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		log.Println("chat:", msg)
		s.SetContext(msg)
		return "recv " + msg
	})

	SocketioServer.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	SocketioServer.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	SocketioServer.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})
}
