package repositories

import (
	"context"
	"data-report-update/convertors"
	"data-report-update/domain/csv"
	"data-report-update/domain/model"
	"data-report-update/logger"
	"fmt"
	"gorm.io/gorm"
)

type UserDataDailySummariesRepository struct {
	db *gorm.DB
}

func NewUserDataDailySummariesRepository(db *gorm.DB) *UserDataDailySummariesRepository {
	return &UserDataDailySummariesRepository{
		db: db,
	}
}

func (a *UserDataDailySummariesRepository) UpdateUserDataDailySummaries(csvDataArray []*csv.UserDataDailySummariesCsvData) (err error) {
	// start transaction
	logger.Info(context.Background(), "UpdateUserDataDailySummaries started.")
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", r)
			logger.Error(context.Background(), "UpdateUserDataDailySummaries catch panic, error : %+v", err)
		}
	}()

	layout := "2006-01-02"
	insertRows := 0
	updateRows := 0
	for _, csvData := range csvDataArray {
		var userDataDailySummaries model.UserDataDailySummaries
		if err = tx.First(&userDataDailySummaries, "happened_date = ? AND user_id = ? AND bidding_type = ?", csvData.HappenedDate.Time.Format(layout), csvData.UserId, csvData.BiddingType).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 传递指针才能用自增id
				userDataDailySummaries = convertors.ConvertUserDataDailySummaries(csvData)
				tx.Create(&userDataDailySummaries)
				insertRows++
				continue
			}
			logger.Error(context.Background(), "UpdateUserDataDailySummaries select campaign_data_daily_summaries table error : %+v", err)
			tx.Rollback()
			return
		}
		userDataDailySummaries.ViewCount += csvData.ViewCount
		userDataDailySummaries.ClickCount += csvData.ClickCount
		tx.Save(&userDataDailySummaries)
		updateRows++
	}
	tx.Commit()
	logger.Info(context.Background(), "UpdateUserDataDailySummaries succeed, insertRows : %d, updateRows :%d", insertRows, updateRows)
	return
}
