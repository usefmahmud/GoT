package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use: "init",
	Short: "initialize your repository",
	RunE: func(cmd *cobra.Command, args []string) error {
    fmt.Println("initialized empty got repository")
		return nil
	},
}