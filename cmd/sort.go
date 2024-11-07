package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

var priorityMap = map[string]int{
	"high":   3,
	"medium": 2,
	"low":    1,
}

func sortPriorityTasks(data []taskListWrite) (t []taskListWrite, err error) {
	sort.Slice(data, func(i, j int) bool {
		return priorityMap[data[i].Priority] > priorityMap[data[j].Priority]
	})
	return data, err
}

func printTasks(t []taskListWrite) {
	for _, t := range t {
		fmt.Printf("Index: %d Task: %s Priority: %s Status: %s \n", t.Id, t.Tasks, t.Priority, t.Status)
	}
}

func sortTasks(prioritySort string, statusSort string) (t []taskListWrite, err error) {
	// if sortingtype != "priority" && sortingtype != "status"{
	// 	fmt.Println("Please slect valid sort option to sort by either priority or status")
	// }
	file, err := os.ReadFile(Path)
	if err != nil {
		panic(err)
	}
	data := []taskListWrite{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}
	statusdata := []taskListWrite{}
	if statusSort != "" {
		for _, task := range data {
			if task.Status == statusSort {
				statusdata = append(statusdata, task)
			}
		}
	}
	if prioritySort != "" {
		if len(statusdata) != 0 {
			data, err := sortPriorityTasks(statusdata)
			if err != nil {
				fmt.Println("Error sorting")
			}
			printTasks(data)
		} else {
			data, err := sortPriorityTasks(data)
			if err != nil {
				panic(err)
			}
			printTasks(data)
		}
	} else {
		printTasks(statusdata)
	}
	return []taskListWrite{}, nil
}
