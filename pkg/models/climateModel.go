package models

import (
	"gorm.io/gorm"
)

type Climate struct {
	gorm.Model
	Temperature string `json:"temperature" binding:"required"`
	Humidity    string `json:"humidity" binding:"required"`
	Node        string `json:"node" binding:"required"`
	Location    string `json:"location" binding:"required"`
}
