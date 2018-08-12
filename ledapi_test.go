package main

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/brotherlogic/goserver"
	pb "github.com/brotherlogic/led/proto"
)

type testledwriter struct{}

func (p *testledwriter) topline(str string, d time.Duration) {
	log.Printf("%v for %v", str, d)
}

func InitTestServer() *Server {
	s := &Server{
		&goserver.GoServer{},
		&testledwriter{},
	}
	return s
}

func TestDisplay(t *testing.T) {
	s := InitTestServer()
	_, err := s.Display(context.Background(), &pb.DisplayRequest{FullText: "blah"})
	if err != nil {
		t.Errorf("Bad Display")
	}
}
