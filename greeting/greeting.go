package greeting

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SendGreetings(ctx context.Context, greeting *Greeting) (*Greeting, error) {

	log.Printf("Received Greeting from the client From: %v, To: %v", greeting.Content, greeting.To)

	return &Greeting{Content: "Roger!"}, nil

}
