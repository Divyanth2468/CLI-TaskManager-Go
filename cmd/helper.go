package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

func (t taskListWrite) String() string {
		return fmt.Sprintf("Index: %d Task: %s Priority: %s Status: %s", t.Id, t.Tasks, t.Priority, t.Status)
}

func readFile(path string, caller string)([]taskListWrite, error) {

	// For Json file
	file,err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error Reading Tasks", err)
		}
	if len(file) == 0 {
		file = []byte("[]")
	}
		TaskList := []taskListWrite{}
		err=json.Unmarshal(file, &TaskList)
		if err!=nil {
			panic(err)
		}
		if len(TaskList) == 0 {
				fmt.Println("No Tasks to in list")
			return []taskListWrite{}, nil
		}
		if caller=="showtasks"{
			for _,task := range TaskList{
				fmt.Println(task)
			}
		}

		// For text file
		// for i, task := range strings.Split(string(file), "\n"){
		// 	if task==""{
		// 		continue
		// 	}
		// 	TaskList = append(TaskList, taskListRead{
		// 		Id:i,
		// 		Tasks:task,
		// 	})
		// 	if caller=="showTasks" {
		// 		fmt.Println(TaskList[i])
		// 	}
		// }
		return TaskList, nil
}
