/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 不支持的属性
type UnsupportSpec struct {
	SiteSet       []string      `json:"site_set,omitempty"`
	Name          string        `json:"name,omitempty"`
	UnsupportType UnsupportType `json:"unsupport_type,omitempty"`
}
