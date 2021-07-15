package services

import (
	"github.com/LucasCarioca/pi-net-collector/pkg/models"
	"github.com/LucasCarioca/pi-net-collector/pkg/config"
	"gorm.io/gorm"
)

type ClimateService struct {
	db *gorm.DB
}

func NewClimateService() ClimateService {
	return ClimateService{
		db: config.GetDataSource(),
	}
}

func (s *ClimateService) CreateClimateRecord(temperature string, humidity string, node string, location string) models.Climate {
	record := &models.Climate{
		Temperature: temperature, 
		Humidity: humidity,
		Node: node,
		Location: location,
	}
	s.db.Create(record)
	return *record
}

func (s *ClimateService)  GetClimateRecords() []models.Climate {
	var climateRecords []models.Climate
	s.db.Find(&climateRecords)
	return climateRecords
}