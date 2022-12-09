package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

// 普通消息
func main() {
	// 生成producer
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"43.143.172.37:9876"}))
	if err != nil {
		panic(err)
	}

	// 启动producer
	err = p.Start()
	if err != nil {
		panic(err)
	}

	res, err := p.SendSync(context.Background(), &primitive.Message{
		Topic: "imooc1",
		Body:  []byte("this is imooc1"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("发送成功: %s\n", res.String())

	err = p.Shutdown()
	if err != nil {
		panic(err)
	}
}
