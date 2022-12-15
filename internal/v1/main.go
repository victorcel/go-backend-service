package main

import (
	"github.com/victorcel/go-enertbit-backend-service/internal/v1/energyMeter"
	energyMeterApiv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/services/rest_config"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", ":"+os.Getenv("GO_BACKEND_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Process service to the server
	energyMeterApiv1.RegisterEnergyMetersServicesServer(grpcServer, &energyMeter.Server{})

	// start the server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
