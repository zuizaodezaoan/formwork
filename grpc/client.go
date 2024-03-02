package grpc

import (
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/zuizaodezaoan/formwork/config"
	"github.com/zuizaodezaoan/formwork/nacos"
)

func RegisterApi(serverName string) (*grpc.ClientConn, error) {
	//_, err := credentials.NewClientTLSFromFile("./cert.pem", "x.test.example.com")
	//if err != nil {
	//	log.Fatalf("failed to load credentials: %v", err)
	//	return nil, err
	//}
	nacosConfig, err := nacos.NacosConfig(serverName)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(nacosConfig), &config.Usersrv)
	if err != nil {
		return nil, err
	}
	fmt.Println("Registe", config.Usersrv.Host, config.Usersrv.Port)
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Usersrv.Host, config.Usersrv.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	return conn, err
}
