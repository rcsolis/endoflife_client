package rpc

import (
	"context"
	"log"
	"time"

	pb "github.com/rcsolis/endoflife_client/internal/rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
	TIMEOUT = (1 * time.Second)
)

var opts []grpc.DialOption

func init() {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func connect() (*grpc.ClientConn, error) {
	return grpc.NewClient(address, opts...)
}

func GetAll() {
	conn, err := connect()
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCycleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	response, err := client.GetAllLanguages(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get languages: %v", err)
	}

	log.Printf("Languages: %v", response.Languages)
}

func GetAllVersions(name string) {
	conn, err := connect()
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCycleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	responseStream, err := client.GetAllVersions(ctx, &pb.Language{Name: name})
	if err != nil {
		log.Fatalf("could not get versions: %v", err)
	}

	for {
		response, err := responseStream.Recv()
		if err != nil {
			log.Fatalf("could not receive versions: %v", err)
		}
		log.Printf("Item: %v", response)
	}
}

func GetDetails(name string, version string) {
	conn, err := connect()
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCycleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	response, err := client.GetDetails(ctx, &pb.DetailsRequest{Name: name, Version: version})
	if err != nil {
		log.Fatalf("could not get details: %v", err)
	}

	log.Printf("Details: %v", response)
}
