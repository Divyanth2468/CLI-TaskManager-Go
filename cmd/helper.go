package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func (t taskListWrite) String() string {
	return fmt.Sprintf("Inex: %d Task: %s Description: %s Priority: %s Status: %s Deadline: %s", t.Id, t.Tasks, t.Description, t.Priority, t.Status, t.Time.UTC().Format(time.RFC1123))
}

func readFile(path string, caller string) ([]taskListWrite, error) {

	// For Json file
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error Reading Tasks", err)
	}
	if len(file) == 0 {
		file = []byte("[]")
	}
	TaskList := []taskListWrite{}
	err = json.Unmarshal(file, &TaskList)
	if err != nil {
		panic(err)
	}
	if len(TaskList) == 0 {
		fmt.Println("No Tasks to in list")
		return []taskListWrite{}, nil
	}
	if caller == "showtasks" {
		for _, task := range TaskList {
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
