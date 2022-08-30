package main

import (
	"fmt"
	"testing"
	"zero-mal/common/tool"
)

func main() {
	chs := make(chan struct{})
	worker, err := tool.NewOrderWorker(0)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 105; i++ {
		go func() {
			id := worker.NextOrderId()
			fmt.Println(id)
		}()

	}
	fmt.Println(1)
	<-chs
	fmt.Println(1)
}

// 结果：
// 7783621591040
// 7783621591041
// 7783621591042
// 7783621591043
// 7783621591044

func BenchmarkID(b *testing.B) {
	worker, _ := tool.NewOrderWorker(0)
	for i := 0; i < b.N; i++ {
		worker.NextOrderId()
	}
}

//BenchmarkID-16           4902658（执行次数）              244.5 ns/op（平均每次执行所需时间）
