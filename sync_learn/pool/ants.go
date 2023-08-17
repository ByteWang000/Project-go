package pool

//受限于服务器硬件内存大小，如果不对goroutine数量进行限制，会出现Out of Memory错误
//ants是一个高性能的 goroutine 池

// goroutine 只能自己运行结束，外部没有任何手段可以强制j结束一个 goroutine。
// 如果一个 goroutine 因为某种原因没有自行结束，就会出现 goroutine 泄露。此外，频繁创建 goroutine 也是一个开销。

// go get -u github.com/panjf2000/ants/v2

//func ants_learn() {
//	// 创建默认的协程池
//	pool, err := ants.NewPool(10) // 10 是协程池的容量
//
//	// 创建带有自定义任务函数的协程池
//	poolfunc, err := ants.NewPoolWithFunc(10, func(task interface{}) {
//		// 处理任务的逻辑
//	})
//	//使用 Submit() 方法将任务提交给协程池
//	err := pool.Submit(func() {
//		// 处理任务的逻辑
//	})
//	// 关闭协程池
//	defer pool.Release()
//
//}
