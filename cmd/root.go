// Package cmd provides the CLI commands for Go-Do
package cmd

import (
	"fmt"
	"os"

	"Go-Do/internal/todo"

	"github.com/spf13/cobra"
)

const filename = "tasks.json"

var myList todo.TodoList

var rootCmd = &cobra.Command{
	Use:   "go-do",
	Short: "A terminal-based task manager",
	Long:  `Go-Do is a CLI task manager that helps you manage your todos from the terminal.`,
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		return myList.LoadFromFile(filename)
	},
	RunE: func(_ *cobra.Command, _ []string) error {
		printHeader()
		printProgress(&myList)
		fmt.Println(todo.StyledBar("OPEN TASKS "))
		myList.Display(false)
		return nil
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(listCmd)
}

func printHeader() {
	fmt.Print("\033[H\033[2J")
	fmt.Print(todo.MegaLogo())
	fmt.Printf("\n   %s %s %s\n",
		todo.Indigo("⚡ TASKS"),
		todo.Dim("v1.0.0"),
		todo.Indigo("●")+" "+todo.Dim("Local Storage Active"))
	fmt.Println(todo.ColorIndigo + "  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" + todo.ColorReset)
	fmt.Println()
}

func printProgress(list *todo.TodoList) {
	total, done := list.GetStats()
	fmt.Printf("\n   %s  \n", todo.StyledBar("Progress:"))
	fmt.Println(todo.StatsBar(total, done))
	fmt.Print("   " + todo.Indigo("● "))
	for i := 0; i < 20; i++ {
		if total > 0 && i < (done*20/total) {
			fmt.Print(todo.Indigo("■"))
		} else {
			fmt.Print(todo.Dim("□"))
		}
	}
	fmt.Print("\n\n")
}
