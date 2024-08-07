package middleware

import (
	"errors"
	"final-project-enigma/dto"
	"final-project-enigma/dto/response"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	applicationName  = "timesheet-app"
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignatureKey  []byte
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	jwtSignatureKey = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateTokenJwt(Id, name, email, roles string, expiredAt int64) (string, error) {
	loginExpDuration := time.Duration(expiredAt) * time.Hour
	issuedAt := time.Now()
	myExpiresAt := issuedAt.Add(loginExpDuration).Unix()
	claims := dto.JwtClaim{
		Id:       Id,
		Username: name,
		Email:    email,
		Roles:    roles,
		StandardClaims: jwt.StandardClaims{
			Issuer:    applicationName,
			ExpiresAt: myExpiresAt,
			IssuedAt:  issuedAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwtSigningMethod, claims)

	signedToken, err := token.SignedString(jwtSignatureKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func JwtAuthWithRoles(userId ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response.NewResponseUnauthorized(c, "Invalid token")
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		claims := &dto.JwtClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSignatureKey, nil
		})

		if err != nil {
			response.NewResponseUnauthorized(c, "Invalid token")
			c.Abort()
			return
		}

		if !token.Valid {
			response.NewResponseUnauthorized(c, "Invalid token")
			c.Abort()
			return
		}

		// utils role
		validRole := false
		if len(userId) > 0 {
			for _, role := range userId {
				if role == claims.Roles {
					validRole = true
					break
				}
			}
		}

		if !validRole {
			response.NewResponseUnauthorized(c, "Invalid token")
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetIdFromToken(authHeader string) (string, error) {
	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return "", errors.New("invalid authorization format")
	}
	tokenString := splitToken[1]

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your_secret_key"), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to get claims")
	}
	id, ok := claims["id"].(string)
	if !ok {
		return "", errors.New("failed to get ID from token")
	}
	return id, nil
}
