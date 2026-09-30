package cmd

import (
	"fmt"

	"Go-Do/internal/todo"

	"github.com/spf13/cobra"
)

var addTagCmd = &cobra.Command{
	Use:   "create <tagName> <tagColor>",
	Short: "Create a new tag",
	Args:  cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		tagName := args[0]
		tagColour := args[1]

		myTags.CreateTag(tagName, tagColour)

		if err := myTags.SaveToFile(filenameTags); err != nil {
			return err
		}

		_, _ = fmt.Fprintln(
			cmd.OutOrStdout(),
			todo.StyleTextWithTagName(
				tagColour,
				"\n Tag added: "+tagName,
			),
		)

		if err := todo.PrintTags(cmd.OutOrStdout(), &myTags); err != nil {
			return err
		}

		_, _ = fmt.Fprintln(
			cmd.OutOrStdout(),
			todo.StyledBar("OPEN TASKS "),
		)

		myList.Display(cmd.OutOrStdout(), false)

		return nil
	},
}
