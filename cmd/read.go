/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a specific task",
	Long: `Read a specific task`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if _, err := readFile(Path, "showtasks"); err!=nil {
				fmt.Println("Not able to read tasks", err)
			}
		} else {
			content, err := readFile(Path, "read")
			if err!=nil {
				fmt.Println("Not able to read tasks", err)
			}
			i,err := strconv.Atoi(args[0])
			if err!=nil {
				fmt.Println("Please give a number")
			}
			if len(content)> i {
				fmt.Println(content[i])
			} else{
				fmt.Println("Please specify a number within the task list")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
