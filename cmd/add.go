package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/pgzisis/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		_, err := db.CreateTask(task)
		if err != nil {
			log.Fatalln("Something went wrong:", err.Error())
		}

		fmt.Printf("Task \"%s\" was added to your list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
