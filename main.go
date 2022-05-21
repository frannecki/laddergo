package main

import (
	"github.com/frannecki/laddergo/handler"
	"github.com/frannecki/laddergo/server"
)

func main() {
	http_server := server.NewServer(uint16(8070))
	handler := handler.NewGoHandler()
	http_server.SetRequestCallback(handler)
	http_server.Start()
}
