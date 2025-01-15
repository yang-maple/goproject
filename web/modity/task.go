package modity

type Task struct {
	ID     int
	Name   string
	Status string
}

var taskslice = make([]*Task, 0, 1024)

// 查询功能
func Findtask() []*Task {
	return taskslice
}

// 新增功能
func Addtask(task Task) {

	taskslice = append(taskslice, &Task{task.ID, task.Name, task.Status})

}

// 删除功能
func Deltask(tasks Task) {
	for i, v := range taskslice {
		if tasks.ID == v.ID || tasks.Name == v.Name {
			taskslice = append(taskslice[:i], taskslice[i+1:]...)
		}
	}
}

// 更新功能
func Updatatask(tasks Task) {
	for _, v := range taskslice {
		if v.ID == tasks.ID || tasks.Name == v.Name {
			v.ID = tasks.ID
			v.Name = tasks.Name
			v.Status = tasks.Status
		}
	}
}
