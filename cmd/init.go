package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use: "init [directory]",
	Short: "initialize your repository",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Failed to get current working directory")
			return err
		}

		// if user give a folder name "git init folder-name", so we use it as root directory
		if( len(args) > 0) {
			wd = filepath.Join(wd, args[0])

			if _, err := os.Stat(wd); os.IsNotExist(err) {
				if MkdirErr := os.Mkdir(wd, 0755); MkdirErr != nil {
					return fmt.Errorf("failed to create directory %s", args[0])
				}
			}
		}
		

		gotDir := filepath.Join(wd, ".got")
		if _, err :=  os.Stat(gotDir); err == nil {
			return nil
		}

		if err := os.Mkdir(gotDir, 0755); err != nil {
			return fmt.Errorf("failed to initialize the repository")
		}

		return nil
	},
}