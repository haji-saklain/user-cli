package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	// Cobra root command
	var rootCmd = &cobra.Command{
		Use:   "user",
		Short: "User management CLI",
	}

	// Add subcommands to the root command
	rootCmd.AddCommand(user.GetUserByIDCmd)
	rootCmd.AddCommand(user.CreateUserCmd)
	rootCmd.AddCommand(user.UpdateUserCmd)
	rootCmd.AddCommand(user.DeleteUserCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
