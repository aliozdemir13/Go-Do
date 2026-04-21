package cmd

import (
	"fmt"
	"strings"

	"Go-Do/internal/todo"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <title>",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := strings.Join(args, " ")
		myList.Add(title)
		if err := myList.SaveToFile(filename); err != nil {
			return err
		}
		fmt.Println(todo.Indigo("\n Task added: ") + title)
		fmt.Println(todo.StyledBar("OPEN TASKS "))
		myList.Display(false)
		return nil
	},
}
