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

func (p *testledwriter) write(top, bot string, d time.Duration) {
	log.Printf("%v and %v for %v", top, bot, d)
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
	_, err := s.Display(context.Background(), &pb.DisplayRequest{TopLine: "blah", BottomLine: "blah2"})
	if err != nil {
		t.Errorf("Bad Display")
	}
}
