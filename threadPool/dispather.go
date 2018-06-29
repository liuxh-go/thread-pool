package threadPool

/*
 * 任务分配器
 *
 */

type Dispather struct {
	// 线程池
	threadPool chan chan task

	// 线程池最大数量
	maxThreadCount int32

	// 任务通道
	taskQueue chan task
}

// Run 启动任务分配器
func (this *Dispather) Run() {
	for i := 0; i < int(this.maxThreadCount); i++ {
		threadObj := newThread(this.threadPool)
		threadObj.start()
	}

	go this.dispather()
}

// AddTask 添加任务
// 参数
// _task:任务对象
func (this *Dispather) AddTask(_task task) {
	this.taskQueue <- _task
}

// 分配任务
func (this *Dispather) dispather() {
	for {
		select {
		// 从任务队列读取一个任务对象
		case taskObj := <-this.taskQueue:
			go func(_taskObj task) {
				// 从线程池取出一条空闲线程通道,将任务对象发送到空闲的线程通道中
				threadChannel := <-this.threadPool
				threadChannel <- _taskObj
			}(taskObj)

		default:
		}
	}
}

// NewDispather 新建任务分配器
func NewDispather(_maxWorkerCount int32) *Dispather {
	return &Dispather{
		threadPool:     make(chan chan task, _maxWorkerCount),
		maxThreadCount: _maxWorkerCount,
		taskQueue:      make(chan task),
	}
}
