package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

type OrderListener struct {
}

func (o *OrderListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("start exec local logic")
	time.Sleep(time.Second * 3)
	fmt.Println("the logic is error")
	// 本地执行逻辑,无缘无故失败,代码异常,宕机等,这时候回查就起到作用了
	return primitive.UnknowState
}

// CheckLocalTransaction 就算服务挂了,服务器重启之后依然会有一个回查的动作
func (o *OrderListener) CheckLocalTransaction(ext *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("check local reback")
	time.Sleep(time.Second * 3)
	return primitive.CommitMessageState
}

func main() {
	// 生成producer
	p, err := rocketmq.NewTransactionProducer(
		&OrderListener{},
		producer.WithNameServer([]string{"43.143.172.37:9876"}))
	if err != nil {
		panic(err)
	}

	// 启动producer
	err = p.Start()
	if err != nil {
		panic(err)
	}

	res, err := p.SendMessageInTransaction(context.Background(), &primitive.Message{
		Topic: "imooc1",
		Body:  []byte("我打你的母牛"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("send msg success -- %v\n", res.String())
	time.Sleep(time.Hour)
	p.Shutdown()
}
