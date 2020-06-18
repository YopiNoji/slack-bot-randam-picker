package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	// "strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.POST("/randam-pick", func(c *gin.Context) {
		list := []string{"<@YopiNoji>", "<@here>", "<@channel>"}
		r1 := pickup(list)
		r2 := pickup(list)
		c.JSON(http.StatusOK, gin.H{
			"text":          "Hello," + r1 + " " + r2,
			"response_type": "in_channel",
		})
	})
	router.Run(":" + port)
}

func pickup(s []string) string {
    rand.Seed(time.Now().UnixNano())
    i := rand.Intn(len(s))
    return s[i]
}
