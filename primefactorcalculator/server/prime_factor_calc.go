package main

import (
	"log"

	pb "github.com/AdekunleDally/grpc-server-streaming-api/primefactorcalculator/proto"
)

func (s *Server) PrimeFactorCalc(myNumb *pb.NumberRequest, stream pb.PrimeFactorService_PrimeFactorCalcServer) error {
	log.Printf("Prime Factor was invoked: %v\n", myNumb)
	var d int64 = 2
	var q int64 = 0
	for myNumb.InputNumber > 1 {
		if myNumb.InputNumber%d == 0 {

			stream.Send(&pb.PrimeFactorResponse{
				Result: d,
			})

			q = myNumb.InputNumber / d
			myNumb.InputNumber = q
		} else {
			d = d + 1
		}
	}
	return nil

}
