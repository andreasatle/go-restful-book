// Package main contains a client for a RPC.
package main

import (
	"log"
	"net/rpc"
)

// Args is an empty struct
type Args struct{}

func main() {
	var reply int64
	args := Args{}
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("dialing: %v\n", err)
	}
	err = client.Call("TimeServer.GiveServerTime", args, &reply)
	if err != nil {
		log.Fatalf("arith error: %v\n", err)
	}
	log.Printf("%d", reply)
}
