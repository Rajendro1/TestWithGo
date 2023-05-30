package main

import (
	"log"
	"os"

	"github.com/Rajendro1/AccuKnox/Api/routers"
	pgdatabase "github.com/Rajendro1/AccuKnox/pgDatabase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("DB_PORT: ", os.Getenv("DB_PORT"))
	log.Println("DB_PASSWORD: ", os.Getenv("DB_PASSWORD"))
	log.Println("DB_USERNAME: ", os.Getenv("DB_USERNAME"))
	log.Println("DB_HOST: ", os.Getenv("DB_HOST"))
}
func main() {
	pgdatabase.Connect(&gin.Context{})
	routers.HandleRequest()
}
