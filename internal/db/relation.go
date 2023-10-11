package database

import (
	"errors"
	"zychimne/instant/pkg/model"

	"gorm.io/gorm"
)

func GetFollowings(userID uint, offset int, limit int, followings *[]model.Following) error {
	return PostgresDB.Where("user_id = ?", userID).Order("id desc").Limit(limit).Offset(offset).Find(&followings).Error
}

func GetJointFollowings(userID uint, offset int, limit int, jointFollowing *[]model.JointFollowing) error {
	return PostgresDB.Table("followings").Select("followings.*, users.username, users.nickname, users.avatar").Joins("left join users on followings.target_id = users.id").Where("followings.user_id = ?", userID).Order("id desc").Limit(limit).Offset(offset).Scan(&jointFollowing).Error
}

func GetFollowers(userID uint, offset int, limit int, followers *[]model.Following) error {
	return PostgresDB.Where("target_id = ?", userID).Order("id desc").Limit(limit).Offset(int(offset)).Find(&followers).Error
}

func GetJointFollowers(userID uint, offset int, limit int, followerResponses *[]model.JointFollowing) error {
	return PostgresDB.Table("following").Select("followings.*, users.username, users.nickname, users.avatar").Joins("left join users on followings.user_id = users.id").Where("followings.target_id = ?", userID).Order("id desc").Limit(limit).Offset(offset).Scan(&followerResponses).Error
}

// Find people who follow me yet I don't follow them
func GetPotentialFollowings(userID uint, offset int, limit int, users *[]model.User) error {
	return nil
}

func AddFollowing(following *model.Following) error {
	if following.UserID == following.TargetID {
		return errors.New("userID and followingID must not be the same")
	}
	tx := PostgresDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	res := tx.Table("followings").Where("user_id = ?", following.TargetID).Where("target_id = ?", following.UserID).Update("is_friend", true) // save a query by using update
	if res.Error == nil {
		if res.RowsAffected == 1 {
			following.IsFriend = true
		}
	} else if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return res.Error
	}
	var targetUser model.User
	if err := tx.Select("type").Where("id = ?", following.TargetID).First(&targetUser).Error; err != nil {
		tx.Rollback()
		return err
	}
	following.TargetType = targetUser.Type
	if err := tx.Create(&following).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Table("users").Where("id = ?", following.UserID).Update("following_count", gorm.Expr("following_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Table("users").Where("id = ?", following.TargetID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func RemoveFollowing(following *model.Following) error {
	if following.UserID == following.TargetID {
		return errors.New("userID and followingID must not be the same")
	}
	tx := PostgresDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Table("followings").Where("user_id = ?", following.TargetID).Where("target_id = ?", following.UserID).Update("is_friend", false).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}
	if err := tx.Where("user_id = ?", following.UserID).Where("target_id = ?", following.TargetID).Unscoped().Delete(&following).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Table("users").Where("id = ?", following.UserID).Update("following_count", gorm.Expr("following_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Table("users").Where("id = ?", following.TargetID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func GetFriends(userID uint, offset int, limit int, friends *[]model.Following) error {
	return PostgresDB.Where("user_id = ?", userID).Where("is_friend = ?", true).Order("id desc").Limit(limit).Offset(offset).Find(&friends).Error
}

func GetJointFriends(userID uint, offset int, limit int, friends *[]model.JointFollowing) error {
	return PostgresDB.Table("followings").Select("followings.*, users.username, users.nickname, users.avatar").Joins("left join users on followings.target_id = users.id").Where("followings.user_id = ?", userID).Where("followings.is_friend = ?", true).Order("id desc").Limit(limit).Offset(offset).Scan(&friends).Error
}
