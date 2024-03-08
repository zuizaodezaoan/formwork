package main

import (
	"log"

	"github.com/zuizaodezaoan/formwork/rabbitmq"
)

var err error

func main() {
	////普通
	//err = rabbitmq.Consumer("2108A_Simple", Take)
	//if err != nil {
	//	return
	//}

	//广播
	err = rabbitmq.BroadcastConsumption("2108A_Broadcast", "timeout", Take)
	if err != nil {
		return
	}
}

func Take(s string) {
	log.Println(s)
}
