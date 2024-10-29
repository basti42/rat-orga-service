package middlewares

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/basti42/orga-service/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func extractBearerToken(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("No 'authorization' header found")
	}
	tokenStrings := strings.Split(header, " ")
	if len(tokenStrings) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}
	return tokenStrings[1], nil
}

func parseTokenString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad signing method received")
		}
		return []byte(utils.JWT_SECRET), nil
	})
	if err != nil {
		return nil, errors.New("bad jwt")
	}
	return token, nil
}

func UserValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		jwtString, err := extractBearerToken(c)
		if err != nil {
			log.Printf("error extracting token: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		token, err := parseTokenString(jwtString)
		if err != nil {
			log.Printf("error parsing token: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Printf("error getting token claims")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "unable to parse claims"})
			return
		}

		data := claims["data"].(map[string]any)

		userUUIDstring, ok := data["user_uuid"].(string)
		if !ok {
			log.Printf("missing: user-uuid")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "missing payload in jwt"})
			return
		}

		subscriptionType, ok := data["subscription_type"].(string)
		if !ok {
			log.Printf("missing: subscription type")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "missing payload in jwt"})
			return
		}

		userRole, ok := data["role"].(string)
		if !ok {
			log.Printf("missing: user-role")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "missing payload in jwt"})
			return
		}

		c.Set("user-uuid", userUUIDstring)
		c.Set("user-role", userRole)
		c.Set("subscription-type", subscriptionType)
		c.Next()
	}
}
