package main

import "fmt"

type Task struct {
	Description string
	Done        bool
}

func main() {
	// Part 1: Create a pointer to a Task using new
	task1 := new(Task)
	task1.Description = "Read chapter"
	task1.Done = false

	// Part 2: Create a pointer to a Task using &
	task2 := &Task{
		Description: "Complete practice problems",
		Done:        false,
	}

	// Part 3: Store tasks in a slice
	tasks := []*Task{task1, task2}

	// TODO: Append one more dynamically allocated task
	task3 := new(Task{
		Description: "Review pointer examples",
		Done:        true,
	})
	tasks = append(tasks, task3)

	// Part 4: Store task owner in a map
	owner := make(map[*Task]string)

	owner[task1] = "Alice"
	owner[task2] = "Bob"

	// TODO: Add owner for third task
	owner[task3] = "Carol"

	// Part 5: Print all tasks and owners
	for _, task := range tasks {
		fmt.Printf("%s owns task: %s, done: %v\n",
			owner[task],
			task.Description,
			task.Done)
	}
}
