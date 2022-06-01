/*
Copyright © 2022 Parsa <carnoval@protonmail.com>

*/
package cmd

import (
	//"encoding/json"
	"fmt"
	//"os"

	// "github.com/Carnoval15/Todo-Boi/cmd/database"

	"github.com/Carnoval15/Todo-Boi/cmd/database"
	"github.com/spf13/cobra"
)

// newTaskCmd represents the newTask command that adds a new task to the list
var newTaskCmd = &cobra.Command{
	Use:     "newTask",
	Aliases: []string{"rev"},
	Short:   "A brief description of your command",
	Long:    `A longer description`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Task:", args[0])

		database.Add(args[0])
	},
}

var taskImportance int64
var task string

func init() {
	rootCmd.AddCommand(newTaskCmd)

	newTaskCmd.Flags().Int64VarP(&taskImportance, "importance", "i", 0, "Importance of the new task from 1 to 10")
}
