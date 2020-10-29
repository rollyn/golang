package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"github.com/rollyn/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
	"time"
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
//	doUnary(c)

	//doServerStreaming(c)
	//doClientStreaming(c)
	doBidirectionalStream(c)
}

func doBidirectionalStream(c greetpb.GreetServiceClient) {
	fmt.Println("executing doBidirectionalStream...")

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting:&greetpb.Greeting{
				FirstName:"ROLLYN",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting:&greetpb.Greeting{
				FirstName:"JOY",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting:&greetpb.Greeting{
				FirstName:"ETHAN",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting:&greetpb.Greeting{
				FirstName:"CHENILLE",
			},
		},
	}

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("unable to stream %v", err)
		return
	}
	wc := make(chan struct{})

	go func() {
		// send
		for _,req := range requests {
			fmt.Printf("Sending request %v\n",req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	go func() {
		// receive
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("unable to receive %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res)
		}
		close(wc)
	}()
	<-wc
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("executing doClientStreaming...")
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greet:&greetpb.Greeting{
				FirstName:"ROLLYN",
			},
		},
		&greetpb.LongGreetRequest{
			Greet:&greetpb.Greeting{
				FirstName:"JOY",
			},
		},
		&greetpb.LongGreetRequest{
			Greet:&greetpb.Greeting{
				FirstName:"ETHAN",
			},
		},
		&greetpb.LongGreetRequest{
			Greet:&greetpb.Greeting{
				FirstName:"CHENILLE",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("unable to stream %v", err)
	}
	for _, req := range requests {
		fmt.Printf("Sending request %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving %v", err)
	}
	fmt.Printf("Response: %v\n", res)
}
func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("executing doServerStreaming...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "ROLLYN",
			LastName:"MOISES",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("unable to receive %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// close stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream %v", err)
		}
		fmt.Printf("Response: %v\n",msg.GetResult())
	}

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
