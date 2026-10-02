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

func (managerTask *ManagerTask) GetListTask() []Task {
	return managerTask.listTask
}

func (managerTask *ManagerTask) GetId() int {
	return managerTask.id
}

func (managerTask *ManagerTask) SetListTask(newList []Task) {
	managerTask.listTask = newList
}

func (managerTask *ManagerTask) SetId(newId int) {
	managerTask.id = newId
}

func (managerTask *ManagerTask) List() {
	if len(managerTask.listTask) == 0 {
		fmt.Println("Задач нет!")
		return
	}

	for _, task := range managerTask.listTask {

		if task.done {
			fmt.Print(task.name, " - Выполнено!", " [", task.id, "]", "\n")
		} else {
			fmt.Print(task.name, " - Не выполнено :(", " [", task.id, "]", "\n")
		}
	}
}

func (managerTask *ManagerTask) Add() {
	reader := bufio.NewReader(os.Stdin)
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	if strings.ContainsAny(task, "+-*/=") || task == "" {
		fmt.Println("Неверный ввод!")
		return
	}

	nTask := newTask(len(managerTask.listTask)+1, task)
	managerTask.listTask = append(managerTask.listTask, nTask)
	managerTask.id++
	fmt.Println("Успех!")
}

func (managerTask *ManagerTask) Delete() {
	var id int
	fmt.Scan(&id)

	if id < 1 || id > len(managerTask.listTask) {
		fmt.Println("Такой задачи не существует!")
		return
	}

	managerTask.listTask = slices.Delete(managerTask.listTask, id-1, id)
	managerTask.id--

	managerTask.reId()
	fmt.Println("Успех")
}

func (managerTask *ManagerTask) Done() {
	var id int
	fmt.Scan(&id)

	if id < 1 || id > len(managerTask.listTask) {
		fmt.Println("Такой задачи не существует!")
		return
	}

	managerTask.listTask[id-1].do()
	fmt.Println("Успех")
}

func (managerTask *ManagerTask) reId() {
	for i := range managerTask.listTask {
		managerTask.listTask[i].id = i + 1
	}
}
