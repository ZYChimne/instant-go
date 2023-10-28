package database

import (
	"zychimne/instant/pkg/model"
)

func GetFeed(userID uint, offset int, limit int, instants *[]model.Feed) error {
	return PostgresDB.Where("user_id = ?", userID).Order("created_at desc").Offset(offset).Limit(limit).Find(&instants).Error
}
