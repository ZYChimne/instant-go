package database

import (
	"zychimne/instant/config"
	"zychimne/instant/pkg/model"

	"gorm.io/gorm"
)

func GetInstants(userID uint, targetID uint, offset int, limit int, instants *[]model.Instant) error {
	return PostgresDB.Where("user_id = ?", targetID).Offset(offset).Limit(limit).Find(&instants).Error
}

func AddInstant(instant *model.Instant) error {
	tx := PostgresDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	var user model.User
	err := PostgresDB.Where("id = ?", instant.UserID).First(&user).Error
	if err != nil {
		return err
	}
	instant.Username = user.Username
	instant.Nickname = user.Nickname
	instant.Avatar = user.Avatar
	instant.InstantType = user.UserType
	err = tx.Create(&instant).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var followings []model.Following
	err = PostgresDB.Where("following_id = ?", instant.UserID).Find(&followings).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, following := range followings {
		err = fanOutOnWrite(tx, instant.ID, following.UserID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = fanOutOnWrite(tx, instant.ID, instant.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func fanOutOnWrite(tx *gorm.DB, instantID uint, userID uint) error {
	var count int64
	err := tx.Model(&model.Feed{}).Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return err
	}
	if count >= int64(config.Conf.Instant.MaxFeed)-1 {
		err = tx.Where("user_id = ?", userID).Order("created").Limit(1).Delete(&model.Feed{}).Error
		if err != nil {
			return err
		}
	}
	return tx.Create(&model.Feed{
		UserID:    userID,
		InstantID: instantID,
	}).Error
}

func UpdateInstant(instant *model.Instant) error {
	return PostgresDB.Model(&instant).Where("id = ?", instant.ID).Updates(&instant).Error
}

func DeleteInstant(instant *model.Instant) error {
	return PostgresDB.Where("user_id = ?", instant.UserID).Delete(&instant).Error
}