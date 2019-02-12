package main

import "mq/gnsq/mq"

func main() {
	go mq.Consumer()
	mq.Producer()
}
