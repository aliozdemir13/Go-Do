package cmd

import (
	"fmt"

	"Go-Do/internal/todo"

	"github.com/spf13/cobra"
)

var showDone bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show tasks (open by default)",
	RunE: func(cmd *cobra.Command, _ []string) error {
		todo.PrintHeader(cmd.OutOrStdout())
		todo.PrintProgress(cmd.OutOrStdout(), &myList)
		if showDone {
			fmt.Fprintln(cmd.OutOrStdout(), todo.StyledBar("COMPLETED TASKS "))
			myList.Display(cmd.OutOrStdout(), true)
		} else {
			fmt.Fprintln(cmd.OutOrStdout(), todo.StyledBar("OPEN TASKS "))
			myList.Display(cmd.OutOrStdout(), false)
		}
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVar(&showDone, "done", false, "Show completed tasks")
}
