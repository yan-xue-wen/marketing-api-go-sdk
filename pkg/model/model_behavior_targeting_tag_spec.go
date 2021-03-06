/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 获取行为定向标签的条件，type 为 BEHAVIOR 时必填
type BehaviorTargetingTagSpec struct {
	QueryMode TargetingTagQueryMode `json:"query_mode,omitempty"`
	QuerySpec *QuerySpec            `json:"query_spec,omitempty"`
}
