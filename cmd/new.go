/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "to Create your solid project",
	RunE: createProject,
}

func createProject(cmd *cobra.Command, args []string) error {
	projectName := args[0]
	
	gitCmd := exec.Command("git", "clone", "--depth", "1", "-b", "template", "https://github.com/wrhz/Solid.git", projectName)

	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr

	err := gitCmd.Run()
	if err != nil {
		os.Exit(1)
	}

	fmt.Print("\nDo you want to run \"npm install\"?(y/n): ")

	var npm string

	if _, err := fmt.Scanln(&npm); err != nil {
		return err
	}

	if npm == "y" {
		npmCmd := exec.Command("npm", "install")

		npmCmd.Stdout = os.Stdout
		npmCmd.Stderr = os.Stderr
		npmCmd.Dir = projectName

		err := npmCmd.Run()
		if err != nil {
			os.Exit(1)
		}
	}

	fmt.Print("\nDo you want to run \"go mod init\"?(y/n): ")

	var init string

	if _, err := fmt.Scanln(&init); err != nil {
		return err
	}

	if init == "y" {
		initCmd := exec.Command("go", "mod", "init", projectName)

		initCmd.Stdout = os.Stdout
		initCmd.Stderr = os.Stderr
		initCmd.Dir = projectName

		err := initCmd.Run()
		if err != nil {
			os.Exit(1)
		}

		fmt.Print("\nDo you want to run \"go mod tidy\"?(y/n): ")

		var tidy string

		if _, err := fmt.Scanln(&tidy); err != nil {
			return err
		}

		if tidy == "y" {
			tidyCmd := exec.Command("go", "mod", "tidy")

			tidyCmd.Stdout = os.Stdout
			tidyCmd.Stderr = os.Stderr
			tidyCmd.Dir = projectName

			err := tidyCmd.Run()
			if err != nil {
				os.Exit(1)
			}
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)
}
