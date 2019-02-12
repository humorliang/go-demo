package mq

import (
	"github.com/nsqio/go-nsq"
	"log"
	"time"
	"fmt"
)

//生产者
func Producer() {
	//nsq配置项
	cfg := nsq.NewConfig()

	//定义一个生产者
	producer, err := nsq.NewProducer("127.0.0.1:4152", cfg)
	if err != nil {
		log.Fatal("producer", err)
	}

	//发送消息
	for true {
		//定义两个topic
		if err := producer.Publish("test", []byte("this is test send msg ")); err != nil {
			log.Fatal("producer error:" + err.Error())
		}
		time.Sleep(2 * time.Second)
	}
}

//消费者
func Consumer() {
	//nsq配置
	cfg := nsq.NewConfig()

	cfg.MaxInFlight = 2
	//消费者  topic ---- 1: n ----> channel
	consumer, err := nsq.NewConsumer("test", "user01", cfg)
	consumer2, err := nsq.NewConsumer("test", "user02", cfg)

	if err != nil {
		log.Fatal("consumer:", err)
	}

	//设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		fmt.Println("test  user01:", string(msg.Body))
		return nil
	}))

	consumer2.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		fmt.Println("test user02:", string(msg.Body))
		return nil
	}))

	// 连接到单例nsqd
	if err := consumer.ConnectToNSQD("127.0.0.1:4152"); err != nil {
		log.Fatal(err)
	}
	if err := consumer2.ConnectToNSQD("127.0.0.1:4152"); err != nil {
		log.Fatal(err)
	}

	//接受停止信号
	<-consumer.StopChan
	<-consumer2.StopChan
}
