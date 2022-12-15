package useCases

import (
	"github.com/victorcel/go-enertbit-backend-service/internal/v1/repository"
	energyMeterv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/energyMeter"
	energyMeterApiv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/services/rest_config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Insert(meters *energyMeterv1.EnergyMeters) (*energyMeterApiv1.ResponseEnergyMeter, error) {

	find, err := repository.Find(&energyMeterv1.EnergyMeters{
		Serial: meters.GetSerial(),
		Brand:  meters.GetBrand(),
	})

	if err != nil {
		return nil, err
	}

	if len(find.GetEnergyMeter()) > 0 {
		err := status.Error(codes.NotFound, "Ya existe serial y marca")
		return nil, err
	}

	insert, err := repository.Insert(meters)
	if err != nil {
		err := status.Error(codes.NotFound, err.Error())
		return nil, err
	}

	return insert, nil

}
