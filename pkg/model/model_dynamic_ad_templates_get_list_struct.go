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
type DynamicAdTemplatesGetListStruct struct {
	DynamicAdTemplateId        int64                      `json:"dynamic_ad_template_id,omitempty"`
	DynamicAdTemplateName      string                     `json:"dynamic_ad_template_name,omitempty"`
	DynamicAdTemplateType      DynamicAdTemplateType      `json:"dynamic_ad_template_type,omitempty"`
	ProductItemDisplayQuantity ProductItemDisplayQuantity `json:"product_item_display_quantity,omitempty"`
	DynamicAdTemplateWidth     int64                      `json:"dynamic_ad_template_width,omitempty"`
	DynamicAdTemplateHeight    int64                      `json:"dynamic_ad_template_height,omitempty"`
	ImageUrl                   string                     `json:"image_url,omitempty"`
	VideoUrl                   string                     `json:"video_url,omitempty"`
}