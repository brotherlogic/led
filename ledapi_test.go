package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/led/proto"
)

func InitTestServer() *Server {
	s := Init()
	return s
}

func TestDisplay(t *testing.T) {
	s := InitTestServer()
	_, err := s.Display(context.Background(), &pb.DisplayRequest{FullText: "blah"})
	if err == nil {
		t.Errorf("Bad Move")
	}
}
