package models

import "time"

type Airline struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Code      string    `gorm:"unique" json:"code"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}
