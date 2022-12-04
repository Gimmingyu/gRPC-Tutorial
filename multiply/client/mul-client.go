package main

import (
	"context"
	"flag"
	"gRPC-Tutorial/api/multiply"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultValue = "10"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
	x    = flag.String("x", defaultValue, "10")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := multiply.NewMultiplierClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Mul(ctx, &multiply.MulReq{X: *x})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResult())
}
