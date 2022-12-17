package useCases

import (
	"github.com/victorcel/go-enertbit-backend-service/internal/v1/repository"
	energyMeterv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/energyMeter"
	energyMeterApiv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/services/rest_config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Insert(meters *energyMeterv1.EnergyMeters) (*energyMeterApiv1.ResponseEnergyMeter, error) {

	if meters.GetLines() == 0 || meters.GetLines() > 10 {
		err := status.Error(codes.NotFound, "Número de líneas conectadas, puede ir de 1 a 10")
		return nil, err
	}

	findSerial, err := repository.Find(&energyMeterv1.EnergyMeters{
		Serial: meters.GetSerial(),
	})

	if err != nil {
		return nil, err
	}

	for _, serial := range findSerial.GetEnergyMeter() {
		if serial.GetSerial() == meters.GetSerial() {
			err := status.Error(codes.NotFound, "El Serial ya existe")
			return nil, err
		}
	}

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

func Update(id *energyMeterv1.IdRequest, meter *energyMeterv1.EnergyMeters) (*energyMeterv1.BoolResponse, error) {
	return repository.Update(id, meter)
}

func Delete(id string) (bool, error) {
	find, err := repository.Find(&energyMeterv1.EnergyMeters{Id: id})

	if err != nil {
		return false, err
	}

	if len(find.GetEnergyMeter()) == 0 {
		err := status.Error(codes.NotFound, "Id no existe")
		return false, err
	}

	for _, energyMeter := range find.EnergyMeter {
		if energyMeter.IsActive == true {
			return false, status.Error(codes.NotFound, err.Error())
		}
	}

	responseDelete, err := repository.Delete(id)
	if err != nil {
		return responseDelete, err
	}

	return responseDelete, nil
}
