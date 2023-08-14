package pool

import (
	"fmt"
	"sync"
)

//Pool 池里的元素随时可能释放掉，释放策略完全由 runtime 内部管理；
//Get 获取到的元素对象可能是刚创建的，也可能是之前创建好 cache 住的。使用者无法区分；
//Pool 池里面的元素个数你无法知道；

// 大小可伸缩，动态扩容，存在池中不活跃对象会被自动清理
type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

func Pool_learn() {
	// 声明对象池
	var studentPool = sync.Pool{
		New: func() interface{} {
			return new(Student)
		},
	}
	// Get() 用于从对象池中获取对象，因为返回值是 interface{}，因此需要断言
	stu := studentPool.Get().(*Student)
	// Put() 则是在对象使用完毕后，返回对象池
	defer studentPool.Put(stu)

}
func Hello() {
	fmt.Println("OK！！！")
}
