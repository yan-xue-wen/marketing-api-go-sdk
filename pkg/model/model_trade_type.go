/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// TradeType : 交易类型
type TradeType string

// List of TradeType
const (
	TradeType_CHARGE        TradeType = "CHARGE"
	TradeType_PAY           TradeType = "PAY"
	TradeType_TRANSFER_BACK TradeType = "TRANSFER_BACK"
	TradeType_EXPIRE        TradeType = "EXPIRE"
)
