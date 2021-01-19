package model

// Task 任务对象
type Task struct {
	Func     func(*Param)
	ParamObj *Param
}

// NewTask 新建任务对象
func NewTask(f func(*Param), paramObj *Param) *Task {
	return &Task{
		Func:     f,
		ParamObj: paramObj,
	}
}
