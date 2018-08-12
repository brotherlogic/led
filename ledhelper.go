// +build !maintest

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/led/lib"
)

type prodledwriter struct{}

func (p *prodledwriter) topline(str string, d time.Duration) {
	lib.Topline(str, d)
}

// Init builds the server
func Init() *Server {
	s := &Server{
		&goserver.GoServer{},
		&prodledwriter{},
	}
	return s
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.PrepServer()
	server.Register = server

	server.RegisterServer("led", false)
	server.Serve()
}
