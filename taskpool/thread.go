package taskpool

import (
	"sync"

	"github.com/liuxh-go/thread-pool/model"
)

type thread struct {
	threadPool  chan chan *model.Task
	taskChannel chan *model.Task
	quit        chan bool
	wg          *sync.WaitGroup
}

func (t *thread) start() {
	go func() {
		for {
			// 将线程通道注册到线程池中
			t.threadPool <- t.taskChannel

			// 从线程通道中读取线程对象并执行任务
			select {
			case taskObj := <-t.taskChannel:
				func() {
					// TODO:可添加recover错误处理
					defer t.wg.Done()
					taskObj.Func(taskObj.ParamObj)
				}()

			case <-t.quit:
				return
			}
		}
	}()
}

func (t *thread) stop() {
	t.quit <- true
}

func newThread(threadPool chan chan *model.Task, wg *sync.WaitGroup) *thread {
	return &thread{
		threadPool:  threadPool,
		taskChannel: make(chan *model.Task),
		quit:        make(chan bool),
		wg:          wg,
	}
}
