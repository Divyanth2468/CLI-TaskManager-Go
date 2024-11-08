package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func alert() {
	file, err := os.ReadFile(Path)
	if err != nil {
		panic(err)
	}
	data := []taskListWrite{}
	if len(file) != 0 {
		if err = json.Unmarshal(file, &data); err != nil {
			panic(err)
		}
	}
	updatedTasks := []taskListWrite{}
	for _, task := range data {
		if task.Time.UTC().Before(time.Now().UTC()) || task.Status == "completed" {
			continue
		}

		if time.Until(task.Time) < 24*time.Hour {
			var message string = task.Description + " " + task.Time.Format("2006-01-02") //+ "File to task file " + openTaskFile(task)
			err := beeep.Notify(task.Tasks, message, "assets/information.png")
			if err != nil {
				fmt.Println("Notification error for task:", task.Tasks, err)
			}
		}
		updatedTasks = append(updatedTasks, task)
	}

	for i := range updatedTasks {
		updatedTasks[i].Id = i
	}

	newTaskList, err := json.Marshal(updatedTasks)
	if err != nil {
		fmt.Println("Error marshalling tasks to JSON:", err)
	}
	err = os.WriteFile(Path, newTaskList, 0664)
	if err != nil {
		fmt.Println("Error writing updated tasks to file:", err)
	}
}

// func openTaskFile(t taskListWrite) string {
// 	taskFileName := fmt.Sprintf("%s/%s_task.txt", PathText, t.Tasks)
// 	fmt.Println(taskFileName)
// 	taskFileContent := fmt.Sprintf("Task: %s\nPriority: %s\nStatus: %s\nDescription: %s\nDue Date: %s\n", t.Tasks, t.Priority, t.Status, t.Description, t.Time.Format("2006-01-02"))

// 	if err := os.WriteFile(taskFileName, []byte(taskFileContent), 0664); err != nil {
// 		fmt.Println("Error writing to file:", err)
// 	}

// 	return taskFileName

// }
