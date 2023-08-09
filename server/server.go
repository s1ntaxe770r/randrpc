package main

import (
	"context"
	"crypto/rand"
	"math/big"
	"net"
	"time"

	"github.com/charmbracelet/log"
	pb "github.com/s1ntaxe770r/randrpc/proto/rand"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = ":7070"
)

type RandServer struct {
	pb.UnimplementedRandServiceServer
}

func (r *RandServer) Rand(ctx context.Context, req *pb.RandRequest) (*pb.RandResponse, error) {
	t := time.Now()
	log.Info("Recived a rand request.....")
	log.Info("Generating random number")

	no, _ := rand.Int(rand.Reader, big.NewInt(int64(req.Max-req.Min)))

	defer func() {
		log.Info("Request took 🔥", time.Since(t))
	}()

	return &pb.RandResponse{
		Rand: int32(no.Int64()),
	}, nil

}

func main() {
	// wire up Grpc stuff
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	grpc := grpc.NewServer()
	reflection.Register(grpc)

	s := &RandServer{}

	pb.RegisterRandServiceServer(grpc, s)

	log.Info("😤 Starting server on", "port", port)
	err = grpc.Serve(lis)
	if err != nil {
		log.Fatal("unable to start rpc server", err.Error())
	}
}

// call the service using grpcurl -plaintext -d '{"min": 1, "max": 100}' localhost:7070 RandService.Rand
