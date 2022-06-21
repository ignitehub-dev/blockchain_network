package main

import (
	"flag"
	"log"

	"./server"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 8000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := server.NewBlockchainServer(uint16(*port))
	app.Run()
}
