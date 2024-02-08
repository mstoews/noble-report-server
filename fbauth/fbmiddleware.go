package fbauth

import (
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/mstoews/prd-backup-server/secrets"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

//

const (
	authorizationHeader = "Authorization"
	apiKeyHeader 		= "X-API-Key"
	cronExecutedHeader	= "X-Appengine-Cron"
	valName = "FIREBASE_ID_TOKEN"
)

func AuthJWT(client *auth.Client) gin.HandlerFunc {
	return func (c *gin.Context){
		startTime := time.Now()

		authHeader := c.Request.Header.Get(authorizationHeader)
		log.Println ("authHeader ", authHeader)
		token := strings.Replace(authHeader, "Bearer", "", 1)

		idToken, err := client.VerifyIDToken(c, token) // usually gets it from the cache

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H {
				"code": http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
		}

		log.Println("Auth Time", time.Since(startTime))

		c.Set(valName, idToken)
		c.Next()

		// Authorization Bearer ID_TOKEN
	}
}

// API key auth middleware
func AuthAPIKey(secretId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.Header.Get(apiKeyHeader)

		secret, err := secrets.GetSecret(secretId)
		if err != nil {
			log.Println("failed to get secret")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("comparing secret with provided key", secret, key)

		if secret != key {
			log.Println("key doesnt match!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("no error during check")
		c.Next()
	}
}

// AppEngine cron authentication
func AuthAppEngineCron() gin.HandlerFunc {
	return func(c *gin.Context) {
		cron := c.Request.Header.Get(cronExecutedHeader)

		if cron != "true" {
			log.Println("not invoked by cron - access denied")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		c.Next()
	}
}
