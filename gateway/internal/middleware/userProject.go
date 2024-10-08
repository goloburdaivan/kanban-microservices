package middleware

import (
	"encoding/json"
	"fmt"
	"gateway/internal/cache"
	"gateway/internal/dto"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func UserInProjectMiddleware(cache *cache.RedisCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetInt("userId")
		projectId, _ := strconv.Atoi(c.Param("id"))
		cacheKey := fmt.Sprintf("user-%d-project-%d", userId, projectId)

		projectInfo, err := cache.Remember(cacheKey, time.Duration(60)*time.Second, func() interface{} {
			log.Println("Making http request to get project info")
			result, err := http.Get(fmt.Sprintf("http://localhost:8082/projects/%d/permissions?userId=%d", projectId, userId))
			if err != nil {
				return map[string]interface{}{
					"error":   true,
					"message": err.Error(),
					"status":  float64(http.StatusInternalServerError),
				}
			}

			defer result.Body.Close()

			if result.StatusCode != http.StatusOK {
				return map[string]interface{}{
					"error":   true,
					"message": "You are not in the project",
					"status":  float64(http.StatusForbidden),
				}
			}

			jsonBody, err := ioutil.ReadAll(result.Body)
			if err != nil {
				return map[string]interface{}{
					"error":   true,
					"message": "Failed to read body",
					"status":  float64(http.StatusInternalServerError),
				}
			}

			var project dto.PermissionsDTO
			if err := json.Unmarshal(jsonBody, &project); err != nil {
				return map[string]interface{}{
					"error":   true,
					"message": "Failed to parse JSON",
					"status":  float64(http.StatusInternalServerError),
				}
			}

			return map[string]interface{}{
				"status":  float64(http.StatusOK),
				"error":   false,
				"project": project,
			}
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		if projectInfoMap, ok := projectInfo.(map[string]interface{}); ok {
			if projectInfoMap["error"].(bool) {
				statusCode := int(projectInfoMap["status"].(float64))
				c.AbortWithStatusJSON(statusCode, gin.H{
					"success": false,
					"message": projectInfoMap["message"].(string),
				})
				return
			}

			c.Set("project", projectInfoMap["project"])
		}

		c.Next()
	}
}

func forbiddenResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"success": false,
		"message": "You are not in the project",
	})
}
