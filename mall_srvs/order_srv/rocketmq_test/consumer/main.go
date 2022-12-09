package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

// 读取消息(消费)
func main() {
	// push模式更加省资源
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"43.143.172.37:9876"}),
		consumer.WithGroupName("mall"), //达到负载的效果
	)
	if err != nil {
		panic(err)
	}

	err = c.Subscribe("imooc1", consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := 0; i < len(ext); i++ {
			fmt.Printf("获取到值: %v\n", ext[i])
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		panic(err)
	}
	c.Start()

	// 不能让主goroutine退出
	time.Sleep(time.Hour)
	c.Shutdown()
}
