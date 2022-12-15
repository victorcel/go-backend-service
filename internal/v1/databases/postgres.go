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

func NewPostgres(db *gorm.DB) (*Postgres, error) {

	err := db.AutoMigrate(&energyMeterv1.EnergyMeters{})

	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) Insert(meters *energyMeterv1.EnergyMeters) (*energyMeterApiv1.ResponseEnergyMeter, error) {
	uuid, _ := uuid.NewV4()
	create := p.db.Create(&energyMeterv1.EnergyMeters{
		Id:               uuid.String(),
		Brand:            meters.Brand,
		Address:          meters.Address,
		InstallationDate: meters.InstallationDate,
		Serial:           meters.Serial,
		Lines:            meters.Lines,
		IsActive:         meters.IsActive,
		CreatedAt:        time.Now().String(),
	})

	if create.Error != nil {
		return nil, status.Error(codes.NotFound, create.Error.Error())
	}

	var response = &energyMeterApiv1.ResponseEnergyMeter{}

	response.EnergyMeter = meters

	return response, nil

}

func (p *Postgres) Update() {

}

func (p *Postgres) Get() {

}

func (p *Postgres) Delete() {

}

func (p *Postgres) Find(meters *energyMeterv1.EnergyMeters) (*energyMeterApiv1.ResponseGetEnergyMeters, error) {

	var response = &energyMeterApiv1.ResponseGetEnergyMeters{}
	var count int64
	find := p.db.Where(meters).Find(&response.EnergyMeter).Count(&count)

	if find.Error != nil {
		return nil, status.Error(codes.NotFound, find.Error.Error())
	}

	return response, nil
}

func (p *Postgres) InstalledCutOrInactive() {

}

func (p *Postgres) RecentInstallationEnergyMeter() {

}
