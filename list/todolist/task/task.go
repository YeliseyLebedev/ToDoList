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

func (task *Task) GetId() int {
	return task.id
}

func (task *Task) GetName() string {
	return task.name
}

func (task *Task) GetDone() bool {
	return task.done
}

func (task *Task) SetId(newId int) {
	task.id = newId
}

func (task *Task) SetName(newName string) {
	task.name = newName
}

func (task *Task) SetDone(flag bool) {
	task.done = flag
}
