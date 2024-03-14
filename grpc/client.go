package grpc

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	return conn, err
}
