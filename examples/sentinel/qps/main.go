package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
)

func main() {

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
