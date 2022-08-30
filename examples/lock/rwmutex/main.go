package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
)

func main() {
	//var (
	//	lock sync.Mutex
	//	m    = make(map[int]int)
	//)

	//lock.Lock()
	//defer lock.Unlock()
	//for i := 0; i < 100; i++ {
	//	i := i
	//	go func() {
	//		lock.Lock()
	//		fmt.Println("lock ok")
	//		m[i] = i
	//		lock.Unlock()
	//		fmt.Println("lock Unlock")
	//	}()
	//}
	//time.Sleep(time.Second * 5)
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}

	chanClose := make(chan struct{})

	for i := 0; i < 100; i++ {
		go func() {
			ee, bb := sentinel.Entry("goods-list", sentinel.WithTrafficType(base.Inbound))
			if bb != nil {
				//logx.Error("限流了")
				fmt.Println("限流了")
				return
			} else {
				fmt.Println("正常")
			}
			ee.Exit()
		}()
	}

	<-chanClose
}
