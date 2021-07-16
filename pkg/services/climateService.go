package services

import (
	"fmt"
	"github.com/LucasCarioca/pi-net-collector/pkg/config"
	"github.com/LucasCarioca/pi-net-collector/pkg/models"
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
		Humidity:    humidity,
		Node:        node,
		Location:    location,
	}
	s.db.Create(record)
	return *record
}

func (s *ClimateService) GetClimateRecords() []models.Climate {
	var climateRecords []models.Climate
	s.db.Find(&climateRecords)
	return climateRecords
}

func (s *ClimateService) GetClimateRecordsBy(field string, value string) []models.Climate {
	var climateRecords []models.Climate
	s.db.Find(&climateRecords, fmt.Sprintf("%s = ?", field), value)
	return climateRecords
}

func (s *ClimateService) GetLastClimateRecordBy(field string, value string) models.Climate {
	var climateRecord models.Climate
	s.db.Last(&climateRecord, fmt.Sprintf("%s = ?", field), value)
	return climateRecord
}

func (s *ClimateService) GetClimateRecordById(id int) models.Climate {
	var climateRecord models.Climate
	s.db.First(&climateRecord, id)
	return climateRecord
}

func (s *ClimateService) DeleteClimateRecord(id int) {
	s.db.Delete(&models.Climate{}, id)
}
