package cmd

import (
	"fmt"
	"log"

	"github.com/pgzisis/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadAllTasks()
		if err != nil {
			log.Fatalln("Something went wrong:", err.Error())
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!")
		} else {
			fmt.Println("You have the following tasks:")
			for i, t := range tasks {
				fmt.Printf("%d. %s\n", i+1, t.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
