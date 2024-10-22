package middlewares

import (
	"github.com/gin-gonic/gin"
)

// func extractBearerToken(c *gin.Context) (string, error) {
// 	header := c.GetHeader("Authorization")
// 	if header == "" {
// 		return "", errors.New("No 'authorization' header found")
// 	}
// 	tokenStrings := strings.Split(header, " ")
// 	if len(tokenStrings) != 2 {
// 		return "", errors.New("incorrectly formatted authorization header")
// 	}
// 	return tokenStrings[1], nil
// }

// func parseTokenString(tokenString string) (*jwt.Token, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("bad signing method received")
// 		}
// 		return []byte(system.JWT_SECRET), nil
// 	})
// 	if err != nil {
// 		return nil, errors.New("bad jwt")
// 	}
// 	return token, nil
// }

func UserValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// jwtString, err := extractBearerToken(c)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		// 	return
		// }

		// token, err := parseTokenString(jwtString)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		// 	return
		// }

		// _, ok := token.Claims.(jwt.MapClaims)
		// if !ok {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "unable to parse claims"})
		// 	return
		// }

		// TODO extract claims and set into context

		c.Set("user-uuid", "1858f3ab-9cb3-4b15-96b6-c4407cc7cfc7")
		c.Next()
	}
}
