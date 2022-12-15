package energyMeter

import (
	"context"
	energyMeterApiv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/services/rest_config"
)

var errorPath = "ERR internal/energyMeter/energyMeter.go"

type Server struct {
	energyMeterApiv1.EnergyMetersServicesServer
}

func (e Server) GetEnergyMeters(ctx context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseGetEnergyMeters, error) {
	//TODO implement me
	panic("implement me")
}

func (e Server) CreateEnergyMeters(ctx context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseEnergyMeter, error) {
	//TODO implement me
	panic("implement me")
}

func (e Server) UpdateEnergyMeters(ctx context.Context, meters *energyMeterApiv1.RequestUpdateEnergyMeters) (*interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (e Server) DeleteEnergyMeters(ctx context.Context, e2 *interface{}) (*interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (e Server) InstalledCutOrInactiveEnergyMeter(ctx context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseGetEnergyMeters, error) {
	//TODO implement me
	panic("implement me")
}

func (e Server) RecentInstallationEnergyMeter(ctx context.Context, meter *energyMeterApiv1.RequestEnergyMeter) (*energyMeterApiv1.ResponseEnergyMeter, error) {
	//TODO implement me
	panic("implement me")
}
