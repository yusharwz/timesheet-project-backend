package helper

import (
	"final-project-enigma/middleware"
	"fmt"
)

func GetTokenJwt(userId, username, email, roles string) (string, error) {

	token, err := middleware.GenerateTokenJwt(userId, username, email, roles, 720)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return token, nil
}
