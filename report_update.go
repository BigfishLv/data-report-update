package main

import (
	"context"
	"data-report-update/config"
	"data-report-update/dependencies"
	"data-report-update/domain/csv"
	"data-report-update/domain/repositories"
	"data-report-update/logger"
	"data-report-update/reader"
)

func main() {
	params := config.NewParams()
	config := config.NewConfig(params)
	logger.InitLogger(config.Logger.Path, config.Logger.Level, config.Logger.Console)

	db := dependencies.NewMySQLClient(config)
	adPvClickCostDailyRepository := repositories.NewAdPvClickCostDailyRepository(db)
	campaignDataDailySummariesRepository := repositories.NewCampaignDataDailySummariesRepository(db)
	creativeDataDailySummariesRepository := repositories.NewCreativeDataDailySummariesRepository(db)
	userDataDailySummariesRepository := repositories.NewUserDataDailySummariesRepository(db)
	allUsersDataDailySummariesRepository := repositories.NewAllUsersDataDailySummariesRepository(db)

	adPvFileProcessor := reader.CsvFileProcessor[csv.AdPvClickCostDailyCsvData]{}
	filePath := config.Csv.Path + config.Csv.AdPvClickCostDailyCsvFileName
	adPvClickCostCsvDataArray, err := adPvFileProcessor.Read(filePath)
	if err != nil {
		logger.Error(context.Background(), "adPvFileProcessor read filePath : %s, error : %+v", filePath, err)
		return
	}
	err = adPvClickCostDailyRepository.UpdateAdPvClickCostDaily(adPvClickCostCsvDataArray)
	if err != nil {
		logger.Error(context.Background(), "adPvClickCostDailyRepository UpdateAdPvClickCostDaily error : %+v", err)
		return
	}

	campaignDataDailyFileProcessor := reader.CsvFileProcessor[csv.CampaignDataDailySummariesCsvData]{}
	filePath = config.Csv.Path + config.Csv.CampaignDataDailySummariesCsvFileName
	campaignDataDailyCsvDataArray, err := campaignDataDailyFileProcessor.Read(filePath)
	if err != nil {
		logger.Error(context.Background(), "campaignDataDailyFileProcessor read filePath : %s, error : %+v", filePath, err)
		return
	}
	err = campaignDataDailySummariesRepository.UpdateCampaignDataDailySummaries(campaignDataDailyCsvDataArray)
	if err != nil {
		logger.Error(context.Background(), "campaignDataDailySummariesRepository UpdateCampaignDataDailySummaries error : %+v", err)
		return
	}

	creativeDataDailyFileProcessor := reader.CsvFileProcessor[csv.CreativeDataDailySummariesCsvData]{}
	filePath = config.Csv.Path + config.Csv.CreativeDataDailySummariesCsvFileName
	creativeDataDailyCsvDataArray, err := creativeDataDailyFileProcessor.Read(filePath)
	if err != nil {
		logger.Error(context.Background(), "creativeDataDailyFileProcessor read filePath : %s, error : %+v", filePath, err)
		return
	}
	err = creativeDataDailySummariesRepository.UpdateCreativeDataDailySummaries(creativeDataDailyCsvDataArray)
	if err != nil {
		logger.Error(context.Background(), "creativeDataDailySummariesRepository UpdateCreativeDataDailySummaries error : %+v", err)
		return
	}

	userDataDailyFileProcessor := reader.CsvFileProcessor[csv.UserDataDailySummariesCsvData]{}
	filePath = config.Csv.Path + config.Csv.UserDataDailySummariesCsvFileName
	userDataDailyCsvDataArray, err := userDataDailyFileProcessor.Read(filePath)
	if err != nil {
		logger.Error(context.Background(), "userDataDailyFileProcessor read filePath : %s, error : %+v", filePath, err)
		return
	}
	err = userDataDailySummariesRepository.UpdateUserDataDailySummaries(userDataDailyCsvDataArray)
	if err != nil {
		logger.Error(context.Background(), "userDataDailySummariesRepository UpdateUserDataDailySummaries error : %+v", err)
		return
	}

	allUsersDataDailyFileProcessor := reader.CsvFileProcessor[csv.AllUsersDataDailySummariesCsvData]{}
	filePath = config.Csv.Path + config.Csv.AllUsersDataDailySummariesCsvFileName
	allUsersDataDailyCsvDataArray, err := allUsersDataDailyFileProcessor.Read(filePath)
	if err != nil {
		logger.Error(context.Background(), "allUsersDataDailyFileProcessor read filePath : %s, error : %+v", filePath, err)
		return
	}
	err = allUsersDataDailySummariesRepository.UpdateAllUsersDataDailySummaries(allUsersDataDailyCsvDataArray)
	if err != nil {
		logger.Error(context.Background(), "allUsersDataDailySummariesRepository UpdateAllUsersDataDailySummaries error : %+v", err)
		return
	}

	logger.Info(context.Background(), "data report update succeed.")
}
