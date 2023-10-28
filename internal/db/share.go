package database

import (
	"zychimne/instant/pkg/model"
)

func GetShares(instantID uint, offset int, limit int, instants *[]model.Instant) error {
	return PostgresDB.Where("ref_id = ?", instantID).Find(&instants).Error
}

func ShareInstant(instant *model.Instant) error {
	return nil
}
