package model

import "time"

type City struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	StateID     uint
	StateName   string
	CountryID   uint
	CountryName string
	Name        string
}

type State struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CountryID   uint
	CountryName string
	Name        string
	Cities      []City `gorm:"foreignKey:StateID"`
}

type Country struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	States    []State `gorm:"foreignKey:CountryID"`
	Cities    []City  `gorm:"foreignKey:CountryID"`
}

type GeoChecker struct {
	CountryName string
	StateName   string
}
