/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Path string = "/Users/uppuluridivyanthsatya/Desktop/Go/Tasks.json"

var priority string
var status string

type taskListWrite struct {
	Id       int    `json:"Id"`
	Tasks    string `json:"string"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tasks to your list",
	Long:  `Add tasks to your list`,
	Run: func(cmd *cobra.Command, args []string) {
		// file, err :=os.OpenFile("/Users/uppuluridivyanthsatya/Desktop/Go/Tasks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664);
		// if err!=nil{
		// 	fmt.Println("Error Loading File", err)
		// }

		// defer file.Close()
		// for _, arg := range args {
		// 	_, err := file.WriteString(arg + "\n")
		// 	if err !=nil{
		// 		fmt.Println("Error writing to file", err)
		// 	}
		// }
		if len(args) == 0 {
			fmt.Println("Please enter the task with priority and status , the defaults are low priority and running status")
		} else {
			data, err := os.ReadFile(Path)
			if err != nil {
				fmt.Println("Error reading json file", err)
			}
			if len(data) == 0 {
				data = []byte("[]")
			}
			data_old := []taskListWrite{}
			err = json.Unmarshal(data, &data_old)
			if err != nil {
				fmt.Println("Error unmarshalling data", err)
			}
			var index int = len(data_old)
			data_new := taskListWrite{}
			for _, arg := range args {
				data_new.Id = index
				data_new.Tasks = arg
				data_new.Priority = priority
				data_new.Status = status
				data_old = append(data_old, taskListWrite(data_new))
				index += 1
			}
			jsonData, err := json.Marshal(data_old)
			if err != nil {
				fmt.Println("Error converting into json", err)
			}
			err = os.WriteFile(Path, jsonData, 0664)
			if err != nil {
				fmt.Println("Error writing to file", err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&priority, "priority", "p", "low", "Set the priority of the task")
	addCmd.Flags().StringVarP(&status, "status", "s", "running", "Set the status of the task")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
