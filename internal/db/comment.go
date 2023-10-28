package database

import (
	"zychimne/instant/pkg/model"

	"gorm.io/gorm"
)

func GetComments(instantID uint, offset int, limit int, comments *[]model.Comment) error {
	return PostgresDB.Table("comments").Where("instant_id = ?", instantID).Order("created desc").Offset(offset).Limit(limit).Scan(&comments).Error
}

func AddComment(comment *model.Comment) error {
	var user model.BasicUser
	user.UserID = comment.UserID
	if err := GetBasicAccount(&user); err != nil {
		return err
	}
	comment.Username = user.Username
	comment.Nickname = user.Nickname
	comment.Avatar = user.Avatar
	comment.UserType = user.UserType
	tx := PostgresDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	err := tx.Create(&comment).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Table("instants").Where("id = ?", comment.InstantID).Update("commentCount", gorm.Expr("commentCount + ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func DeleteComment(comment *model.Comment) error {
	tx := PostgresDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	err := tx.Where("id = ?", comment.ID).Where("user_id = ?", comment.UserID).Limit(1).Delete(&comment).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Table("instants").Where("id = ?", comment.InstantID).Update("commentCount", gorm.Expr("commentCount - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
