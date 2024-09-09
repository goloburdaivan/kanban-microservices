package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserInProjectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetInt("userId")
		projectId, _ := strconv.Atoi(c.Param("id"))
		result, err := http.Get(fmt.Sprintf("http://localhost:8082/projects/%d/permissions?userId=%d", projectId, userId))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		defer result.Body.Close()

		if result.StatusCode != http.StatusOK {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "You are not in the project",
			})
			return
		}

		c.Next()
	}
}
