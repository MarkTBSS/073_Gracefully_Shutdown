package main

import (
	"github.com/MarkTBSS/073_Gracefully_Shutdown/config"
	"github.com/MarkTBSS/073_Gracefully_Shutdown/server"
)

func main() {
	conf := config.ConfigGetting()
	server := server.NewEchoServer(conf)

	server.Start()
}
