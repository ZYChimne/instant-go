package database

import (
	"zychimne/instant/pkg/model"

	"gorm.io/gorm"
)

func LikeInstant(like *model.InstantLike) error {
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
	err := PostgresDB.Where("id = ?", like.UserID).First(&user).Error
	if err != nil {
		return err
	}
	like.Username = user.Username
	like.Nickname = user.Nickname
	like.Avatar = user.Avatar
	like.UserType = user.UserType
	err = tx.Create(&like).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Table("instants").Where("id = ?", like.InstantID).Update("likeCount", gorm.Expr("likeCount + ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Table("feeds").Where("user_id = ? and instant_id = ?", like.UserID, like.InstantID).Update("attitude", like.Attitude).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func UnlikeInstant(like *model.InstantLike) error {
	tx := PostgresDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	err := tx.Where("user_id = ? and instant_id = ?", like.UserID, like.InstantID).Delete(&like).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Table("instants").Where("id = ?", like.InstantID).Update("likeCount", gorm.Expr("likeCount - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Table("feeds").Where("user_id = ? and instant_id = ?", like.UserID, like.InstantID).Update("attitude", like.Attitude).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func GetLikes(instantID uint, offset int, limit int, likes *[]model.InstantLike) error {
	return PostgresDB.Where("instant_id = ?", instantID).Find(&likes).Order("created desc").Offset(offset).Limit(limit).Error
}
