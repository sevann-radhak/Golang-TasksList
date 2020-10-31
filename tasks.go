package main

import "fmt"

type task struct {
	description string
	done        bool
	name        string
}

type tasksList struct {
	tasks []*task
}

func main() {
	t := &task{
		description: "Finish Golang course",
		done:        false,
		name:        "Golang Course",
	}

	fmt.Printf("%+v\n", t)

	t.updateDescription("Finish Golang course. Edited")
	t.updateName("Golang Course. Edited")
	t.finishTask()

	fmt.Printf("%+v\n", t)

	fmt.Println(" ************************************* ")

	t1 := &task{
		description: "Task 1 Description",
		done:        false,
		name:        "Task 1",
	}

	t2 := &task{
		description: "Task 2 Description",
		done:        false,
		name:        "Task 2",
	}

	t3 := &task{
		description: "Task 3 Description",
		done:        false,
		name:        "Task 3",
	}

	tlist := &tasksList{
		tasks: []*task{
			t1,
			t2,
			t3,
		},
	}

	fmt.Println("Total tasks: ", len(tlist.tasks))
	fmt.Println("Add 1 taks")
	tlist.addToList(t3)
	fmt.Println("Total tasks: ", len(tlist.tasks))
	fmt.Println("Delete 1 taks")
	tlist.deleteFromList(2)
	fmt.Println("Total tasks: ", len(tlist.tasks))

	// for
	for i := 0; i < len(tlist.tasks); i++ {
		fmt.Println("Tarea: ", i, " - ", tlist.tasks[i].name)
	}

	// range
	for index, task := range tlist.tasks {
		fmt.Println("Tarea: ", index, " - ", task.name)
	}

	fmt.Println(" ******** TASK LIST *****************")
	tlist.printTasksList()
	fmt.Println(" ******** COMPLETED *****************")
	tlist.tasks[1].finishTask()
	tlist.printTasksCompleted()

	fmt.Println(" ******** MAPS *****************")
	mapTasks := make(map[string]*tasksList)
	mapTasks["Sevann"] = tlist
	mapTasks["Radhak"] = tlist

	mapTasks["Sevann"].printTasksList()
}

func (t *tasksList) addToList(task *task) {
	t.tasks = append(t.tasks, task)
}

func (t *tasksList) deleteFromList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *task) finishTask() {
	t.done = true
}

func (p *tasksList) printTasksCompleted() {
	for index, task := range p.tasks {
		if task.done {
			fmt.Println(index, " - ", task.name, " - Descriptiton: ", task.description, " - Done: ", task.done)
		}
	}
}

func (p *tasksList) printTasksList() {
	for index, task := range p.tasks {
		fmt.Println(index, " - ", task.name, " - Descriptiton: ", task.description, " - Done: ", task.done)
	}
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}
