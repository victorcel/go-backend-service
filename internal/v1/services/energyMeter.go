package services

import (
	"context"
	"errors"
	"fmt"
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

func (EnergyMetersServer) GetEnergyMeters(
	_ context.Context,
	_ *energyMeterApiv1.RequestEnergyMeter,
) (*energyMeterApiv1.ResponseGetEnergyMeters, error) {
	return useCases.Get()

}

func (EnergyMetersServer) CreateEnergyMeters(
	_ context.Context,
	meter *energyMeterApiv1.RequestEnergyMeter,
) (*energyMeterApiv1.ResponseEnergyMeter, error) {
	_, err := useCases.Insert(meter.EnergyMeter)

	if err != nil {
		err := status.Error(codes.NotFound, err.Error())
		return nil, err
	}

	response := &energyMeterApiv1.ResponseEnergyMeter{}

	response.EnergyMeter = meter.GetEnergyMeter()

	return response, nil
}

func (EnergyMetersServer) UpdateEnergyMeters(
	_ context.Context,
	meters *energyMeterApiv1.RequestUpdateEnergyMeters,
) (*energyMeterv1.BoolResponse, error) {
	return useCases.Update(meters.GetIdRequest(), meters.GetEnergyMeter())
}

func (EnergyMetersServer) DeleteEnergyMeters(
	_ context.Context,
	request *energyMeterv1.IdRequest,
) (*energyMeterv1.BoolResponse, error) {
	db, err := useCases.Delete(request.GetId())

	if err != nil {
		return &energyMeterv1.BoolResponse{Response: db}, errors.New(fmt.Sprintf("%s %s", errorPath, err))
	}
	return &energyMeterv1.BoolResponse{Response: db}, nil
}

func (EnergyMetersServer) InstalledCutOrInactiveEnergyMeter(
	_ context.Context,
	meter *energyMeterApiv1.RequestEnergyMeter,
) (*energyMeterApiv1.ResponseGetEnergyMeters, error) {
	find, err := useCases.InstalledCutOrInactiveEnergyMeter()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s %s", errorPath, err))
	}

	return find, err
}

func (EnergyMetersServer) RecentInstallationEnergyMeter(
	_ context.Context,
	meter *energyMeterApiv1.RequestEnergyMeter,
) (*energyMeterApiv1.ResponseEnergyMeter, error) {

	find, err := useCases.RecentInstallationEnergyMeter(meter)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s %s", errorPath, err))
	}

	return find, err
}
