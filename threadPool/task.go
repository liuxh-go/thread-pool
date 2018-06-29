package threadPool

import "time"

/*
 * 任务对象
 *
 */

type task func(dtNow time.Time)
