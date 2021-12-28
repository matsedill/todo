/*
Copyright © 2021 Alex Gunnarsson

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A brief description of your application",
	Long: `Tackle your life with a single command. Efficiency, simplicity 
	are the value of todo.  The tool can be used to for tracking your tasks, 
	inside and outside of a professional environment.
	Examples
	
	todo add -p2 "Finish this descriptio"

	todo list
	
	todo edit 1 -p1 "Finish this description"
	
	todo x # checks off the first todo
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().String(
		"dataFile",
		os.ExpandEnv("$HOME")+string(os.PathSeparator)+".todos.json",
		"Path to Todos")
	viper.BindPFlag("dataFile", rootCmd.PersistentFlags().Lookup("dataFile"))
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
