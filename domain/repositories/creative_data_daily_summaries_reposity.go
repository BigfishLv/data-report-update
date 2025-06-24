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

type CreativeDataDailySummariesRepository struct {
	db *gorm.DB
}

func NewCreativeDataDailySummariesRepository(db *gorm.DB) *CreativeDataDailySummariesRepository {
	return &CreativeDataDailySummariesRepository{
		db: db,
	}
}

func (a *CreativeDataDailySummariesRepository) UpdateCreativeDataDailySummaries(csvDataArray []*csv.CreativeDataDailySummariesCsvData) (err error) {
	// start transaction
	logger.Info(context.Background(), "UpdateCreativeDataDailySummaries started.")
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", r)
			logger.Error(context.Background(), "UpdateCreativeDataDailySummaries catch panic, error : %+v", err)
		}
	}()

	layout := "2006-01-02"
	insertRows := 0
	updateRows := 0
	for _, csvData := range csvDataArray {
		var creativeDataDailySummaries model.CreativeDataDailySummaries
		if err = tx.First(&creativeDataDailySummaries, "happened_date = ? AND creative_id = ?", csvData.HappenedDate.Time.Format(layout), csvData.CreativeId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 传递指针才能用自增id
				creativeDataDailySummaries = convertors.ConvertCreativeDataDailySummaries(csvData)
				tx.Create(&creativeDataDailySummaries)
				insertRows++
				continue
			}
			logger.Error(context.Background(), "UpdateCreativeDataDailySummaries select campaign_data_daily_summaries table error : %+v", err)
			tx.Rollback()
			return
		}
		creativeDataDailySummaries.ViewCount += csvData.ViewCount
		creativeDataDailySummaries.ClickCount += csvData.ClickCount
		tx.Save(&creativeDataDailySummaries)
		updateRows++
	}
	tx.Commit()
	logger.Info(context.Background(), "UpdateCreativeDataDailySummaries succeed, insertRows : %d, updateRows :%d", insertRows, updateRows)
	return
}
