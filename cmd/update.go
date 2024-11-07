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

var newTaskValue string
var newPriority string
var newStatus string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update specific tasks",
	Long: `Update specific tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("missing index and updated task value")
		}
			data, err := os.ReadFile(Path)
			if err != nil {
				panic(err)
			}
			ind, err := strconv.Atoi(args[0])
			if err!=nil {
				panic(err)
			}
			content := []taskListWrite{}
			if err=json.Unmarshal(data, &content); err!=nil {
				panic(err)
			}
			if newTaskValue != "" {
				content[ind].Tasks = newTaskValue
			} 
			if newPriority !="" {
				content[ind].Priority = newPriority
			} 
			if newStatus != "" {
				content[ind].Status = newStatus
			}
			data_new, err:=json.Marshal(content)
			if err!=nil{
				panic(err)
			}
			os.WriteFile(Path, data_new, 0666)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&newTaskValue, "task-value", "t", "", "task value to update")
	updateCmd.Flags().StringVarP(&newPriority, "priority-value", "p", "", "priority value to update")
	updateCmd.Flags().StringVarP(&newStatus, "status-value", "s", "", "status value to update")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
