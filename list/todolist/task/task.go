package task

type Task struct {
	id   int
	name string
	done bool
}

func newTask(id int, name string) Task {
	return Task{
		id:   id,
		name: name,
		done: false,
	}
}

func (task *Task) do() {
	task.done = true
}
