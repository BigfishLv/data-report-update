
mysql -uroot -p -e "SELECT happened_date, user_id, campaign_id, creative_id, bidding_type, view_count, click_count, spent, should_spent, user_balance, version, created_at, updated_at FROM bp.ad_pv_click_cost_daily where happened_date >= '2025-06-18' and happened_date <= '2025-06-30'"  | sed 's/\t/,/g' > ad_pv_click_cost_daily.csv

mysql -uroot -p -e "SELECT campaign_id, happened_date, view_count, click_count, spent, should_spent, version, created_at, updated_at FROM bp.campaign_data_daily_summaries where happened_date >= '2025-06-18' and happened_date <= '2025-06-30'"  | sed 's/\t/,/g' > campaign_data_daily_summaries.csv

mysql -uroot -p -e "SELECT creative_id, happened_date, view_count, click_count, spent, should_spent, version, created_at, updated_at FROM bp.creative_data_daily_summaries where happened_date >= '2025-06-18' and happened_date <= '2025-06-30'"  | sed 's/\t/,/g' > creative_data_daily_summaries.csv

mysql -uroot -p -e "SELECT user_id, happened_date, view_count, click_count, spent, should_spent, bidding_type, balance, version, created_at, updated_at FROM bp.user_data_daily_summaries where happened_date >= '2025-06-18' and happened_date <= '2025-06-30'"  | sed 's/\t/,/g' > user_data_daily_summaries.csv

mysql -uroot -p -e "SELECT happened_date, view_count, click_count, spent, should_spent, bidding_type, balance, version, created_at, updated_at FROM bp.all_users_data_daily_summaries where happened_date >= '2025-06-18' and happened_date <= '2025-06-30'"  | sed 's/\t/,/g' > all_users_data_daily_summaries.csv



