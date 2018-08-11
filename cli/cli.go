package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/brotherlogic/goserver/utils"
	"google.golang.org/grpc"

	pbgs "github.com/brotherlogic/goserver/proto"
	pb "github.com/brotherlogic/led/proto"
	pbt "github.com/brotherlogic/tracer/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
)

func main() {
	host, port, err := utils.Resolve("led")
	if err != nil {
		log.Fatalf("Unable to reach organiser: %v", err)
	}
	conn, err := grpc.Dial(host+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}

	client := pb.NewLedServiceClient(conn)
	ctx, cancel := utils.BuildContext("led-cli", pbgs.ContextType_LONG)
	defer cancel()

	_, err = client.Display(ctx, &pb.DisplayRequest{FullText: os.Args[1]})
	if err != nil {
		log.Fatalf("Error on GET: %v", err)
	}

	utils.SendTrace(ctx, "End of CLI", time.Now(), pbt.Milestone_END, "led-cli")
}
