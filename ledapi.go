package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/led/proto"
)

// Display shows the text
func (s *Server) Display(ctx context.Context, req *pb.DisplayRequest) (*pb.DisplayResponse, error) {
	go s.ledwriter.write(req.GetTopLine(), req.GetBottomLine(), time.Minute)
	s.received++
	s.Log(fmt.Sprintf("WRITING %v,%v", req.GetTopLine(), req.GetBottomLine()))
	return &pb.DisplayResponse{}, nil
}
