package router

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Initialize() {
	r := gin.Default()

	InitializeRoutes(r)

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080"
	}

	r.Run("0.0.0.0:" + port)
}
