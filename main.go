package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pratyushanand26/web-scrapper/db"
	"github.com/pratyushanand26/web-scrapper/handlers"
)

func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	jwtsecret := os.Getenv("JWT_SECRET")
	db_password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("postgres://%s:localhost:5432/testdb?sslmode=disable", db_password)

	if port == "" {
		log.Fatal("port not found")
	}

	if jwtsecret == "" {
		log.Fatal("jwtsecret must be present in .env")
	}

	if db_password == "" {
		log.Fatal("port not found")
	}

	r := gin.Default()
	gormdb, err := db.New(dsn)

	if err != nil {
		log.Fatal("error is establishing connection with db")
	}

	api := r.Group("/api")

	api.POST("/register", func(c *gin.Context) {
		handlers.Register(c, gormdb)
	})

	r.Run(":" + port)
	fmt.Println("server started on port", port)

}
