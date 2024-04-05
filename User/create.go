package user

import (
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
}

var userData map[int]User

func CreateUser(name, email string) {
	userID := len(userData) + 1
	user := User{
		ID:    userID,
		Name:  name,
		Email: email,
	}
	userData[userID] = user
	fmt.Println("User created successfully")
	return
}
func init() {
	userData = make(map[int]User)
}
