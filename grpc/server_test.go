package grpc

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/timfame/random-number-service/generator"
	"google.golang.org/grpc"
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

func TestRandom(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	server, err := StartServer(&wg, WithPort(8008), WithSeed(uint64(time.Now().Unix())))
	require.NoError(t, err)

	var testClients sync.WaitGroup
	testClients.Add(100)

	for g := 0; g < 100; g++ {
		go func() {
			for i := 0; i < 100; i++ {
				response, err := server.GetRandomNumbers(context.Background(), &generator.RandomNumbersRequest{
					Number: 1000,
					Max:    1000,
				})
				require.NoError(t, err)
				require.Equal(t, 1000, len(response.Numbers))
			}
			testClients.Done()
		}()
	}
	testClients.Wait()

	server.Stop()
	wg.Wait()
}
