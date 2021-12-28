/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"todo/todos"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos",
	Long:  ` Lists all the todos with Priority number, date added and tag`,
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	var items []todos.Item = todos.GetItems(rootCmd.PersistentFlags().Lookup("dataFile").Value.String())
	fmt.Println("-")
	for _, item := range items {
		fmt.Println("[ ]", item.Text)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
