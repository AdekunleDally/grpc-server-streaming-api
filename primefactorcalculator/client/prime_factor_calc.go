package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/AdekunleDally/grpc-server-streaming-api/primefactorcalculator/proto"
)

func calculatePrimeFactor(c pb.PrimeFactorServiceClient) {
	var input_number int64
	var primeFactorArray []int64
	log.Println("CalculatePrimeFactor was invoked")

	log.Println("Please Enter the Number which you want to find its Prime Factors:")
	fmt.Scanln(&input_number)
	req := &pb.NumberRequest{
		InputNumber: input_number,
	}

	stream, err := c.PrimeFactorCalc(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to call the prime Factors : %v\n ", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Error while reading the prime factor data stream")
		}

		//log.Print("The PrimeFactors are: ", msg.Result)
		primeFactorArray = append(primeFactorArray, msg.Result)

	}
	log.Println("The Prime Factors are:", primeFactorArray)
}
