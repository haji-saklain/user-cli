package User

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteUserCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete user by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Executing delete user command")
	},
}
