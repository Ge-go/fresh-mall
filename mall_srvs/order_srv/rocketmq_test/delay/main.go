package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

// 延时消息
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

	msg := &primitive.Message{
		Topic: "imooc1",
		Body:  []byte("this is imooc1"),
	}
	msg.WithDelayTimeLevel(2)
	res, err := p.SendSync(context.Background(), msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("发送成功: %s\n", res.String())

	err = p.Shutdown()
	if err != nil {
		panic(err)
	}

	//  支付的时候,淘宝,12306,购票, 超时归还  -  定时执行逻辑
	//  轮询的问题,多久执行依次轮询一次  无法达到统一
	//  采用rocketmq的延时队列(仅需查询该订单号)
}
