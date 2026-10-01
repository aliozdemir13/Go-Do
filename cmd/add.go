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

		if err := myList.Add(title, tag, &myTags); err != nil {
			return err
		}

		if err := myList.SaveToFile(filename); err != nil {
			return err
		}

		err := todo.PrintHeader(cmd.OutOrStdout())
		if err != nil {
			return err
		}

		e := todo.PrintProgress(cmd.OutOrStdout(), &myList)
		if e != nil {
			return e
		}

		myTags.LoadFromFile(filenameTags)
		if err := todo.PrintTags(cmd.OutOrStdout(), &myTags); err != nil {
			return err
		}

		_, _ = fmt.Fprintln(
			cmd.OutOrStdout(),
			todo.Indigo("\n Task added: ")+title+"\n",
		)

		_, _ = fmt.Fprintln(
			cmd.OutOrStdout(),
			todo.StyledBar("OPEN TASKS "),
		)

		myList.Display(cmd.OutOrStdout(), false, "")

		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(
		&tag,
		"tag",
		"t",
		"",
		"Tag to assign to the task",
	)
}
