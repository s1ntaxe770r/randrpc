package main

import (
	"context"
	"github.com/charmbracelet/log"
	pb "github.com/s1ntaxe770r/randrpc/proto/rand"
	"google.golang.org/grpc"
	"os"
	"time"
)

var (
	serverAddr = os.Getenv("SERVER_ADDR")
)

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRandServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	for {
		// Make a request to the server
		response, err := client.Rand(ctx, &pb.RandRequest{Min: 1, Max: 100})
		if err != nil {
			log.Error(err.Error())
		} else {
			log.Info("Random number", "=", response.Rand)
		}

		log.Info("Random number", "", response.Rand)
		// Sleep for 1 minute before making the next request
		time.Sleep(time.Second)
	}
}
