# Prime FactorCalculator Service

This application provides a calculator service using a gRPC server streaming API to perform by taking in an input from a user and returns the prime factor(s) of that number.Through the use of gRPC protocolbuffer, it defines and takes a message request and returns a stream of Responses that represent the prime number decomposition of that number.
## Setup

This project requires a `gcc` compiler installed and the `protobuf` code generation tools.

### Install protobuf compiler

Install the `protoc` tool using the instructions available at [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/).

## Generating gRPC Code 

To generate the gRPC code, run the following command:
 
```bash
protoc -Iprimefactorcalculator/proto --go_out=. --go_opt=module=github.com/AdekunleDally/grpc-server-streaming-api --go-grpc_out=. --go-grpc_opt=module=github.com/AdekunleDally/grpc-server-streaming-api primefactorcalculator/proto/pf_calculator.proto 

protoc -Icalculator/proto --go_out=. --go_opt=module=github.com/AdekunleDally/grpc-unary-api --go-grpc_out=. --go-grpc_opt=module=github.com/AdekunleDally/grpc-unary-api calculator/proto/calculator.proto
