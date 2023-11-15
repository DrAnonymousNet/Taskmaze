package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "taskMaze",
	Short: "A simple task manager",
	Long: "A simple task manager that allows you to create tasks todo in the command line. It also allows you to mark tasks as done and delete them.",
}

