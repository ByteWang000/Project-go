package main

import (
	"context"
	"fmt"
	"time"
)

// context 用来解决 goroutine 之间退出通知、元数据传递的功能。

func main() {

}
func context_learn() {
	// Background是任何Context树的根,它永远不会被取消:
	// 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
	ctx := context.Background() //返回一个空的Context。 它永远不会取消，没有截止日期，没有价值。
	// WithCancel返回一个父进程的副本，该父进程的Done通道被尽快关闭。 关闭Done或调用cancel。
	ctx, cancel := context.WithCancel(ctx)
	cancel()

	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))

}

// 轮询外卖小哥位置
func find_location() {
	// WithTimeout函数来创建一个具有超时功能（超过指定时间段，上下文自动取消）的上下文。
	// cancel：取消函数，用于手动取消上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	go func(ctx1 context.Context) {
		for {
			// 计算位置
			// 发送位置
			select {
			case <-ctx1.Done():
				// 被取消，直接返回
				return
			case <-time.After(time.Second):
				// block 1 秒
			}
		}
	}(ctx)
	cancel() // 调用后——>ctx1.Done()返回的通道有struct{}，从而return
}
func context_list() {
	// 创建了一个空的上下文父ctx。
	ctx := context.Background()
	// 此时ctx还没有写入值
	process(ctx)
	// 创建了一个衍生的ctx，并使 key=traceId,value=317722950
	ctx = context.WithValue(ctx, "traceId", "317722950")
	process(ctx)
}

func process(ctx context.Context) {
	// 获取上下文中键为"traceId"的值
	traceId, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Printf("process over.trace_id = %s\n", traceId)
	} else {
		fmt.Println("process over. no trace_id !")
	}

}
