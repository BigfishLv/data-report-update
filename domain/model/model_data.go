package model

import (
	"time"
)

type AdPvClickCostDaily struct {
	Id           uint64    `gorm:"primarykey;autoIncrement" json:"id"`
	HappenedDate time.Time `gorm:"type:date;uniqueIndex:uniq_row" json:"happened_date"`
	UserId       int64     `json:"user_id"`
	CampaignId   int64     `gorm:"uniqueIndex:uniq_row" json:"campaign_id"`
	CreativeId   int64     `gorm:"uniqueIndex:uniq_row" json:"creative_id"`
	BiddingType  int32     `json:"bidding_type"`
	ViewCount    int64     `json:"view_count"`
	ClickCount   int64     `json:"click_count"`
	Spent        int64     `json:"spent"`
	ShouldSpent  int64     `json:"should_spent"`
	UserBalance  int64     `json:"user_balance"`
	Version      int64     `json:"version"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (AdPvClickCostDaily) TableName() string {
	return "ad_pv_click_cost_daily" // 返回实际表名
}

type CampaignDataDailySummaries struct {
	Id           uint64    `gorm:"primarykey;autoIncrement" json:"id"`
	HappenedDate time.Time `gorm:"type:date;uniqueIndex:uniq_campaign_id_happened_date" json:"happened_date"`
	CampaignId   int64     `gorm:"uniqueIndex:uniq_campaign_id_happened_date" json:"campaign_id"`
	ViewCount    int64     `json:"view_count"`
	ClickCount   int64     `json:"click_count"`
	Spent        int64     `json:"spent"`
	ShouldSpent  int64     `json:"should_spent"`
	Version      int64     `json:"version"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (CampaignDataDailySummaries) TableName() string {
	return "campaign_data_daily_summaries" // 返回实际表名
}

type CreativeDataDailySummaries struct {
	Id           uint64    `gorm:"primarykey;autoIncrement" json:"id"`
	HappenedDate time.Time `gorm:"type:date;uniqueIndex:uniq_creative_id_happened_date" json:"happened_date"`
	CreativeId   int64     `gorm:"uniqueIndex:uniq_creative_id_happened_date" json:"campaign_id"`
	ViewCount    int64     `json:"view_count"`
	ClickCount   int64     `json:"click_count"`
	Spent        int64     `json:"spent"`
	ShouldSpent  int64     `json:"should_spent"`
	Version      int64     `json:"version"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (CreativeDataDailySummaries) TableName() string {
	return "creative_data_daily_summaries" // 返回实际表名
}

type UserDataDailySummaries struct {
	Id           uint64    `gorm:"primarykey;autoIncrement" json:"id"`
	HappenedDate time.Time `gorm:"type:date;uniqueIndex:uniq_user_id_happened_date_bidding_type" json:"happened_date"`
	UserId       int64     `gorm:"uniqueIndex:uniq_user_id_happened_date_bidding_type" json:"user_id"`
	ViewCount    int64     `json:"view_count"`
	ClickCount   int64     `json:"click_count"`
	Spent        int64     `json:"spent"`
	ShouldSpent  int64     `json:"should_spent"`
	BiddingType  int32     `gorm:"uniqueIndex:uniq_user_id_happened_date_bidding_type" json:"bidding_type"`
	Balance      int64     `json:"balance"`
	Version      int64     `json:"version"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (UserDataDailySummaries) TableName() string {
	return "user_data_daily_summaries" // 返回实际表名
}

type AllUsersDataDailySummaries struct {
	Id           uint64    `gorm:"primarykey;autoIncrement" json:"id"`
	HappenedDate time.Time `gorm:"type:date;uniqueIndex:uniq_happened_date_bidding_type" json:"happened_date"`
	ViewCount    int64     `json:"view_count"`
	ClickCount   int64     `json:"click_count"`
	Spent        int64     `json:"spent"`
	ShouldSpent  int64     `json:"should_spent"`
	BiddingType  int32     `gorm:"uniqueIndex:uniq_happened_date_bidding_type" json:"bidding_type"`
	Balance      int64     `json:"balance"`
	Version      int64     `json:"version"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (AllUsersDataDailySummaries) TableName() string {
	return "all_users_data_daily_summaries" // 返回实际表名
}

type CampaignDataSummaries struct {
	Id          uint64    `gorm:"primarykey;autoIncrement" json:"id"`
	CampaignId  int64     `gorm:"uniqueIndex:uniq_campaign_id" json:"campaign_id"`
	ViewCount   int64     `json:"view_count"`
	ClickCount  int64     `json:"click_count"`
	Spent       int64     `json:"spent"`
	ShouldSpent int64     `json:"should_spent"`
	Version     int64     `json:"version"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (CampaignDataSummaries) TableName() string {
	return "campaign_data_summaries" // 返回实际表名
}

type CreativeDataSummaries struct {
	Id          uint64    `gorm:"primarykey;autoIncrement" json:"id"`
	CreativeId  int64     `gorm:"uniqueIndex:uniq_creative_id" json:"creative_id"`
	ViewCount   int64     `json:"view_count"`
	ClickCount  int64     `json:"click_count"`
	Spent       int64     `json:"spent"`
	ShouldSpent int64     `json:"should_spent"`
	Version     int64     `json:"version"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (CreativeDataSummaries) TableName() string {
	return "creative_data_summaries" // 返回实际表名
}
