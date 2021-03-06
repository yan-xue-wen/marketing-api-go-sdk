/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 预算优化
type OptBudgetStruct struct {
	RaiseDayBudgetFlag            bool  `json:"raise_day_budget_flag,omitempty"`
	CurrentAdgroupDayBudget       int64 `json:"current_adgroup_day_budget,omitempty"`
	CurrentCampaignDayBudget      int64 `json:"current_campaign_day_budget,omitempty"`
	CurrentAccountDayBudget       int64 `json:"current_account_day_budget,omitempty"`
	AdgroupCostDaily              int64 `json:"adgroup_cost_daily,omitempty"`
	CampaignCostDaily             int64 `json:"campaign_cost_daily,omitempty"`
	AccountCostDaily              int64 `json:"account_cost_daily,omitempty"`
	AdgroupBalance                int64 `json:"adgroup_balance,omitempty"`
	CampaignBalance               int64 `json:"campaign_balance,omitempty"`
	AccountBalance                int64 `json:"account_balance,omitempty"`
	OptimizeAdgroupDayBudgetFlag  bool  `json:"optimize_adgroup_day_budget_flag,omitempty"`
	OptimizeCampaignDayBudgetFlag bool  `json:"optimize_campaign_day_budget_flag,omitempty"`
	OptimizeAccountDayBudgetFlag  bool  `json:"optimize_account_day_budget_flag,omitempty"`
	OptimizeAdgroupDayBudget      int64 `json:"optimize_adgroup_day_budget,omitempty"`
	OptimizeCampaignDayBudget     int64 `json:"optimize_campaign_day_budget,omitempty"`
	OptimizeAccountDayBudget      int64 `json:"optimize_account_day_budget,omitempty"`
	RaiseAccountBalanceFlag       bool  `json:"raise_account_balance_flag,omitempty"`
	RecommendRecharge             int64 `json:"recommend_recharge,omitempty"`
}
