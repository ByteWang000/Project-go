package cond_learn

import (
	"fmt"
	"sync"
)

// 用于实现在指定条件下阻塞和唤醒协程的操作.
func Hello() {
	fmt.Println("OK！！！")
}
func Cond_learn() {
	// 创建实例
	cond := sync.NewCond(&sync.Mutex{})
	// Broadcast 唤醒所有等待条件变量 c 的 goroutine，无需锁保护。
	cond.Broadcast()
	// Signal 只唤醒任意1个等待条件变量 c 的 goroutine，无需锁保护。
	cond.Signal()

	cond.L.Lock() // 关联的互斥锁加锁
	for !(1 == 1) {
		// 条件不满足在此等待
		// 当前goroutine会被阻塞，直到其它goroutine调用cond.Signal()或cond.Broadcast()来通知条件满足
		cond.Wait() // 自动释放c.L挂起调用者所在goroutine。
	}
	// 执行需要在条件满足时进行的操作
	// do your work here !
	cond.L.Unlock()
}
