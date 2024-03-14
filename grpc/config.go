package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RegisterGrpc 函数用于注册并启动 gRPC 服务器
// 参数 host 是服务器主机名或 IP 地址
// 参数 port 是服务器端口号
// 参数 register 是一个函数，用于注册 gRPC 服务
// 如果成功启动服务器，则函数返回 nil；否则返回非 nil 错误
func RegisterGrpc(host string, port int, register func(c *grpc.Server), cert, key string) error {
	// 监听指定的主机和端口
	lister, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		// 如果监听失败，输出错误并返回错误
		log.Fatalf("监听错误: %v", err)
		return err
	}

	n := grpc.NewServer()
	// 注册服务器的反射服务，便于调试
	reflection.Register(n)

	// 调用用户提供的注册函数，注册 gRPC 服务
	register(n)

	// 输出服务器监听地址
	log.Printf("监听服务器地址: %v", lister.Addr())

	// 启动 gRPC 服务器，并开始接受传入的连接
	if err := n.Serve(lister); err != nil {
		// 如果服务器启动失败，输出错误并返回错误
		log.Fatalf("启动失败: %v", err)
		return err
	}

	// 返回 nil 表示服务器启动成功
	return nil
}
