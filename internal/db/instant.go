package database

import (
	"errors"
	"reflect"
	"time"
	"zychimne/instant/config"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"

	"gorm.io/gorm"
)

func GetInboxInstants(userID uint, offset int, limit int, instants *[]model.InboxInstant) error {
	return PostgresDB.Where("user_id = ?", userID).Order("created_at desc").Offset(offset).Limit(limit).Find(&instants).Error
}

func GetInstant(userID uint, instantID uint, instant *model.Instant) error {
	return PostgresDB.Where("id = ?", instantID).First(&instant).Error
}

func GetInstants(userID uint, targetID uint, offset int, limit int, instants *[]model.Instant) error {
	return PostgresDB.Where("user_id = ?", targetID).Offset(offset).Limit(limit).Find(&instants).Error
}

func AddInstant(instant *schema.UpsertInstantRequest, userID uint) error {
	tx := PostgresDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	now := time.Now()
	err := tx.Raw("INSERT INTO instants (created_at, updated_at, instant_type, content, visibility, ref_origin_id, user_id, username, nickname, avatar, user_type) SELECT ?, ?, ?, ?, ?, u.username, u.nickname, u.avatar, u.user_type FROM users u WHERE u.id = ? RETURNING id", now, now, instant.InstantType, instant.Content, instant.Visibility, instant.RefOriginID, userID, userID).Scan(&instant.ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var followerIDs []uint
	if instant.Visibility <= 1 {
		err = PostgresDB.Table("followings").Select("id").Where("target_id = ?", userID).Scan(&followerIDs).Error
	} else if instant.Visibility == 2 {
		err = PostgresDB.Table("followings").Select("id").Where("target_id = ? and is_friend = true", userID).Scan(&followerIDs).Error
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	followerIDs = append(followerIDs, userID)
	if err := upsertIntoFeed(tx, followerIDs, instant.ID, userID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func upsertIntoFeed(tx *gorm.DB, followerIDs []uint, instantID uint, userID uint) error {
	feeds := make([]model.InboxInstant, len(followerIDs))
	for i, followerID := range followerIDs {
		feeds[i] = model.InboxInstant{
			InstantID: instantID,
			UserID:    followerID,
		}
	}
	if err := tx.CreateInBatches(&feeds, config.Conf.Database.App.CreateInstantBatchSize).Error; err != nil {
		return err
	}
	return fanOutOnWrite(tx, followerIDs)
}

func fanOutOnWrite(tx *gorm.DB, userIDs []uint) error {
	for _, userID := range userIDs {
		if err := tx.Exec("DELETE FROM feeds WHERE id IN (SELECT id FROM feeds WHERE user_id = ? AND (SELECT COUNT(*) FROM feeds WHERE user_id = ?) > ? ORDER BY created_at ASC LIMIT 1)", userID, userID, config.Conf.Database.App.MaxFeed-1).Error; err != nil {
			return err
		}
	}
	return nil
}

func UpdateInstant(instant *model.UpsertInstant) error {
	var originalInstant model.UpsertInstant
	if err := PostgresDB.Table("instants").Where("id = ?", instant.ID).Scan(&originalInstant).Error; err != nil {
		return err
	}
	if originalInstant.UserID != instant.UserID {
		return errors.New(UserIDDoesNotMatchError)
	}
	if reflect.DeepEqual(originalInstant, instant) {
		return errors.New(NothingToUpdateError)
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
	if originalInstant.Visibility != instant.Visibility {
		if originalInstant.Visibility <= 1 {
			if instant.Visibility == 2 {
				if err := tx.Exec("DELETE FROM feeds WHERE instant_id = ? and user_id in (SELECT user_id FROM following WHERE target_id = ? and is_friend = false)", originalInstant.ID, originalInstant.UserID).Error; err != nil {
					tx.Rollback()
					return err
				}
			} else if instant.Visibility == 3 {
				if err := tx.Exec("DELETE FROM feeds WHERE instant_id = ? and user_id in (SELECT user_id FROM following WHERE target_id = ?)", originalInstant.ID, originalInstant.UserID).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		} else if originalInstant.Visibility == 2 {
			if instant.Visibility <= 1 {
				var followerIDs []uint
				if err := tx.Table("followings").Select("id").Where("target_id = ? and is_friend = ?", originalInstant.UserID, false).Scan(&followerIDs).Error; err != nil {
					tx.Rollback()
					return err
				}
				if err := upsertIntoFeed(tx, followerIDs, originalInstant.ID, originalInstant.UserID); err != nil {
					tx.Rollback()
					return err
				}
			} else if instant.Visibility == 3 {
				if err := tx.Exec("DELETE FROM feeds WHERE instant_id = ? and user_id in (SELECT user_id FROM following WHERE target_id = ? and is_friend = true)", originalInstant.ID, originalInstant.UserID).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		} else if originalInstant.Visibility == 3 {
			if instant.Visibility <= 1 {
				var followerIDs []uint
				if err := tx.Table("followings").Select("id").Where("target_id = ?", originalInstant.UserID).Scan(&followerIDs).Error; err != nil {
					tx.Rollback()
					return err
				}
				if err := upsertIntoFeed(tx, followerIDs, originalInstant.ID, originalInstant.UserID); err != nil {
					tx.Rollback()
					return err
				}
			} else if instant.Visibility == 2 {
				var followerIDs []uint
				if err := tx.Table("followings").Select("id").Where("target_id = ? and is_friend = true", originalInstant.UserID).Scan(&followerIDs).Error; err != nil {
					tx.Rollback()
					return err
				}
				if err := upsertIntoFeed(tx, followerIDs, originalInstant.ID, originalInstant.UserID); err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}
	if err := tx.Model(&instant).Where("id = ?", instant.ID).Updates(&instant).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func DeleteInstant(instant *model.Instant) error {
	tx := PostgresDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	err := tx.Where("user_id = ?", instant.UserID).Delete(&instant).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Where("instant_id = ?", instant.ID).Delete(&model.InstantLike{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Where("instant_id = ?", instant.ID).Delete(&model.Comment{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Where("instant_id = ?", instant.ID).Delete(&model.InboxInstant{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
