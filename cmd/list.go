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
		err := todo.PrintHeader(cmd.OutOrStdout())
		if err != nil {
			return err
		}
		e := todo.PrintProgress(cmd.OutOrStdout(), &myList)
		if e != nil {
			return e
		}
		if showDone {
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), todo.StyledBar("COMPLETED TASKS "))
			myList.Display(cmd.OutOrStdout(), true)
		} else {
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), todo.StyledBar("OPEN TASKS "))
			myList.Display(cmd.OutOrStdout(), false)
		}
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVar(&showDone, "done", false, "Show completed tasks")
}
