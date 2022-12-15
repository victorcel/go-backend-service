package repository

import (
	energyMeterv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/energyMeter"
	energyMeterApiv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/services/rest_config"
)

type EnergyMeterRepository interface {
	Insert(meters *energyMeterv1.EnergyMeters) (*energyMeterApiv1.ResponseEnergyMeter, error)
	Find(meters *energyMeterv1.EnergyMeters) (*energyMeterApiv1.ResponseGetEnergyMeters, error)
}

var implementationEnergyMeter EnergyMeterRepository

func SetEnergyMeterRepository(repository EnergyMeterRepository) {
	implementationEnergyMeter = repository
}

func Insert(meters *energyMeterv1.EnergyMeters) (*energyMeterApiv1.ResponseEnergyMeter, error) {
	return implementationEnergyMeter.Insert(meters)
}

func Find(meters *energyMeterv1.EnergyMeters) (*energyMeterApiv1.ResponseGetEnergyMeters, error) {
	return implementationEnergyMeter.Find(meters)
}
