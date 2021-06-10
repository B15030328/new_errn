package main

import (
	"context"
	"error_test/errn"
	pb "error_test/grpcdemo"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

const (
	//gRPC服务地址
	Address = "127.0.0.1:50052"
)

type helloService struct{}

func (h helloService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	Err := errn.Forbidden
	//Err := status.Error(codes.Aborted, "status error")
	resp.Data = "test data"
	return resp, Err
}

var HelloServer = helloService{}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
	}

	//实现gRPC Server
	s := grpc.NewServer()
	//注册helloServer为客户端提供服务
	pb.RegisterHelloServer(s, HelloServer) //内部调用了s.RegisterServer()
	fmt.Println("Listen on" + Address)

	s.Serve(listen)
}
