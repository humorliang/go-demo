package kafka

import (
	"github.com/Shopify/sarama"
	"fmt"
)

// 同步模式
func SyncProducer(addr string) {
	config := sarama.NewConfig()
	fmt.Println(config)
}
