package user

import (
	"fmt"
)

func UpdateUser(userID int, name, email string) error {
	user, ok := userData[userID]
	if !ok {
		return fmt.Errorf("user not found")
	}
	user.Name = name
	user.Email = email
	userData[userID] = user
	fmt.Println("User updated successfully")
	return nil
}
