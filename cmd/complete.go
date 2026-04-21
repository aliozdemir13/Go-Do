package cmd

import (
	"fmt"
	"strconv"

	"Go-Do/internal/todo"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete <id>",
	Short: "Mark a task as complete",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID: %s", args[0])
		}
		if err := myList.Complete(todo.TaskID(id)); err != nil {
			return err
		}
		if err := myList.SaveToFile(filename); err != nil {
			return err
		}
		fmt.Println(todo.Indigo("\n Task completed!"))
		fmt.Println(todo.StyledBar("OPEN TASKS "))
		myList.Display(false)
		return nil
	},
}
