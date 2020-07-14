package main

import (
	"log"
	"os"

	"github.com/cuongduong257/go-boilerplate/driver"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		host     = os.Getenv("MYSQL_HOST")
		port     = os.Getenv("MYSQL_PORT")
		user     = os.Getenv("MYSQL_USER")
		password = os.Getenv("MYSQL_PASSWORD")
		dbname   = os.Getenv("MYSQL_DB")
	)

	db := driver.Connect(host, port, user, password, dbname)

	if err := db.Database.DB().Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Success connect to database!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() //listen and serve on 0.0.0.0:8080
}
