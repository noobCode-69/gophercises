package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide at least one task ID")
			return
		}

		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse argument:", arg)
				return
			}
			ids = append(ids, id)
		}

		tasks, err := db.AllTask()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		exists := make(map[int]bool)
		for _, t := range tasks {
			exists[t.Key] = true
		}

		for _, id := range ids {
			if !exists[id] {
				fmt.Printf("Task %d does not exist\n", id)
				continue
			}

			err := db.DeleteTask(id)
			if err != nil {
				fmt.Printf("Failed to complete task %d: %v\n", id, err)
				continue
			}

			fmt.Printf("✓ Task %d completed\n", id)
		}
	},
}


func init() {
	RootCmd.AddCommand(doCmd)
}