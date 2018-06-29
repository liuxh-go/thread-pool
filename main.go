package main

import (
	"fmt"
	"time"
	"wshhz.com/goThreadPool/threadPool"
)

func main() {
	dispather := threadPool.NewDispather(20)
	dispather.Run()

	for i := 0; i < 50; i++ {
		go func() {
			dispather.AddTask(test)
		}()
	}

	time.Sleep(time.Second * 20)
}

func test(dtNow time.Time) {
	time.Sleep(1 * time.Second)

	fmt.Println(dtNow)
}
