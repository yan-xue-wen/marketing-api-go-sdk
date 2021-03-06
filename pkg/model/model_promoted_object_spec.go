/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 推广目标详细信息
type PromotedObjectSpec struct {
	JdItem       *EcInfo                  `json:"jd_item,omitempty"`
	JdShop       *EcInfo                  `json:"jd_shop,omitempty"`
	DianpingShop *ProductTypeDianpingShop `json:"dianping_shop,omitempty"`
}
