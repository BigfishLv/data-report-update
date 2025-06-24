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

type CampaignDataDailySummariesRepository struct {
	db *gorm.DB
}

func NewCampaignDataDailySummariesRepository(db *gorm.DB) *CampaignDataDailySummariesRepository {
	return &CampaignDataDailySummariesRepository{
		db: db,
	}
}

func (a *CampaignDataDailySummariesRepository) UpdateCampaignDataDailySummaries(csvDataArray []*csv.CampaignDataDailySummariesCsvData) (err error) {
	// start transaction
	logger.Info(context.Background(), "UpdateCampaignDataDailySummaries started.")
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", r)
			logger.Error(context.Background(), "UpdateCampaignDataDailySummaries catch panic, error : %+v", err)
		}
	}()

	layout := "2006-01-02"
	insertRows := 0
	updateRows := 0
	for _, csvData := range csvDataArray {
		var campaignDataDailySummaries model.CampaignDataDailySummaries
		if err = tx.First(&campaignDataDailySummaries, "happened_date = ? AND campaign_id = ?", csvData.HappenedDate.Time.Format(layout), csvData.CampaignId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 传递指针才能用自增id
				campaignDataDailySummaries = convertors.ConvertCampaignDataDailySummaries(csvData)
				tx.Create(&campaignDataDailySummaries)
				insertRows++
				continue
			}
			logger.Error(context.Background(), "UpdateCampaignDataDailySummaries select campaign_data_daily_summaries table error : %+v", err)
			tx.Rollback()
			return
		}
		campaignDataDailySummaries.ViewCount += csvData.ViewCount
		campaignDataDailySummaries.ClickCount += csvData.ClickCount
		tx.Save(&campaignDataDailySummaries)
		updateRows++
	}
	tx.Commit()
	logger.Info(context.Background(), "UpdateCampaignDataDailySummaries succeed, insertRows : %d, updateRows :%d", insertRows, updateRows)
	return
}
