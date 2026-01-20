package main

import (
	"grpc/protos/protos/currency"
	"grpc/protos/server"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	grpcServer := grpc.NewServer()
	currencyServer := server.NewCurrency(log)

	currency.RegisterCurrencyServer(grpcServer, currencyServer)

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Error("unable to listen on port", "error", err)
		os.Exit(1)
	}

	grpcServer.Serve(lis)
}
