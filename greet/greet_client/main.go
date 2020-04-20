package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/sajanjswl/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("hello i am client")

	cc, err := grpc.Dial("10.101.235.155:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client %f ", c)
	doUnary(c)

	doGreetManyTimes(c)
}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("starting to do Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Veer",
			LastName:  "Jaiswal",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling greet Rpc %v", err)
	}
	log.Printf("Response from greet : %v", res.Result)
}

func doGreetManyTimes(c greetpb.GreetServiceClient) {

	fmt.Println("starting to do a server Streamin  RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Veer",
			LastName:  "Jaiswal",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling greet Rpc %v", err)
	}
	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream %v", err)
		}
		log.Printf("Response from greetManyTimes : %v", msg.Result)
	}

}
