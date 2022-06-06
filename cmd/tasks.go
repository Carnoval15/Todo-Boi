/*
Copyright © 2022 Parsa <carnoval@protonmail.com>

*/
package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
	"github.com/Carnoval15/Todo-Boi/cmd/database"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Prints your Tasks on the screen",

	Run: func(cmd *cobra.Command, args []string) {
		database.Read()
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
