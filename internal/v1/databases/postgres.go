package databases

import (
	"github.com/gofrs/uuid"
	energyMeterv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/energyMeter"
	energyMeterApiv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/services/rest_config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

type Postgres struct {
	db *gorm.DB
}

type EnergyMeters struct {
	ID               string `gorm:"not null"`
	Brand            string `gorm:"not null"`
	Address          string `gorm:"not null"`
	InstallationDate int64  `gorm:"not null"`
	RetirementDate   int64  `gorm:"default:null"`
	Serial           string `gorm:"not null"`
	Lines            int    `gorm:"not null"`
	IsActive         bool   `gorm:"not null"`
	CreatedAt        int64  `gorm:"not null"`
}

func NewPostgres(db *gorm.DB) (*Postgres, error) {

	err := db.AutoMigrate(&EnergyMeters{})

	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) Insert(meter *energyMeterv1.EnergyMeters) (*energyMeterApiv1.ResponseEnergyMeter, error) {

	newV4, _ := uuid.NewV4()

	parseInstallationDate, err := time.Parse("02/01/2006 15:04", meter.InstallationDate)

	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	create := p.db.Create(&EnergyMeters{
		ID:               newV4.String(),
		Brand:            meter.GetBrand(),
		Address:          meter.GetAddress(),
		InstallationDate: parseInstallationDate.Unix(),
		Serial:           meter.GetSerial(),
		Lines:            int(meter.GetLines()),
		IsActive:         meter.GetIsActive(),
		CreatedAt:        time.Now().Unix(),
	})

	if create.Error != nil {
		return nil, status.Error(codes.NotFound, create.Error.Error())
	}

	var response = &energyMeterApiv1.ResponseEnergyMeter{}

	response.EnergyMeter = meter

	response.EnergyMeter.Id = newV4.String()

	return response, nil

}

func (p *Postgres) Update(
	id *energyMeterv1.IdRequest, meter *energyMeterv1.EnergyMeters,
) (*energyMeterv1.BoolResponse, error) {

	if id.GetId() == "" {
		return &energyMeterv1.BoolResponse{Response: false}, status.Error(codes.NotFound, "No existe ID")
	}

	find := p.db.Where("id = ?", id.GetId())

	if find.Error != nil {
		return &energyMeterv1.BoolResponse{Response: false}, status.Error(codes.Internal, find.Error.Error())
	}

	if meter.GetRetirementDate() != "" {
		parseRetirementDateDate, error := time.Parse("02/01/2006 15:04", meter.GetRetirementDate())
		if error != nil {
			return &energyMeterv1.BoolResponse{Response: false}, status.Error(codes.NotFound, error.Error())
		}

		update := find.Update(id.GetId(), &EnergyMeters{
			Address:        meter.GetAddress(),
			RetirementDate: parseRetirementDateDate.Unix(),
			Lines:          int(meter.GetLines()),
			IsActive:       meter.GetIsActive(),
		})

		if update.Error != nil {
			return &energyMeterv1.BoolResponse{Response: false}, status.Error(codes.Internal, update.Error.Error())
		}

	}

	update := find.Updates(&EnergyMeters{
		Address:  meter.GetAddress(),
		Lines:    int(meter.GetLines()),
		IsActive: meter.GetIsActive(),
	})

	if update.Error != nil {
		return &energyMeterv1.BoolResponse{Response: false}, status.Error(codes.Internal, update.Error.Error())
	}

	return &energyMeterv1.BoolResponse{Response: true}, nil
}

func (p *Postgres) Get() (*energyMeterApiv1.ResponseGetEnergyMeters, error) {
	var response = &energyMeterApiv1.ResponseGetEnergyMeters{}

	find := p.db.Where(&EnergyMeters{}).Find(&response.EnergyMeter)

	if find.Error != nil {
		return nil, status.Error(codes.NotFound, find.Error.Error())
	}

	return response, nil
}

func (p *Postgres) Delete(id string) (bool, error) {

	db := p.db.Where(&energyMeterv1.EnergyMeters{}).Delete(&energyMeterv1.EnergyMeters{Id: id})

	if db.Error != nil {
		return false, status.Error(codes.NotFound, db.Error.Error())
	}

	return true, nil
}

func (p *Postgres) Find(meter *energyMeterv1.EnergyMeters) (*energyMeterApiv1.ResponseGetEnergyMeters, error) {

	var response = &energyMeterApiv1.ResponseGetEnergyMeters{}

	find := p.db.Where(meter).Find(&response.EnergyMeter)

	if find.Error != nil {
		return nil, status.Error(codes.NotFound, find.Error.Error())
	}

	return response, nil
}

func (p *Postgres) InstalledCutOrInactive() (*energyMeterApiv1.ResponseGetEnergyMeters, error) {
	var response = &energyMeterApiv1.ResponseGetEnergyMeters{}

	find := p.db.
		Model(&EnergyMeters{}).
		Where("installation_date > 0  and is_active  = false").
		Find(&response.EnergyMeter)

	if find.Error != nil {
		return nil, status.Error(codes.NotFound, find.Error.Error())
	}

	return response, nil
}

func (p *Postgres) RecentInstallationEnergyMeter(
	meter *energyMeterApiv1.RequestEnergyMeter,
) (*energyMeterApiv1.ResponseEnergyMeter, error) {

	var response = &energyMeterApiv1.ResponseEnergyMeter{}

	find := p.db.
		Model(&EnergyMeters{}).
		Where(&EnergyMeters{Serial: meter.GetEnergyMeter().GetSerial(), Brand: meter.GetEnergyMeter().GetBrand()}).
		Order("installation_date desc").
		Limit(1).
		Find(&response.EnergyMeter)

	if find.Error != nil {
		return nil, status.Error(codes.NotFound, find.Error.Error())
	}

	return response, nil
}
