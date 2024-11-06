/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type taskList struct{
	id int
	tasks string
}

// showtasksCmd represents the showtasks command
var showtasksCmd = &cobra.Command{
	Use:   "showtasks",
	Short: "Shows list of tasks you have added",
	Long: `Shows list of tasks you have added`,
	Run: func(cmd *cobra.Command, args []string) {
	 if _,err := readFile("/Users/uppuluridivyanthsatya/Desktop/Go/Tasks.txt", "showtasks"); err!=nil {
		fmt.Println("Error showing tasks", err)
	 }

	},
}

func init() {
	rootCmd.AddCommand(showtasksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showtasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showtasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}