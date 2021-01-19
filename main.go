package main

import (
	"fmt"
	"time"

	"github.com/liuxh-go/thread-pool/model"
	pool "github.com/liuxh-go/thread-pool/taskpool"
)

func main() {
	dispather := pool.NewDispather(5)
	dispather.Run()

	for i := 0; i < 50; i++ {
		dispather.AddTask(model.NewTask(test, &model.Param{A: int32(i)}))
	}

	dispather.WaitStop()
}

func test(paramObj *model.Param) {
	fmt.Println(paramObj.A)
	time.Sleep(time.Second * 5)
}
