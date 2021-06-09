package main

import (
	"context"
	pb "error_test/grpcdemo"
	"fmt"

	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//初始化客户端
	c := pb.NewHelloClient(conn)

	//调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Status = "test status"
	reply, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		fmt.Println(err)
	}
	err, msg := reply.Err.ConvertProto()
	fmt.Println(reply.Err.GetErrs(), reply.Err.GetMessage(), reply.GetData())
}
