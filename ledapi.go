package main

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/led/proto"
)

// Display shows the text
func (s *Server) Display(ctx context.Context, req *pb.DisplayRequest) (*pb.DisplayResponse, error) {
	return nil, fmt.Errorf("Not Implemented")
}
