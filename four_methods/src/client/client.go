package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-demo/four_methods/src/protobuf"
	"io"
	"log"
	"sync"
	"time"
)

var (
	addr = flag.String("addr", "localhost:58888", "the server address to connect")
	data = flag.String("data", "rpctest", "the test data")
)

func UnaryRPC_test(client pb.DemoClient) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// UnaryRPC test
	rpcResponse, err := client.UnaryRPC(ctx, &pb.DemoRequest{Data: *data})
	if err != nil {
		log.Fatalf("failed to client: %v \n %v", err, rpcResponse)
	}
	log.Printf("%v", rpcResponse)
}

func ServerStreamRPC_test(client pb.DemoClient) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// ServerStreamRPC test
	serverStreamRpcResponse, err := client.ServerStreamRPC(ctx, &pb.DemoRequest{Data: *data})
	if err != nil {
		log.Fatalf("failed to client: %v \n %v", err, serverStreamRpcResponse)
	}
	for {
		recv, err := serverStreamRpcResponse.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ServerStreamRPC failed: %v", err)
		}
		log.Printf("client receive %v", recv.Data)
	}
}

func ClientStreamRPC_test(client pb.DemoClient) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// ClientStreamRPC test
	clientStreamRPC, err := client.ClientStreamRPC(ctx)

	if err != nil {
		log.Fatalf("failed to client: %v \n %v", err, clientStreamRPC)
	}
	for _, c := range *data {
		if err := clientStreamRPC.Send(&pb.DemoRequest{Data: fmt.Sprintf("%c", c)}); err != nil {
			log.Fatalf("client.ClientStreamRPC: stream.Send(%v) failed: %v", fmt.Sprintf("%c", c), err)
		}
	}
	response, err := clientStreamRPC.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.ClientStreamRPC failed %v", err)
	}
	log.Printf("client receive %v", response)
}

func BiStreamRPC_test(client pb.DemoClient) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	biStreamRPC, err := client.BiStreamRPC(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()

	// 发送
	go func() {
		defer wg.Done()
		for _, c := range *data {
			err = biStreamRPC.Send(&pb.DemoRequest{Data: fmt.Sprintf("%c", c)})
			if err != nil {
				log.Fatalf("client stream send err %v", err)
			}
			time.Sleep(time.Second)
		}
		err := biStreamRPC.CloseSend()
		if err != nil {
			return
		}
	}()

	// 接收
	go func() {
		defer wg.Done()
		for {
			response, err := biStreamRPC.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("client stream recv %v", err)
			}
			log.Printf("client recv %v", response)
		}
	}()
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to client: %v", err)
		}
	}(conn)
	client := pb.NewDemoClient(conn)

	//UnaryRPC_test(client)
	//ServerStreamRPC_test(client)
	//ClientStreamRPC_test(client)
	BiStreamRPC_test(client)
}
