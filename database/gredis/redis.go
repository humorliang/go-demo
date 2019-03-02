package main

import (
	"github.com/go-redis/redis"
	"github.com/smallnest/rpcx/log"
	"fmt"
	"time"
)

func NewClient() (*redis.Client, error) {
	/*
	Redis服务相关的命令:
	select 选择数据库(数据库编号0-15)
	quit 退出连接
	info 获得服务的信息与统计
	monitor 实时监控
	config get 获得服务配置
	flushdb 删除当前选择的数据库中的key
	flushall 删除所有数据库中的key
	*/
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	//连接
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
func main() {
	c, err := NewClient()
	if err != nil {
		log.Fatal("client error:", err)
	}
	err = c.Set("key", "value", time.Duration(5*time.Second)).Err()
	if err != nil {
		log.Fatal(err)
	}
	val, err := c.Get("key").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key:", val)
	val2, err := c.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 is not exsit")
	} else if err != nil {
		fmt.Println("key error:", err)
	} else {
		fmt.Println("key2:", val2)
	}
}
