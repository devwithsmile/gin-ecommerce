package main

import (
	"devwithsmile/gin-ecommerce/internal/http"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	log.Printf("JWT_SECRET = [%s]", os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	http.Run()
}
