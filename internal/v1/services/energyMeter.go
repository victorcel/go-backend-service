package services

import (
	"context"
	energyMeterv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/energyMeter"
	energyMeterApiv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/services/rest_config"
)

var errorPath = "ERR internal/services/services.go"

type EnergyMetersServer struct {
	energyMeterApiv1.EnergyMetersServicesServer
}

func (EnergyMetersServer) GetEnergyMeters(ctx context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseGetEnergyMeters, error) {
	//TODO implement me
	panic("implement me")
}

func (EnergyMetersServer) CreateEnergyMeters(ctx context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseEnergyMeter, error) {
	//TODO implement me
	panic("implement me")
}

func (EnergyMetersServer) UpdateEnergyMeters(ctx context.Context, meters *energyMeterApiv1.RequestUpdateEnergyMeters) (*energyMeterv1.BoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (EnergyMetersServer) DeleteEnergyMeters(ctx context.Context, request *energyMeterv1.IdRequest) (*energyMeterv1.BoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (EnergyMetersServer) InstalledCutOrInactiveEnergyMeter(ctx context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseGetEnergyMeters, error) {
	//TODO implement me
	panic("implement me")
}

func (EnergyMetersServer) RecentInstallationEnergyMeter(ctx context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseEnergyMeter, error) {
	//TODO implement me
	panic("implement me")
}
