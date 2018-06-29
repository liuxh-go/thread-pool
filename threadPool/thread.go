package threadPool

import (
	"time"
)

/*
 * 线程对象
 *
 */

// 线程对象
type thread struct {
	// 线程池
	threadPool chan chan task

	// 线程通道
	taskChannel chan task

	// 退出通道
	quit chan bool
}

// 开始工作
func (this *thread) start() {
	go func() {
		for {
			// 将线程通道注册到线程池中
			this.threadPool <- this.taskChannel

			// 从线程通道中读取线程对象并执行任务
			select {
			case taskObj := <-this.taskChannel:
				taskObj(time.Now())
			case <-this.quit:
				return
			}
		}
	}()
}

// 停止工作
func (this *thread) stop() {
	this.quit <- true
}

// 新建线程对象
func newThread(_threadPool chan chan task) *thread {
	return &thread{
		threadPool:  _threadPool,
		taskChannel: make(chan task),
		quit:        make(chan bool),
	}
}
