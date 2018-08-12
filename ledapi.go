package main

import (
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/led/proto"
)

// Display shows the text
func (s *Server) Display(ctx context.Context, req *pb.DisplayRequest) (*pb.DisplayResponse, error) {
	s.ledwriter.topline(req.GetFullText(), time.Minute)
	return &pb.DisplayResponse{}, nil
}
