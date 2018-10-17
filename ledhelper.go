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

func (p *prodledwriter) write(top string, bot string, d time.Duration) {
	go lib.Topline(top, d)
	go lib.Bottomline(bot, d)
}

// Init builds the server
func Init() *Server {
	s := &Server{
		&goserver.GoServer{},
		&prodledwriter{},
		int64(0),
	}
	return s
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	var local = flag.Bool("local", false, "Run local test")
	flag.Parse()

	if *local {
		lib.Topline("hello", time.Minute)
		lib.Bottomline("natalie", time.Minute)
		return
	}

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.PrepServer()
	server.Register = server

	server.RegisterServer("led", false)
	server.RunSudo()
	server.Serve()
}
