package cmd

import (
	"fmt"
	"os"
	"strings"
)

func (t taskList) String() string {
		return fmt.Sprintf("%d %s", t.id, t.tasks)
}

func readFile(path string, caller string)([]taskList, error) {
	file,err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error Reading tasks", err)
		}
		if string(file) == "" {
			if caller=="showtasks" { 
				fmt.Println("No tasks to in list")
			}
			return []taskList{}, nil
		}
		TaskList := []taskList{}
		for i, task := range strings.Split(string(file), "\n"){
			if task==""{
				continue
			}
			TaskList = append(TaskList, taskList{
				id:i,
				tasks:task,
			})
			if caller=="showtasks" {
				fmt.Println(TaskList[i])
			}
		}
		return TaskList, nil
}
