package task

import (
	"fmt"
)

func Show() {
	fmt.Print("help - сводка команд\n\n")

	manager := newManagerTask()

	active := true

	for active {
		fmt.Print("todo ")
		task := ""
		fmt.Scan(&task)

		switch task {
		case "add":
			manager.add()

		case "list":
			manager.list()

		case "done":
			manager.done()

		case "delete":
			manager.delete()

		case "help":
			help()

		case "exit":
			active = false

		default:
			fmt.Println("Неправильный ввод!")
		}
	}
}

func help() {
	fmt.Println("add 'название задачи' - добавить задачу")
	fmt.Println("list - вывести все задачи")
	fmt.Println("done 'номер задачи' - завершить задачу")
	fmt.Println("delete 'номер задачи' - удалить задачу")
	fmt.Print("exit - выйти из программы\n\n")
}
