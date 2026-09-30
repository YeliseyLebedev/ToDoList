package task

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type ManagerTask struct {
	id       int
	listTask []Task
}

func newManagerTask() ManagerTask {
	return ManagerTask{
		listTask: []Task{},
		id:       0,
	}
}

func (managerTask *ManagerTask) list() {
	for _, task := range managerTask.listTask {

		if task.done {
			fmt.Print(task.name, " - Выполнено!", " [", task.id, "]", "\n")
		} else {
			fmt.Print(task.name, " - Не выполнено :(", " [", task.id, "]", "\n")
		}
	}
}

func (managerTask *ManagerTask) add() {
	reader := bufio.NewReader(os.Stdin)
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	nTask := newTask(managerTask.id, task)
	managerTask.listTask = append(managerTask.listTask, nTask)
	managerTask.id++
	fmt.Println("Успех!")
}

func (managerTask *ManagerTask) done() {
	var id int
	fmt.Scan(&id)

	if id > managerTask.id {
		fmt.Println("Такой задачи не существует!")
		return
	}

	managerTask.listTask[id].do()
}

func (managerTask *ManagerTask) delete() {
	var id int
	fmt.Scan(&id)

	if id > managerTask.id {
		fmt.Println("Такой задачи не существует!")
		return
	}

	managerTask.listTask = slices.Delete(managerTask.listTask, id, id+1)
	fmt.Println("Успех")
}
