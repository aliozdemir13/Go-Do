package cmd

import (
	"fmt"

	"Go-Do/internal/todo"

	"github.com/spf13/cobra"
)

var (
	showDone bool
	showOpen bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show tasks (open by default)",
	RunE: func(cmd *cobra.Command, args []string) error {
		printHeader()
		printProgress(&myList)
		if showDone {
			fmt.Println(todo.StyledBar("COMPLETED TASKS "))
			myList.Display(true)
		} else {
			fmt.Println(todo.StyledBar("OPEN TASKS "))
			myList.Display(false)
		}
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVar(&showDone, "done", false, "Show completed tasks")
	listCmd.Flags().BoolVar(&showOpen, "open", false, "Show open tasks")
	listCmd.MarkFlagsMutuallyExclusive("done", "open")
}
