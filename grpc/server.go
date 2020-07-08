package grpc

import (
	"context"
	"fmt"
	"github.com/timfame/random-number-service/generator"
	"github.com/timfame/random-number-service/mt"
	"google.golang.org/grpc"
	"net"

	"sync"
)

type Server struct {
	grpc     *grpc.Server
	listener net.Listener
	random   *mt.MT
}

func StartServer(wg *sync.WaitGroup, opts ...Option) (*Server, error) {
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}

	result := &Server{
		grpc:   grpc.NewServer(),
		random: mt.New(mt.WithSeed(config.seed)),
	}
	generator.RegisterRandomNumberGeneratorServer(result.grpc, result)

	var err error
	result.listener, err = net.Listen("tcp", net.JoinHostPort(config.host, fmt.Sprint(config.port)))
	if err != nil {
		return nil, err
	}

	go func() {
		if err := result.grpc.Serve(result.listener); err != nil {
			result.listener.Close()
		}
		wg.Done()
	}()

	return result, nil
}

func (s *Server) Stop() {
	s.grpc.Stop()
}

func (s *Server) Addr() string {
	return s.listener.Addr().String()
}

func (s *Server) GetRandomNumbers(ctx context.Context, request *generator.RandomNumbersRequest) (*generator.RandomNumbersResponse, error) {
	response := &generator.RandomNumbersResponse{Numbers: make([]uint32, request.Number)}
	for i := range response.Numbers {
		response.Numbers[i] = uint32(s.random.Next() % uint64(request.Max))
	}
	return response, nil
}
