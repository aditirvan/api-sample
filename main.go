package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type status struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/", home)

	router.Run("localhost:8000")
}

func home(c *gin.Context) {
	hostname, _ := os.Hostname()
	var statuses = status{
		Status:  200,
		Message: fmt.Sprint("Running on ", hostname),
	}

	c.IndentedJSON(http.StatusOK, statuses)
}
