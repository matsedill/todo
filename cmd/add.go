/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"todo/todos"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an item to your todo list",
	Long: `Adding items to the todo list is a trivial tasks. However,
	without priority todo lists are nothing. That is why todo only has
	space for 1 p1 task.`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	newTodos := []todos.Item{}
	for _, x := range args {
		newTodos = append(newTodos, todos.Item{Text: x})
	}
	fmt.Println("Adding on to the pile...")
	todos.SaveItems(todos.GetLocation(), newTodos)
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
