package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    pb "grpcDemo/grpcMod"
    "net"
)

const (
    port = ":9090"
)

type server struct {}

func (s *server) SayHello(context.Context, *pb.HelloRequest) (*pb.HelloResponse, error) {
    return &pb.HelloResponse{Reply:"hello word!!!"}, nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        fmt.Println("start service error and error reason is :", err)
    }
    s := grpc.NewServer()
    pb.RegisterHelloServiceServer(s, &server{})
    fmt.Println("service start *************")
    if err = s.Serve(lis); err != nil {
        fmt.Println("start service error and error reason is :", err)
    }
}
