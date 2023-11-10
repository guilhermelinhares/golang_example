package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router := app.Group("/")
	router.GET("/", Hello)
	app.Run(":8080")

}
func Hello(c *gin.Context) {
	db_password := os.Getenv("DB_PASSWORD")
	db_user := os.Getenv("DB_USER")
	db := os.Getenv("DB_SOURCE")
	server_port := os.Getenv("SERVER_PORT")

	c.Writer.Write([]byte("Database Password:\n"))
	c.Writer.Write([]byte(db_password))
	c.Writer.Write([]byte("\nDatabase User:\n"))
	c.Writer.Write([]byte(db_user))
	c.Writer.Write([]byte("\nDatabase:\n"))
	c.Writer.Write([]byte(db))
	c.Writer.Write([]byte("\nPort:\n"))
	c.Writer.Write([]byte(server_port))
}
