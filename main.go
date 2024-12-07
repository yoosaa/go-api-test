package github.com/yoosaa/go-api-test

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
		r.GET("ping/", func(c *gin.Context)) {
			
		}
}