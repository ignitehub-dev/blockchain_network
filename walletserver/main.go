package main

import (
	"flag"
	"log"

	"./server"
)

func init() {
	log.SetPrefix("Wallet Server: ")
}

func main() {
	port := flag.Uint("port", 8080, "TCP Port Number for Wallet Server")
	gateway := flag.String("gateway", "http://0.0.0.0:8000", "Blockchain Gateway")
	flag.Parse()

	app := server.NewWalletServer(uint16(*port), *gateway)
	app.Run()
}
