package main 

import ("context"
		"log"
		"fmt"
		"net"
		"google.golang.org/grpc"
		"github.com/ethanium/grpc-go/greet/greetpb"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	log.Printf("Greet function invoked with %v",req)	
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello "+firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}


func main() {
	
	fmt.Println("hello world")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}