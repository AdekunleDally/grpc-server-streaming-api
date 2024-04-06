package main

import (
	"log"
	"net"

	pb "github.com/AdekunleDally/grpc-server-streaming-api/primefactorcalculator/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.PrimeFactorServiceServer
}

var addr string = "0.0.0.0:8081"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen to : %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()

	pb.RegisterPrimeFactorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}
