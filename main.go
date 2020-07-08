package main

import (
	"flag"
	"fmt"
	"github.com/timfame/random-number-service/grpc"
	"os"
	"sync"
	"time"
)

func main() {
	host := flag.String("host", "", "gRPC host")
	port := flag.Uint64("port", 8008, "gRPC port")
	seed := flag.Uint64("seed", uint64(time.Now().Unix()), "seed for random")

	flag.Parse()

	opts := []grpc.Option{
		grpc.WithHost(*host),
		grpc.WithPort(uint32(*port)),
		grpc.WithSeed(*seed),
	}

	var wg sync.WaitGroup
	wg.Add(1)

	_, err := grpc.StartServer(&wg, opts...)
	fmt.Println("Server started...")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wg.Wait()
	fmt.Println("Server stopped")
}
