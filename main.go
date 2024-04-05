package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type user struct {
	ID    int
	Name  string
	Email string
}

var userData map[int]user

func init() {
	userData = make(map[int]user)
}

func User() {
	var rootCmd = &cobra.Command{
		Use:   "user",
		Short: "User management CLI",
	}

	// Executing the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	create := flag.NewFlagSet("create", flag.ExitOnError)
	get := flag.NewFlagSet("get", flag.ExitOnError)
	update := flag.NewFlagSet("update", flag.ExitOnError)
	delete := flag.NewFlagSet("delete", flag.ExitOnError)

	createName := create.String("name", "", "User name")
	createEmail := create.String("email", "", "User email")

	getID := get.Int("id", 0, "User ID")

	updateID := update.Int("id", 0, "User ID")
	updateName := update.String("name", "", "User name")
	updateEmail := update.String("email", "", "User email")

	deleteID := delete.Int("id", 0, "User ID")

	if len(os.Args) < 2 {
		fmt.Println("Please specify an operation (create, get, update, delete)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		create.Parse(os.Args[2:])
		if *createName == "" || *createEmail == "" {
			create.PrintDefaults()
			os.Exit(1)
		}
		user.CreateUser(*createName, *createEmail)
	case "get":
		get.Parse(os.Args[2:])
		if *getID == 0 {
			get.PrintDefaults()
			os.Exit(1)
		}
		user, err := user.GetUserByID(*getID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("User ID: %d\nName: %s\nEmail: %s\n", user.ID, user.Name, user.Email)
	case "update":
		update.Parse(os.Args[2:])
		if *updateID == 0 || *updateName == "" || *updateEmail == "" {
			update.PrintDefaults()
			os.Exit(1)
		}
		err := user.UpdateUser(*updateID, *updateName, *updateEmail)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("User updated successfully")
	case "delete":
		delete.Parse(os.Args[2:])
		if *deleteID == 0 {
			delete.PrintDefaults()
			os.Exit(1)
		}
		err := user.DeleteUser(*deleteID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("User deleted successfully")
	default:
		fmt.Println("Invalid operation. Please specify create, get, update, or delete.")
		os.Exit(1)
	}
}

// Function to parse ID from string to integer
func parseID(id string) int {
	var userID int
	fmt.Sscanf(id, "%d", &userID)
	return userID
}
