/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type WechatAdFollowersGetResponseData struct {
	List     []WechatAdFollowersGetListStruct `json:"list,omitempty"`
	PageInfo Conf                             `json:"page_info,omitempty"`
}