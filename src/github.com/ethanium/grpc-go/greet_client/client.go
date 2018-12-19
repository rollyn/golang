package main 


import ("context"
		"fmt"
		"google.golang.org/grpc"
		"github.com/ethanium/grpc-go/greet/greetpb"
)

func main() {
	
	fmt.Println("client running")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	doUnary(c)
	
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting doUnary....")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rollyn",
			LastName: "Moises",
		},	
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %v", res.Result)

}