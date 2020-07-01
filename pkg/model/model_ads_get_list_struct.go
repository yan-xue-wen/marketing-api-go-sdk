/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type AdsGetListStruct struct {
	CampaignId              int64             `json:"campaign_id,omitempty"`
	AdgroupId               int64             `json:"adgroup_id,omitempty"`
	AdId                    int64             `json:"ad_id,omitempty"`
	AdName                  string            `json:"ad_name,omitempty"`
	AdcreativeId            int64             `json:"adcreative_id,omitempty"`
	Adcreative              Adcreative        `json:"adcreative,omitempty"`
	ConfiguredStatus        AdStatus          `json:"configured_status,omitempty"`
	SystemStatus            SysStatus         `json:"system_status,omitempty"`
	AuditSpec               []AuditSpecStruct `json:"audit_spec,omitempty"`
	ImpressionTrackingUrl   string            `json:"impression_tracking_url,omitempty"`
	ClickTrackingUrl        string            `json:"click_tracking_url,omitempty"`
	FeedsInteractionEnabled bool              `json:"feeds_interaction_enabled,omitempty"`
	IsDeleted               bool              `json:"is_deleted,omitempty"`
	IsDynamicCreative       bool              `json:"is_dynamic_creative,omitempty"`
	CreatedTime             int64             `json:"created_time,omitempty"`
	LastModifiedTime        int64             `json:"last_modified_time,omitempty"`
	RejectMessage           string            `json:"reject_message,omitempty"`
}