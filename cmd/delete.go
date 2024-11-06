/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_, err := os.OpenFile("/Users/uppuluridivyanthsatya/Desktop/Go/Tasks.txt", os.O_WRONLY|os.O_TRUNC, 0666)
			if err!=nil{
				fmt.Println("Cannot clear file", err)
			}
			fmt.Println("Task List is deleted")
		} else {
			ind, err:=strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Please enter a valid index", err)
			}
			content, err := readFile("/Users/uppuluridivyanthsatya/Desktop/Go/Tasks.txt", "delete")
			if err != nil {
				fmt.Println("Error reading tasks", err)
			}
			file, err := os.OpenFile("../Tasks.txt", os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0666)
			if err!=nil{
				fmt.Println("Cannot clear file", err)
			}

			if ind > len(content) - 1 {
				fmt.Println("Task does not exist")
			} else{
				for _, task := range content{
					if(task.id!=ind){
						file.WriteString(task.tasks + "\n")
					}
				}
			}
			defer file.Close()
	}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
