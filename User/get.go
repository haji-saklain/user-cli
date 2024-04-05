package User

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetUserByIDCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Get user by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Executing get user by ID command")
	},
}
