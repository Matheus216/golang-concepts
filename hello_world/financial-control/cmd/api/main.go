package api

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/wallet")

	router.Run("localhost:9999")
}
