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

type AllUsersDataDailySummariesRepository struct {
	db *gorm.DB
}

func NewAllUsersDataDailySummariesRepository(db *gorm.DB) *AllUsersDataDailySummariesRepository {
	return &AllUsersDataDailySummariesRepository{
		db: db,
	}
}

func (a *AllUsersDataDailySummariesRepository) UpdateAllUsersDataDailySummaries(csvDataArray []*csv.AllUsersDataDailySummariesCsvData) (err error) {
	// start transaction
	logger.Info(context.Background(), "UpdateAllUsersDataDailySummaries started.")
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", r)
			logger.Error(context.Background(), "UpdateAllUsersDataDailySummaries catch panic, error : %+v", err)
		}
	}()

	layout := "2006-01-02"
	insertRows := 0
	updateRows := 0
	for _, csvData := range csvDataArray {
		var allUsersDataDailySummaries model.AllUsersDataDailySummaries
		if err = tx.First(&allUsersDataDailySummaries, "happened_date = ? AND bidding_type = ?", csvData.HappenedDate.Time.Format(layout), csvData.BiddingType).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 传递指针才能用自增id
				allUsersDataDailySummaries = convertors.ConvertAllUsersDataDailySummaries(csvData)
				tx.Create(&allUsersDataDailySummaries)
				insertRows++
				continue
			}
			logger.Error(context.Background(), "UpdateAllUsersDataDailySummaries select campaign_data_daily_summaries table error : %+v", err)
			tx.Rollback()
			return
		}
		allUsersDataDailySummaries.ViewCount += csvData.ViewCount
		allUsersDataDailySummaries.ClickCount += csvData.ClickCount
		tx.Save(&allUsersDataDailySummaries)
		updateRows++
	}
	tx.Commit()
	logger.Info(context.Background(), "UpdateAllUsersDataDailySummaries succeed, insertRows : %d, updateRows :%d", insertRows, updateRows)
	return
}
