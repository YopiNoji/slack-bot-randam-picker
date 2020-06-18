package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	// "strings"
	// "time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.POST("/randam-pick", func(c *gin.Context) {
		list := []string{"<@U013MCT7PS6>", "<!here>", "<!channel>", "<!everyone>"}
		// randamize order of list
		shuffle(list)
		r0 := list[0]
		r1 := list[1]
		r2 := list[2]

		c.JSON(http.StatusOK, gin.H{
			"text":          "Hello," + r0 + "," + r1 + "," + r2,
			"response_type": "in_channel",
		})
	})
	router.Run(":" + port)
}

func shuffle(data []string) {
    n := len(data)
    for i := n - 1; i >= 0; i-- {
        j := rand.Intn(i + 1)
        data[i], data[j] = data[j], data[i]
    }
}
