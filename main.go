package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/victorcel/go-enertbit-backend-service/internal/v1/energyMeter"
	energyMeterApiv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/services/rest_config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to godotenv: %v", err)
	}

	// create a listener on TCP port
	lis, err := net.Listen("tcp", ":"+os.Getenv("GO_BACKEND_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Print("Run Server: " + os.Getenv("GO_BACKEND_SERVICE_PORT"))

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	// attach the Process service to the server
	energyMeterApiv1.RegisterEnergyMetersServicesServer(grpcServer, &energyMeter.EnergyMetersServer{})

	// start the server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
