package User

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateUserCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Executing create user command")
	},
}
