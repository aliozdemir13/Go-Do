// Package cmd provides the CLI commands for Go-Do
package cmd

import (
	"fmt"

	"Go-Do/internal/todo"

	"github.com/spf13/cobra"
)

const filename = "tasks.json"

var myList todo.TodoList

var rootCmd = &cobra.Command{
	Use:     "go-do",
	Short:   "A terminal-based task manager",
	Long:    `Go-Do is a CLI task manager that helps you manage your todos from the terminal.`,
	Version: "1.0.0",
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		return myList.LoadFromFile(filename)
	},
	RunE: func(cmd *cobra.Command, _ []string) error {
		err := todo.PrintHeader(cmd.OutOrStdout())
		if err != nil {
			return err
		}
		e := todo.PrintProgress(cmd.OutOrStdout(), &myList)
		if e != nil {
			return e
		}
		fmt.Fprintln(cmd.OutOrStdout(), todo.StyledBar("OPEN TASKS "))
		myList.Display(cmd.OutOrStdout(), false)
		return nil
	},
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(listCmd)
}
