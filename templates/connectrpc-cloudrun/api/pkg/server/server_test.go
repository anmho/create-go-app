package server

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	greetv1 "github.com/anmho/cloud-run-connectrpc/gen/protos/greet/v1"
	"github.com/stretchr/testify/assert"
)


func TestServer_Greet(t *testing.T) {
	tests := []struct{
		desc string
		request *connect.Request[greetv1.GreetRequest]
		expectedResponse *greetv1.GreetResponse
	}{
		{
			desc: "happy path: greets caller",
			request: &connect.Request[greetv1.GreetRequest]{
				Msg: &greetv1.GreetRequest{
					Name: "Joe",
				},
			},
			expectedResponse: &greetv1.GreetResponse{
				Greeting: "Hello, Joe!",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func (t *testing.T)  {
			s := New()

			resp, err := s.Greet(context.Background(), tc.request)

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, tc.expectedResponse, resp.Msg)
		})
	}
}