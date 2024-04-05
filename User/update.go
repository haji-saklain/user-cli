package User

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateUserCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update user by ID",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		
		fmt.Println("Executing update user command")
	},
}
