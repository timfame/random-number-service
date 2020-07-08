package grpc

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"randomNumberService/generator"
	"sync"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	server, err := StartServer(&wg, WithPort(8008), WithSeed(uint64(time.Now().Unix())))
	require.NoError(t, err)

	conn, err := grpc.Dial("[::]:8008", grpc.WithInsecure())
	require.NoError(t, err)

	client := generator.NewRandomNumberGeneratorClient(conn)

	response, err := client.GetRandomNumbers(context.Background(), &generator.RandomNumbersRequest{
		Number: 5,
		Max:    14,
	})
	require.NoError(t, err)
	require.Equal(t, 5, len(response.Numbers))
	for _, n := range response.Numbers {
		require.Less(t, n, uint32(14))
	}

	server.Stop()
	wg.Wait()

	response, err = client.GetRandomNumbers(context.Background(), &generator.RandomNumbersRequest{
		Number: 5,
		Max:    14,
	})
	require.Error(t, err)
}
