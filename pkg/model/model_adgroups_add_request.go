/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdgroupsAddRequest struct {
	CampaignId                 int64                           `json:"campaign_id,omitempty"`
	AdgroupName                string                          `json:"adgroup_name,omitempty"`
	PromotedObjectType         PromotedObjectType              `json:"promoted_object_type,omitempty"`
	BeginDate                  string                          `json:"begin_date,omitempty"`
	EndDate                    string                          `json:"end_date,omitempty"`
	BillingEvent               BillingEvent                    `json:"billing_event,omitempty"`
	BidAmount                  int64                           `json:"bid_amount,omitempty"`
	OptimizationGoal           OptimizationGoal                `json:"optimization_goal,omitempty"`
	TimeSeries                 string                          `json:"time_series,omitempty"`
	AutomaticSiteEnabled       bool                            `json:"automatic_site_enabled,omitempty"`
	SiteSet                    []string                        `json:"site_set,omitempty"`
	DailyBudget                int64                           `json:"daily_budget,omitempty"`
	PromotedObjectId           string                          `json:"promoted_object_id,omitempty"`
	AppAndroidChannelPackageId string                          `json:"app_android_channel_package_id,omitempty"`
	TargetingId                int64                           `json:"targeting_id,omitempty"`
	Targeting                  WriteTargetingSettingForAdgroup `json:"targeting,omitempty"`
	SceneSpec                  SceneTargetingForWrite          `json:"scene_spec,omitempty"`
	ConfiguredStatus           AdStatus                        `json:"configured_status,omitempty"`
	CustomizedCategory         string                          `json:"customized_category,omitempty"`
	DynamicAdSpec              DynamicAdSpec                   `json:"dynamic_ad_spec,omitempty"`
	UserActionSets             []UserActionSetStruct           `json:"user_action_sets,omitempty"`
	AdditionalUserActionSets   []UserActionSetStruct           `json:"additional_user_action_sets,omitempty"`
	DynamicCreativeId          int64                           `json:"dynamic_creative_id,omitempty"`
	IsRewardedVideoAd          bool                            `json:"is_rewarded_video_ad,omitempty"`
	BidStrategy                BidStrategy                     `json:"bid_strategy,omitempty"`
	ColdStartAudience          []int64                         `json:"cold_start_audience,omitempty"`
	AutoAudience               bool                            `json:"auto_audience,omitempty"`
	ExpandEnabled              bool                            `json:"expand_enabled,omitempty"`
	ExpandTargeting            []string                        `json:"expand_targeting,omitempty"`
	DeepConversionSpec         DeepConversionSpec              `json:"deep_conversion_spec,omitempty"`
	DeepOptimizationActionType DeepOptimizationActionType      `json:"deep_optimization_action_type,omitempty"`
	ConversionId               int64                           `json:"conversion_id,omitempty"`
	DeepConversionBehaviorBid  int64                           `json:"deep_conversion_behavior_bid,omitempty"`
	DeepConversionWorthRate    float64                         `json:"deep_conversion_worth_rate,omitempty"`
	AccountId                  int64                           `json:"account_id,omitempty"`
}