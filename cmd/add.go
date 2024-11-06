/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tasks to your list",
	Long: `Add tasks to your list`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err :=os.OpenFile("/Users/uppuluridivyanthsatya/Desktop/Go/Tasks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664); 
		if err!=nil{
			fmt.Println("Error Loading File", err)
		}

		defer file.Close()
		for _, arg := range args {
			_, err := file.WriteString(arg + "\n")
			if err !=nil{
				fmt.Println("Error writing to file", err)
			}
		}
	},
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
