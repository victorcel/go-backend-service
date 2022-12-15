package repository

import (
	energyMeterv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/energyMeter"
	"gorm.io/gorm"
)

type EnergyMeterRepository interface {
	Insert(meters *energyMeterv1.EnergyMeters) *gorm.DB
}

var implementationEnergyMeter EnergyMeterRepository

func SetEnergyMeterRepository(repository EnergyMeterRepository) {
	implementationEnergyMeter = repository
}

func Insert(meters *energyMeterv1.EnergyMeters) *gorm.DB {
	return implementationEnergyMeter.Insert(meters)
}
