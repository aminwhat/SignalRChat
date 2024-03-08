package src

import (
	"log"
	src_controllers "vazir_hessab/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ApplicationInit() {
	log.Println("Application Init")

	defer log.Println("Application Init Done")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//? Testing...
	// s3Bucket := os.Getenv("S3_BUCKET")
	// secretKey := os.Getenv("SECRET_KEY")

	// now do something with s3 or whatever
}

func RouterInit(router *gin.Engine) {
	router.GET("/version", src_controllers.Get_Version)
}
