package main

import (
	"log"

	pb "github.com/AdekunleDally/grpc-server-streaming-api/primefactorcalculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:8081"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewPrimeFactorServiceClient(conn)

	calculatePrimeFactor(c)
}
