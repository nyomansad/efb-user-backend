package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"-"`
	Role      string         `json:"role"` // Pilot, CoPilot, Admin, MaintenanceCrew
	AirlineID uint           `json:"airline_id"`
	Airline   Airline        `gorm:"foreignKey:AirlineID" json:"airline"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
