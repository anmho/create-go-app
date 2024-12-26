package server

import (
	"context"
	"fmt"
	"log"

	"connectrpc.com/connect"
	greetv1 "github.com/anmho/cloud-run-connectrpc/gen/protos/greet/v1"
	"github.com/anmho/cloud-run-connectrpc/gen/protos/greet/v1/greetv1connect"
)

var _ greetv1connect.GreetServiceHandler = (*GreetServer)(nil)
type GreetServer struct{}

func New() *GreetServer {
	return &GreetServer{}
}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}