package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	// "time"
)

var priorityMap = map[string]int{
	"high":   3,
	"medium": 2,
	"low":    1,
}

// func sortPriorityTasks(data []taskListWrite) (t []taskListWrite, err error) {
// 	sort.Slice(data, func(i, j int) bool {
// 		return priorityMap[data[i].Priority] > priorityMap[data[j].Priority]
// 	})
// 	return data, nil
// }

// func sortStatusTasks(data []taskListWrite) (t []taskListWrite, err error) {
// 	statusdata := []taskListWrite{}
// 	for _, task := range data {
// 		if task.Status == statusSort {
// 			statusdata = append(statusdata, task)
// 		}
// 	}
// 	return statusdata, err
// }

// func sortTimeTasks(data []taskListWrite) (t []taskListWrite, err error) {
// 	sort.Slice(data, func(i, j int) bool {
// 		return data[i].Time.Sub(data[j].Time) < 0
// 	})
// 	return data, nil
// }

// func sortTasks(prioritySort string, statusSort string, timeSort string) (t []taskListWrite, err error) {
// 	file, err := os.ReadFile(Path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	data := []taskListWrite{}
// 	err = json.Unmarshal(file, &data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	finalSort := []taskListWrite{}
// 	// fmt.Println(finalSort)
// 	if prioritySort != "" {
// 		finalSort, err = sortPriorityTasks(data)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	if statusSort != "" {
// 		if len(finalSort) != 0 {
// 			statusSortData, err := sortStatusTasks(finalSort)
// 			if err != nil {
// 				panic(err)
// 			}
// 			finalSort = statusSortData
// 		} else {
// 			finalSort, err = sortStatusTasks(data)
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 	}
// 	if timeSort != "" {
// 		if len(finalSort) != 0 {
// 			timeSortData, err := sortTimeTasks(finalSort)
// 			if err != nil {
// 				panic(err)
// 			}
// 			finalSort = append(finalSort, timeSortData...)
// 		} else {
// 			finalSort, err = sortTimeTasks(data)
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 	}

// 	if len(finalSort) != 0 {
// 		printTasks(finalSort)
// 	} else {
// 		fmt.Println("Please filter by available task filters")
// 	}

//		return []taskListWrite{}, nil
//	}

func printTasks(t []taskListWrite) {
	for _, t := range t {
		fmt.Println(t)
	}
}

func sortTasks(prioritySort string, statusSort string, timeSort string) (t []taskListWrite, err error) {
	file, err := os.ReadFile(Path)
	if err != nil {
		panic(err)
	}
	data := []taskListWrite{}
	if err = json.Unmarshal(file, &data); err != nil {
		panic(err)
	}
	if statusSort != "" {
		finalData := []taskListWrite{}
		for _, task := range data {
			if task.Status == statusSort {
				finalData = append(finalData, task)
			}
		}
		data = finalData
	}

	if prioritySort != "" {
		finalData := []taskListWrite{}
		for _, task := range data {
			if task.Priority == prioritySort {
				finalData = append(finalData, task)
			}
		}
		data = finalData
	}

	if timeSort != "" {
		finalData := []taskListWrite{}
		for _, task := range data {
			var date = task.Time.Format("2006-01-02")
			if date == timeSort {
				finalData = append(finalData, task)
			}
		}
		if len(finalData) != 0 {
			data = finalData
		}
	}

	sort.Slice(data, func(i, j int) bool {
		if prioritySort != "" && priorityMap[data[i].Priority] != priorityMap[data[j].Priority] {
			return priorityMap[data[i].Priority] > priorityMap[data[j].Priority]
		}
		if timeSort == "new" && !data[i].Time.Equal(data[j].Time) {
			return data[i].Time.Sub(data[j].Time) > 0
		}
		if timeSort == "old" && !data[i].Time.Equal(data[j].Time) {
			return data[i].Time.Sub(data[j].Time) < 0
		}
		return false
	})

	if len(data) != 0 {
		printTasks(data)
	} else {
		fmt.Println("No tasks match the provided filters")
	}

	return data, nil
}
