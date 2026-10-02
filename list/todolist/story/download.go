package story

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"study/todolist/task"
)

const savePath = "todolist/story/savefile/save"

func Download() task.ManagerTask {
	file, err := os.Open(savePath)
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return task.ManagerTask{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	manager := task.ManagerTask{}

	if scanner.Scan() {
		id, _ := strconv.Atoi(scanner.Text())
		manager.SetId(id)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
		return task.ManagerTask{}
	}

	var list []task.Task
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) != 3 {
			continue
		}

		id, _ := strconv.Atoi(parts[0])
		name := parts[1]
		done := parts[2] == "1"

		t := task.Task{}
		t.SetId(id)
		t.SetName(name)
		t.SetDone(done)

		list = append(list, t)
	}

	manager.SetListTask(list)
	return manager
}
