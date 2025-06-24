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

type AdPvClickCostDailyRepository struct {
	db *gorm.DB
}

func NewAdPvClickCostDailyRepository(db *gorm.DB) *AdPvClickCostDailyRepository {
	return &AdPvClickCostDailyRepository{
		db: db,
	}
}

func (a *AdPvClickCostDailyRepository) UpdateAdPvClickCostDaily(csvDataArray []*csv.AdPvClickCostDailyCsvData) (err error) {
	// start transaction
	logger.Info(context.Background(), "UpdateAdPvClickCostDaily started.")
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", r)
			logger.Error(context.Background(), "UpdateAdPvClickCostDaily catch panic, error : %+v", err)
		}
	}()

	layout := "2006-01-02"
	insertRows := 0
	updateRows := 0
	for _, csvData := range csvDataArray {
		var adPvClickCostDaily model.AdPvClickCostDaily
		if err = tx.First(&adPvClickCostDaily, "happened_date = ? AND campaign_id = ? AND creative_id = ?", csvData.HappenedDate.Time.Format(layout), csvData.CampaignId, csvData.CreativeId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 传递指针才能用自增id
				adPvClickCostDaily = convertors.ConvertAdPvClickCostDaily(csvData)
				tx.Create(&adPvClickCostDaily)
				insertRows++
				continue
			}
			logger.Error(context.Background(), "UpdateAdPvClickCostDaily select ad_pv_click_cost_daily table error : %+v", err)
			tx.Rollback()
			return
		}
		adPvClickCostDaily.ViewCount += csvData.ViewCount
		adPvClickCostDaily.ClickCount += csvData.ClickCount
		tx.Save(&adPvClickCostDaily)
		updateRows++
	}
	tx.Commit()
	logger.Info(context.Background(), "UpdateAdPvClickCostDaily succeed, insertRows : %d, updateRows :%d", insertRows, updateRows)
	return
}
