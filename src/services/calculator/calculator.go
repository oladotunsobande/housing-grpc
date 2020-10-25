package calculator

import (
	"log"

	"golang.org/x/net/context"
)

// Server is the struct definition for the grpc service
type Server struct {
}

// ComputeBreakEven is the grpc service that computes the break even
func (s *Server) ComputeBreakEven(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
