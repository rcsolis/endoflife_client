package main

import (
	grpc "github.com/rcsolis/endoflife_client/internal/rpc"
)

func main() {

	grpc.GetDetails("Go", "1.23")

}
