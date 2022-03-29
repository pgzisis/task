package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/pgzisis/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		ids := parseArgs(args)

		tasks, err := db.ReadAllTasks()
		if err != nil {
			log.Fatalln("Something went wrong:", err.Error())
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)

				continue
			}

			t := tasks[id-1]
			err := db.DeleteTask(t.Key)
			if err != nil {
				log.Fatalf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed\n", id)
			}
		}
	},
}

func parseArgs(args []string) []int {
	var ids []int

	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Failed to parse the argument:", arg)
		} else {
			ids = append(ids, id)
		}
	}

	return ids
}

func init() {
	RootCmd.AddCommand(doCmd)
}
