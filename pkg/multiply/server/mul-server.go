package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"gRPC-Tutorial/api/multiply"

	"google.golang.org/grpc"
)

var (
	port  = flag.Int("port", 50051, "The port to listen on")
	value = 1
)

type server struct {
	multiply.UnimplementedMultiplierServer
}

func (s *server) Mul(ctx context.Context, in *multiply.MulReq) (*multiply.MulRes, error) {
	number, err := strconv.Atoi(in.X)
	if err != nil {
		log.Fatalf("InvalidArgument")
	}

	value *= number
	return &multiply.MulRes{Result: strconv.Itoa(value)}, nil
}

func main() {
	flag.Parse()
	con, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Println("PORT =", *port)
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	multiply.RegisterMultiplierServer(s, &server{})
	log.Printf("server listening at %v", con.Addr())
	if err := s.Serve(con); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
