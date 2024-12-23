package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/", Home)

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println("Here's error")
	}

}

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "All good"})
}
