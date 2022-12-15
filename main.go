package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/victorcel/go-enertbit-backend-service/internal/v1/databases"
	"github.com/victorcel/go-enertbit-backend-service/internal/v1/repository"
	"github.com/victorcel/go-enertbit-backend-service/internal/v1/services"
	energyMeterApiv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/services/rest_config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to godotenv: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable TimeZone=America/Bogota",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	dbPostGres, err := databases.NewPostgres(db)

	if err != nil {
		log.Fatalf(err.Error())
	}

	repository.SetEnergyMeterRepository(dbPostGres)

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
	energyMeterApiv1.RegisterEnergyMetersServicesServer(grpcServer, &services.EnergyMetersServer{})

	// start the server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
