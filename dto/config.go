package dto

import "github.com/dgrijalva/jwt-go"

type (
	ConfigData struct {
		DbConfig    DbConfig
		AppConfig   AppConfig
		AdminConfig AdminConfig
	}

	DbConfig struct {
		Host        string
		DbPort      string
		User        string
		Pass        string
		Database    string
		MaxIdle     int
		MaxConn     int
		MaxLifeTime string
		LogMode     int
	}

	AppConfig struct {
		Port string
	}

	AdminConfig struct {
		Email    string
		Password string
	}

	JwtClaim struct {
		jwt.StandardClaims
		Id       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Roles    string `json:"role"`
	}
)
