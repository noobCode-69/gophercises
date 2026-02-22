package cmd

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all the task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks , err := db.AllTask()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0  {
			fmt.Println("You have no task! Enjoy")
			return;
		}
		fmt.Println("You have the following tasks:")
		for _, task := range tasks {
    		fmt.Printf("%d. %s\n", task.Key, task.Value)
		}
	},	
}

func init() {
	RootCmd.AddCommand(listCmd)
}