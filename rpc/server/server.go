// Package main contains a server for a RPC.
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// Args is an empty struct
type Args struct{}

// TimeServer is an int64
type TimeServer int64

// GiveServerTime gives the time on the server
func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	fmt.Println("GiveServerTime")
	*reply = time.Now().Unix()
	return nil
}

func main() {
	timeServer := new(TimeServer)
	rpc.Register(timeServer)
	rpc.HandleHTTP()
	server, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	http.Serve(server, nil)
}
