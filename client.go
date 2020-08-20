package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpcDemo/grpcMod"
	"os"
)

const (
	address = ":9090"
)

func main()  {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
    	fmt.Println("connect error:", err)
    	os.Exit(1)
	}
	defer conn.Close()
    cli := pb.NewHelloServiceClient(conn)
    res, err := cli.SayHello(context.Background(), &pb.HelloRequest{Greeting:"rpc"})
    if err != nil {
    	fmt.Println("error:", err)
	}
	print(res.Reply)
}

