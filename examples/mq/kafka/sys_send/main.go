package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"

	"os"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	p, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer p.Close()
	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		for {
			select {
			case err := <-errors:
				if err != nil {
					glog.Errorln(err)
				}
			case <-p.Successes():
				logx.Info("消息发送producer成功")
			}
		}
	}(p)
	for {
		v := "async: " + strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000))
		fmt.Fprintln(os.Stdout, v)
		msg := &sarama.ProducerMessage{
			Topic: "ffcs",
			Value: sarama.ByteEncoder(v),
		}
		p.Input() <- msg
		time.Sleep(time.Second * 1)
	}
}
