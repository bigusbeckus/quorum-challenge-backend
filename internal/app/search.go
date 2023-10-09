package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchHandler(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing query parameter",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "hi fren " + query,
	})
}
