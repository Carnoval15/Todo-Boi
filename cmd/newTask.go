/*
Copyright © 2022 Parsa <carnoval@protonmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newTaskCmd represents the newTask command that adds a new task to the list
var newTaskCmd = &cobra.Command{
	Use:   "newTask",
	Aliases: []string{"rev"},
	Short: "A brief description of your command",
	Long: `A longer description`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if taskImportance > 10 {
			// defer cmd.Execute()
			panic("Importance must be between 1 and 10 \n Exiting the program...")
		}
		fmt.Println("newTask called")
		fmt.Println("Task:", args[0])
		fmt.Println("Task Importance", taskImportance)

		//fmt.Println("taskImportance:", taskImportance)
	},
}

var taskImportance int64
var task string

func init() {
	rootCmd.AddCommand(newTaskCmd)

	newTaskCmd.Flags().Int64VarP(&taskImportance, "importance", "i", 0, "Importance of the new task from 1 to 10")
}
