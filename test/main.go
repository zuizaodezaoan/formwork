package main

import (
	"log"
	"strconv"

	"github.com/zuizaodezaoan/formwork/rabbitmq"
)

func main() {
	for i := 20; i > 10; i-- {
		as := strconv.Itoa(i)
		//简单模式
		//err := rabbitmq.Product("2108A_Simple", []byte("123"+as))
		//if err != nil {
		//	log.Printf(err.Error())
		//	return
		//}

		//广播模式
		err := rabbitmq.Broadcast("2108A_Broadcast", "timeout", []byte("123"+as))
		if err != nil {
			log.Printf(err.Error())
			return
		}
	}

}
