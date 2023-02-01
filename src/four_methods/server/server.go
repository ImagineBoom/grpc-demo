package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-demo/src/four_methods/protobuf"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

var (
	port = flag.Int("port", 58888, "The server port")
)

type server struct {
	pb.UnimplementedDemoServer
}

func (s *server) UnaryRPC(ctx context.Context, req *pb.DemoRequest) (*pb.DemoResponse, error) {
	log.Printf("server receive: %v", req.Data)
	return &pb.DemoResponse{Data: req.GetData() + "-server"}, status.Errorf(codes.OK, "method UnaryRPC not implemented")
}
func (s *server) ServerStreamRPC(req *pb.DemoRequest, stream pb.Demo_ServerStreamRPCServer) error {
	log.Printf("server receive: %v", req.Data)
	for _, c := range req.Data + "server" {
		if err := stream.Send(&pb.DemoResponse{Data: fmt.Sprintf("%c", c)}); err != nil {
			return err
		}
	}
	return status.Errorf(codes.OK, "method ServerStreamRPC not implemented")
}

func (s *server) ClientStreamRPC(stream pb.Demo_ClientStreamRPCServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&pb.DemoResponse{
				Data: endTime.String(),
			})
		}
		if err != nil {
			return err
		}
		log.Printf("server receive %s", request.Data)
	}
}

func (s *server) BiStreamRPC(stream pb.Demo_BiStreamRPCServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()

	// 接收
	go func() {
		defer wg.Done()
		for {
			request, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("server recv %v", err)
			}
			log.Printf("server recv %v", request)
		}
	}()

	// 发送
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			err := stream.Send(&pb.DemoResponse{Data: fmt.Sprintf("server send test %d", i)})
			if err != nil {
				log.Fatalf("server send err %v", err)
			}
			time.Sleep(time.Second)
		}
	}()

	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDemoServer(grpcServer, &server{})
	log.Printf("server listening at %v", lis.Addr())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
