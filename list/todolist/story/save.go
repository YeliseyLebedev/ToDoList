package story

import (
	"fmt"
	"os"
	"strconv"
	"study/todolist/task"
)

func Save(manager task.ManagerTask) {
	file, err := os.OpenFile("todolist/story/savefile/save", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Ошибка сохранения!")
		return
	}
	defer file.Close()

	file.WriteString(strconv.Itoa(manager.GetId()) + "\n")

	saveListTask := manager.GetListTask()
	for _, task := range saveListTask {
		text := ""
		if task.GetDone() {
			text = strconv.Itoa(task.GetId()) + "|" + task.GetName() + "|1"
		} else {
			text = strconv.Itoa(task.GetId()) + "|" + task.GetName() + "|0"
		}
		file.WriteString(text + "\n")
	}
}
