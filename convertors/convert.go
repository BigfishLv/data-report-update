package convertors

import (
	"data-report-update/domain/csv"
	"data-report-update/domain/model"
)

func ConvertAdPvClickCostDaily(csvData *csv.AdPvClickCostDailyCsvData) model.AdPvClickCostDaily {
	return model.AdPvClickCostDaily{
		HappenedDate: csvData.HappenedDate.Time,
		UserId:       csvData.UserId,
		CampaignId:   csvData.CampaignId,
		CreativeId:   csvData.CreativeId,
		BiddingType:  csvData.BiddingType,
		ViewCount:    csvData.ViewCount,
		ClickCount:   csvData.ClickCount,
		Spent:        csvData.Spent,
		ShouldSpent:  csvData.ShouldSpent,
		UserBalance:  csvData.UserBalance,
		Version:      csvData.Version,
		CreatedAt:    csvData.CreatedAt.Time,
		UpdatedAt:    csvData.UpdatedAt.Time,
	}
}

func ConvertAdPvClickCostDailyArray(csvDataArray []*csv.AdPvClickCostDailyCsvData) []model.AdPvClickCostDaily {
	adPvClickCostDailyArray := make([]model.AdPvClickCostDaily, 0)
	for _, csvData := range csvDataArray {
		adPvClickCostDailyArray = append(adPvClickCostDailyArray, ConvertAdPvClickCostDaily(csvData))
	}
	return adPvClickCostDailyArray
}

func ConvertCampaignDataDailySummaries(csvData *csv.CampaignDataDailySummariesCsvData) model.CampaignDataDailySummaries {
	return model.CampaignDataDailySummaries{
		HappenedDate: csvData.HappenedDate.Time,
		CampaignId:   csvData.CampaignId,
		ViewCount:    csvData.ViewCount,
		ClickCount:   csvData.ClickCount,
		Spent:        csvData.Spent,
		ShouldSpent:  csvData.ShouldSpent,
		Version:      csvData.Version,
		CreatedAt:    csvData.CreatedAt.Time,
		UpdatedAt:    csvData.UpdatedAt.Time,
	}
}

func ConvertCreativeDataDailySummaries(csvData *csv.CreativeDataDailySummariesCsvData) model.CreativeDataDailySummaries {
	return model.CreativeDataDailySummaries{
		HappenedDate: csvData.HappenedDate.Time,
		CreativeId:   csvData.CreativeId,
		ViewCount:    csvData.ViewCount,
		ClickCount:   csvData.ClickCount,
		Spent:        csvData.Spent,
		ShouldSpent:  csvData.ShouldSpent,
		Version:      csvData.Version,
		CreatedAt:    csvData.CreatedAt.Time,
		UpdatedAt:    csvData.UpdatedAt.Time,
	}
}

func ConvertUserDataDailySummaries(csvData *csv.UserDataDailySummariesCsvData) model.UserDataDailySummaries {
	return model.UserDataDailySummaries{
		HappenedDate: csvData.HappenedDate.Time,
		UserId:       csvData.UserId,
		ViewCount:    csvData.ViewCount,
		ClickCount:   csvData.ClickCount,
		Spent:        csvData.Spent,
		ShouldSpent:  csvData.ShouldSpent,
		BiddingType:  csvData.BiddingType,
		Balance:      csvData.Balance,
		Version:      csvData.Version,
		CreatedAt:    csvData.CreatedAt.Time,
		UpdatedAt:    csvData.UpdatedAt.Time,
	}
}

func ConvertAllUsersDataDailySummaries(csvData *csv.AllUsersDataDailySummariesCsvData) model.AllUsersDataDailySummaries {
	return model.AllUsersDataDailySummaries{
		HappenedDate: csvData.HappenedDate.Time,
		ViewCount:    csvData.ViewCount,
		ClickCount:   csvData.ClickCount,
		Spent:        csvData.Spent,
		ShouldSpent:  csvData.ShouldSpent,
		BiddingType:  csvData.BiddingType,
		Balance:      csvData.Balance,
		Version:      csvData.Version,
		CreatedAt:    csvData.CreatedAt.Time,
		UpdatedAt:    csvData.UpdatedAt.Time,
	}
}
