package main

import ("fmt"
		"bufio"
		"os"
		"strings"
		"slices"
		"strconv"
		"ToDoList/TaskClass")

func main(){
	fmt.Print("CLI App in which, you can create Tasks to do, edit them, delete and display list of Tasks\n")
	var tasksList []TaskClass.Task
	indexTitleMap:= make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	var action string
	
	for {
		fmt.Print(`
	Choose an action to do:
	q - exit programme
	display - display a list of tasks
	add - add a task to list
	delete - delete a task by name or index
> `)

		action, _ = reader.ReadString('\n')
		action = strings.TrimSpace(action)

		switch action {
			case "q":
				return
			case "add":
				fmt.Print("\nEnter title of the Task: \n")
				var taskTitle string
				taskTitle, _ = reader.ReadString('\n')
				taskTitle = strings.TrimSpace(taskTitle)
				fmt.Print("Enter description of the Task: \n")
				var taskDescription string
				taskDescription, _ = reader.ReadString('\n')
				taskDescription = strings.TrimSpace(taskDescription)
				var newTask TaskClass.Task = TaskClass.New(taskTitle, taskDescription)
				tasksList = append(tasksList, newTask)
				indexTitleMap[taskTitle] = len(tasksList) - 1
			case "display":
				for index, task := range tasksList{
					fmt.Printf("\nTask index: %d\n", index)
					task.DisplayTask()
				}
			case "delete":
				fmt.Print("\nTask by index or name?\n")
				var option string
				option, _ = reader.ReadString('\n')
				option = strings.TrimSpace(option)
				switch option {
					case "name":
						fmt.Print("\nEnter name of task to remove: ")
						var name string
						name, _ = reader.ReadString('\n')
						name = strings.TrimSpace(name)
						index, no_err := indexTitleMap[name]
						if no_err {
							tasksList = slices.Delete(tasksList, index, index + 1)
							fmt.Printf("Deleted Task '%s' successfully\n", name)
						} else {
							fmt.Printf("An error occurred during deleting task '%s'. Is name correct?\n", name)
						}
					case "index":
						fmt.Print("\nEnter index of task to remove: ")
						var index string
						index, _ = reader.ReadString('\n')
						index = strings.TrimSpace(index)
						index_value, err := strconv.Atoi(index)
						if index_value < len(tasksList) && err == nil {
							tasksList = slices.Delete(tasksList, index_value, index_value + 1)
							fmt.Printf("Deleted Task '%d' successfully\n", index_value)
						} else {
							fmt.Printf("An error occurred during deleting task '%d'. Is index correct?\n", index_value)
						}
					default:
						fmt.Print("Unknown option :c")
				}
			default:
				fmt.Print("Unknown option :c")
		}
	}
}