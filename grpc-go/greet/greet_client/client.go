package main

import (
	"context"
	"fmt"
	"log"
	"github.com/rollyn/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("executing client...")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect %v", err)
	}

	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	//fmt.Printf("created client %f",c)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting:&greetpb.Greeting{
			FirstName:"ROLLYN",
			LastName:"MOISES",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet rpc %v",err)
	}
	fmt.Printf(res.Result)
}
