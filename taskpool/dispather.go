package taskpool

import (
	"sync"

	"github.com/liuxh-go/thread-pool/model"
)

// Dispather 任务分配器对象
type Dispather struct {
	threadPool     chan chan *model.Task
	maxThreadCount int32
	taskQueue      chan *model.Task
	threadList     []*thread
	wg             *sync.WaitGroup
}

// Run 启动任务分配器
func (d *Dispather) Run() {
	d.threadList = make([]*thread, 0, d.maxThreadCount)
	for i := 0; i < int(d.maxThreadCount); i++ {
		threadObj := newThread(d.threadPool, d.wg)
		threadObj.start()
		d.threadList = append(d.threadList, threadObj)
	}

	go d.dispather()
}

// WaitStop 等待完成
func (d *Dispather) WaitStop() {
	d.wg.Wait()

	for _, thread := range d.threadList {
		thread.stop()
	}
}

// AddTask 添加任务
func (d *Dispather) AddTask(task *model.Task) {
	d.taskQueue <- task
	d.wg.Add(1)
}

func (d *Dispather) dispather() {
	for {
		select {
		// 从任务队列读取一个任务对象
		case taskObj := <-d.taskQueue:
			go func(taskObj *model.Task) {
				// 从线程池取出一条空闲线程通道,将任务对象发送到空闲的线程通道中
				threadChannel := <-d.threadPool
				threadChannel <- taskObj
			}(taskObj)

		default:
		}
	}
}

// NewDispather 新建任务分配器
func NewDispather(maxWorkerCount int32) *Dispather {
	return &Dispather{
		threadPool:     make(chan chan *model.Task, maxWorkerCount),
		maxThreadCount: maxWorkerCount,
		taskQueue:      make(chan *model.Task),
		wg:             new(sync.WaitGroup),
	}
}
