package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use: "init",
	Short: "initialize your repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Failed to get current working directory")
			return err
		}

		gotDir := filepath.Join(wd, ".got")
    err = os.Mkdir(gotDir, 0755)
		if err != nil {
			fmt.Println("Failed to initialize the repository")
			return nil
		}

		return nil
	},
}