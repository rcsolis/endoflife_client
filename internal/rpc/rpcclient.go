package rpc

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	e "github.com/rcsolis/endoflife_client/internal/error"
	"github.com/rcsolis/endoflife_client/internal/model"
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

/**
 * GetAll fetches all the languages from the remote gRPC server
 */
func GetAll() ([]model.Technology, error) {
	var elements []model.Technology

	conn, err := connect()
	if err != nil {
		log.Printf("%v:could not connect: %v", e.ConnectionError, err)
		return nil, e.Throw(
			e.ConnectionError,
			&e.GRPCError{Msg: fmt.Sprintf("could not connect: %v", err)},
		)
	}
	defer conn.Close()

	client := pb.NewCycleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	response, err := client.GetAllLanguages(ctx, &pb.Empty{})
	if err != nil {
		log.Printf("%v:could not connect: %v", e.GetAllError, err)
		return nil, e.Throw(
			e.GetAllError,
			&e.GRPCError{Msg: fmt.Sprintf("could not get tecnologies: %v", err)},
		)
	}

	for _, element := range response.Languages {
		elements = append(elements, model.Technology{Name: element.Name})
	}
	return elements, nil
}

/**
 * GetAllVersions fetches all the versions of a language from the remote gRPC server
 * @param name string
 */
func GetAllVersions(name string) error {
	var counter = 0
	conn, err := connect()
	if err != nil {
		log.Printf("%v:could not connect: %v", e.ConnectionError, err)
		return e.Throw(
			e.ConnectionError,
			&e.GRPCError{Msg: fmt.Sprintf("could not get tecnologies: %v", err)},
		)
	}
	defer conn.Close()

	client := pb.NewCycleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	responseStream, err := client.GetAllVersions(ctx, &pb.Language{Name: name})
	if err != nil {
		log.Printf("%v:could not get versions: %v", e.GetVersionsError, err)
		return e.Throw(
			e.GetVersionsError,
			&e.GRPCError{Msg: fmt.Sprintf("could not get tecnologies: %v", err)},
		)
	}
	log.Printf("Stream Starts, counter: %d, elements: %d", counter, len(model.TechnologiesCycle))
	for {
		response, err := responseStream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Stream Ends: Stream ends successfully: %v, total found %d", err, counter)
				return nil
			} else {
				log.Printf("%v: data stream closed with error: %v", e.StreamEOF, err)
				return e.Throw(
					e.StreamEOF,
					&e.GRPCError{Msg: fmt.Sprintf("data stream closed with error: %v", err)},
				)
			}
		}
		counter++
		model.TechnologiesCycle = append(model.TechnologiesCycle, model.LanguageCycle{
			Cycle:           response.GetCycle(),
			ReleaseDate:     response.GetReleaseDate(),
			Eol:             response.GetEol(),
			Latest:          response.GetLatest(),
			Link:            response.GetLink(),
			Lts:             response.GetLts(),
			Support:         response.GetSupport(),
			Discontinued:    response.GetDiscontinued(),
			ExtendedSupport: response.GetExtendedSupport(),
		})
	}
}

/**
 * GetDetails fetches the details of a language cycle from the remote gRPC server
 * @param name string
 * @param version string
 */
func GetDetails(name string, version string) (model.LanguageCycle, error) {
	conn, err := connect()
	if err != nil {
		log.Printf("%v:could not connect: %v", e.ConnectionError, err)
		return model.LanguageCycle{}, e.Throw(
			e.ConnectionError,
			&e.GRPCError{Msg: fmt.Sprintf("could not get tecnologies: %v", err)},
		)
	}
	defer conn.Close()

	client := pb.NewCycleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	response, err := client.GetDetails(ctx, &pb.DetailsRequest{Name: name, Version: version})
	if err != nil {
		log.Printf("%v:could not get the cycle details: %v", e.GetDetailsError, err)
		return model.LanguageCycle{}, e.Throw(
			e.GetDetailsError,
			&e.GRPCError{Msg: fmt.Sprintf("could not get the cycle details: %v", err)},
		)
	}
	log.Printf("Response: %v", response)
	return model.LanguageCycle{
		Cycle:           response.GetCycle(),
		ReleaseDate:     response.GetReleaseDate(),
		Eol:             response.GetEol(),
		Latest:          response.GetLatest(),
		Link:            response.GetLink(),
		Lts:             response.GetLts(),
		Support:         response.GetSupport(),
		Discontinued:    response.GetDiscontinued(),
		ExtendedSupport: response.GetExtendedSupport(),
	}, nil
}
