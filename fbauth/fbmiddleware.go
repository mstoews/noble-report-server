package fbauth

import (
	"errors"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/mstoews/prd-backup-server/secrets"
	"log"
	"net/http"
	"strings"
	"time"
)

//

const (
	authorizationHeader    = "Authorization"
	apiKeyHeader           = "X-API-Key"
	cronExecutedHeader     = "X-Appengine-Cron"
	valName                = "FIREBASE_ID_TOKEN"
	authorizationHeaderKey = "authorization"
)

func AuthJWT(client *auth.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		authHeader := ctx.Request.Header.Get(authorizationHeader)
		if len(authHeader) == 0 {
			err := errors.New("Authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		log.Println("authHeader ", authHeader)
		token := strings.Replace(authHeader, "Bearer", "", 1)

		//idToken, err := client.VerifyIDToken(ctx, token) // usually gets it from the cache
		//
		//if err != nil {
		//	ctx.JSON(http.StatusUnauthorized, gin.H{
		//		"code":    http.StatusUnauthorized,
		//		"message": http.StatusText(http.StatusUnauthorized),
		//	})
		//	err := errors.New("Authorization header is not provided")
		//	ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		//	return
		//}

		log.Println("Auth Time", time.Since(startTime))

		ctx.Set(valName, token)
		ctx.Next()

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

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
