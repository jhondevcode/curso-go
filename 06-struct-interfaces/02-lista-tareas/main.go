package main

import "fmt"

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) append(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) remove(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

type Task struct {
	name        string
	description string
	completed   bool
}

func (t *Task) Complete() {
	t.completed = true
}

func (t *Task) ToString() string {
	return fmt.Sprintf("Task { name: %s, description: %s, completed: %t }", t.name, t.description, t.completed)
}

func main() {
	var task_list TaskList = TaskList{}

	var task1 Task = Task{
		name:        "Paint",
		description: "Paint the walls of the house",
		completed:   false,
	}

	task1.Complete()
	task_list.append(&task1)

	var task2 Task = Task{
		name:        "Clean the garden",
		description: "Pick up and throw away dry leaves",
		completed:   true,
	}

	task_list.append(&task2)

	var task3 Task = Task{
		name:        "Feed",
		description: "Feed the animals",
		completed:   false,
	}
	task_list.append(&task3)

	fmt.Println(task_list)

	for i, task := range task_list.tasks {
		fmt.Println(i, task.ToString())
	}

}
