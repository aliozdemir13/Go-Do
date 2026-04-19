// Package main is the command center of the app
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"Go-Do/internal/todo"
)

func printHeader() {
	fmt.Print("\033[H\033[2J") // Clear screen
	fmt.Print(todo.MegaLogo())
	fmt.Printf("\n   %s %s %s\n",
		todo.Indigo("⚡ TASKS"),
		todo.Dim("v1.0.0"),
		todo.Indigo("●")+" "+todo.Dim("Local Storage Active"))
	fmt.Println(todo.ColorIndigo + "  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" + todo.ColorReset)
	fmt.Println()
}

func printProgress(myList *todo.TodoList) {
	// Get the stats
	total, done := myList.GetStats()

	// Print the Stats Bar
	fmt.Printf("\n   %s  \n", todo.StyledBar("Progress:"))
	fmt.Println(todo.StatsBar(total, done))

	// A small visual progress bar
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

func printMenu() {
	fmt.Println(todo.Indigo("\n┌ Management"))
	fmt.Printf("│ %s Add Task    %s Complete    %s Delete\n",
		todo.Dim("[1]"), todo.Dim("[2]"), todo.Dim("[3]"))

	fmt.Println(todo.Indigo("\n┌ View"))
	fmt.Printf("│ %s Open Tasks  %s Completed Tasks   %s Save\n",
		todo.Dim("[4]"), todo.Dim("[5]"), todo.Dim("[6]"))

	fmt.Print("\n" + todo.Indigo("selection ❯ "))
}

func main() {
	myList := todo.TodoList{}
	filename := "tasks.json"

	printHeader()

	// Try loading existing data
	err := myList.LoadFromFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		printProgress(&myList)
		fmt.Println(todo.StyledBar("OPEN TASKS "))
		myList.Display(false) // incomplete tasks display
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			fmt.Print(todo.Indigo("\n Task title: "))
			scanner.Scan()
			title := scanner.Text()
			myList.Add(title)
			fmt.Println("Task added!")
		case "2":
			fmt.Print(todo.Indigo("\n Task ID to complete: "))
			scanner.Scan()
			var id int
			// Convert the string input to an integer manually
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error:", err)
			}
			errComplete := myList.Complete(todo.TaskId(id))
			if errComplete != nil {
				fmt.Println("Error:", errComplete)
			} else {
				myList.Display(false)
			}
		case "3":
			fmt.Print(todo.Indigo("\n Task ID to delete: "))
			scanner.Scan()
			var id int
			// Convert the string input to an integer
			_, err := fmt.Sscanf(scanner.Text(), "%d", &id)
			if err != nil {
				fmt.Println("Error:", err)
			}
			errDelete := myList.Delete(todo.TaskId(id))
			if errDelete != nil {
				fmt.Println("Error:", errDelete)
			} else {
				myList.Display(false)
			}
		case "4":
			fmt.Println(todo.StyledBar("OPEN TASKS "))
			printProgress(&myList)
			myList.Display(false) // incomplete tasks display
		case "5":
			fmt.Println(todo.StyledBar(" \nCOMPLETED TASKS "))
			printProgress(&myList)
			myList.Display(true) // complete tasks display
		case "6":
			err = myList.SaveToFile(filename)
			if err != nil {
				fmt.Println(todo.Red("\n Saved! \n"))
			}
			fmt.Println(todo.Indigo("\n Saved! \n"))
			fmt.Println(todo.StyledBar("OPEN TASKS "))
			printProgress(&myList)
			myList.Display(false) // incomplete tasks display
		case "exit":
			fmt.Println(todo.Indigo("\n Goodbye!"))
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
