package services

import (
	"context"
	"github.com/victorcel/go-enertbit-backend-service/internal/v1/repository"
	"github.com/victorcel/go-enertbit-backend-service/internal/v1/useCases"
	energyMeterv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/energyMeter"
	energyMeterApiv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/services/rest_config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errorPath = "ERR internal/services/services.go"

type EnergyMetersServer struct {
	energyMeterApiv1.EnergyMetersServicesServer
}

func (EnergyMetersServer) GetEnergyMeters(_ context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseGetEnergyMeters, error) {
	//TODO implement me
	panic("implement me")

}

func (EnergyMetersServer) CreateEnergyMeters(ctx context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseEnergyMeter, error) {
	_, err := useCases.Insert(meter.EnergyMeter)

	if err != nil {
		err := status.Error(codes.NotFound, err.Error())
		return nil, err
	}

	response := &energyMeterApiv1.ResponseEnergyMeter{}

	response.EnergyMeter = meter.GetEnergyMeter()

	return response, nil
}

func (EnergyMetersServer) UpdateEnergyMeters(ctx context.Context, meters *energyMeterApiv1.RequestUpdateEnergyMeters) (*energyMeterv1.BoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (EnergyMetersServer) DeleteEnergyMeters(ctx context.Context, request *energyMeterv1.IdRequest) (*energyMeterv1.BoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (EnergyMetersServer) InstalledCutOrInactiveEnergyMeter(_ context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseGetEnergyMeters, error) {
	find, err := repository.Find(meter.EnergyMeter)

	if err != nil {
		return nil, err
	}

	return find, err
}

func (EnergyMetersServer) RecentInstallationEnergyMeter(ctx context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseEnergyMeter, error) {
	//TODO implement me
	panic("implement me")
}
