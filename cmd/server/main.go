package main

import (
	"flag"
	"log"

	"github.com/artacone/go-url-short/internal/server"
)

func main() {
	mem := flag.String("mem", "in", "in or db")
	flag.Parse()
	switch *mem {
	case "in", "db":
		server.Run(mem)
	default:
		log.Fatalf("wrong mem key: %q. Use \"in\" or \"db\"", *mem)
	}
}
