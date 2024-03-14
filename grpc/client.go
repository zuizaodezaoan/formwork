package grpc

import (
	"fmt"
	"log"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"

	"github.com/zuizaodezaoan/formwork/consul"
	"github.com/zuizaodezaoan/formwork/nacos"
)

func RegisterApi(serverName string) (*grpc.ClientConn, error) {
	server, i, err := consul.GetServer(serverName)
	if err != nil {
		return nil, err
	}
	log.Println("consul连接+++++++++", server, i)

	host, port, err := nacos.GetNacosSrv()
	if err != nil {
		return nil, err
	}
	log.Println("54343234543", host, port)
	conn, err := grpc.Dial(fmt.Sprintf("consul://192.168.18.94:8500/"+"shop"+"?wait=14s", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	return conn, err
}
