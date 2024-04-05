package user

import (
	"fmt"
)

// DeleteUser deletes the user with the given ID.
func DeleteUser(userID int) error {
	_, ok := userData[userID]
	if !ok {
		return fmt.Errorf("user not found")
	}
	delete(userData, userID)
	fmt.Println("User deleted successfully")
	return nil
}
