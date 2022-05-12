package main

import "github.com/frannecki/laddergo/server"

func main() {
	server := server.NewServer(uint16(8070))
	server.Start()
}
