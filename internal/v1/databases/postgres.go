package databases

import (
	energyMeterv1 "github.com/victorcel/proto-enertbit-grpc-models/pkg/v1/energyMeter"
	"gorm.io/gorm"
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

func (p *Postgres) Insert(meters *energyMeterv1.EnergyMeters) *gorm.DB {

	return p.db.Create(&meters)

}

func (p *Postgres) Update() {

}

func (p *Postgres) Get() {

}

func (p *Postgres) Delete() {

}

func (p *Postgres) InstalledCutOrInactive() {

}

func (p *Postgres) RecentInstallationEnergyMeter() {

}
