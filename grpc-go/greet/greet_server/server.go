package main

import (
	"context"
	"fmt"
	"github.com/rollyn/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Printf("Greet function invoked %v", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := "Hello " + firstName + " " + lastName
	res := &greetpb.GreetResponse{
		Result:result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {

	fmt.Printf("GreetManyTimes function invoked %v", req)

	firstName := req.GetGreeting().GetLastName()
	lastName := req.GetGreeting().GetLastName()
	for i := 0; i<10;  i++ {
		result := "Hello " + firstName + " " + lastName + ": " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Failed to receive stream %v", err)
			return err
		}
		firstName := req.GetGreeting().FirstName
		result := "Hello " + firstName + "!"
		sendError := stream.Send(&greetpb.GreetEveryoneResponse{
			Result:result,
		})
		if sendError != nil {
			log.Fatalf("Failed to send stream %v", sendError)
			return sendError
		}
	}
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {

	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result:result,
			})
		}
		if err != nil {
			log.Fatalf("Failed to receive stream %v", err)
		}
		firstName := req.GetGreet().FirstName
		result += "Hello "+ firstName + "\n"
	}


	return nil
}

func main() {

	fmt.Println("executing server...")

	ls, err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(ls); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
