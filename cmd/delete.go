/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete specific task by id or the entire task list",
	Long: `Delete specific task by id or the entire task list`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_, err := os.OpenFile(Path, os.O_WRONLY|os.O_TRUNC, 0666)
			if err!=nil{
				fmt.Println("Cannot clear file", err)
			}
			fmt.Println("Task List is deleted")
		} else {
			ind, err:=strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Please enter a valid index", err)
			}
			content, err := readFile(Path, "delete")
			if err != nil {
				fmt.Println("Error reading tasks", err)
			}
			if ind <=len(content)-1 {
				content = append(content[:ind], content[ind+1:]...)
				for j:=ind; j<len(content);j++{
					content[j].Id--
				}
				data, err := json.Marshal(content)
				if err != nil {
					panic(err)
				}
				os.WriteFile(Path, data, 0664)
			} else {
				fmt.Println("Please specify a task within the task list")
			}
			

			// file, err := os.OpenFile("../Tasks.txt", os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0666)
			// if err!=nil{
			// 	fmt.Println("Cannot clear file", err)
			// }

			// if ind > len(content) - 1 {
			// 	fmt.Println("Task does not exist")
			// } else{
			// 	for _, task := range content{
			// 		if(task.id!=ind){
			// 			file.WriteString(task.tasks + "\n")
			// 		}
			// 	}
			// }
			// defer file.Close()
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
