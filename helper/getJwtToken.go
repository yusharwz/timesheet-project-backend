package helper

import (
	"fmt"
	"timesheet-app/middleware"
)

func GetTokenJwt(userId, name, email, roles string) (string, error) {

	token, err := middleware.GenerateTokenJwt(userId, name, email, roles, 720)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return token, nil
}
