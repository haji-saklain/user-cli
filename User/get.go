package user

import (
	"fmt"
)

func GetUserByID(userID int) (User, error) {
	user, ok := userData[userID]
	if !ok {
		return User{}, fmt.Errorf("user not found")
	}
	return user, nil
}
