package TaskClass

import ("fmt"
		"time"
)

type Task struct{
	name string
	description string
	created_at time.Time
}

func New(name string, description string) Task{
	current_time := time.Now()
	newTask := Task {name, description, current_time}
	return newTask
}

func (task Task) DisplayTask(){
	fmt.Printf("Task: %s\nDescription: %s\nCreated at: %s\n", task.name, task.description, task.created_at.Format("2006-01-02 15:04"))
}