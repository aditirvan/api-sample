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
	router.GET("/data", data)

	router.Run("0.0.0.0:8000")
}

func home(c *gin.Context) {
	hostname, _ := os.Hostname()
	var statuses = status{
		Status:  200,
		Message: fmt.Sprint("Running on ", hostname),
	}

	c.IndentedJSON(http.StatusOK, statuses)
}

func data(c *gin.Context) {
	hostname, _ := os.Hostname()
	var statuses = status{
		Status:  200,
		Message: fmt.Sprint("/data Running on ", hostname),
	}

	c.IndentedJSON(http.StatusOK, statuses)
}
