/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// ProcessStatus : 处理状态
type ProcessStatus string

// List of ProcessStatus
const (
	ProcessStatus_PENDING    ProcessStatus = "PENDING"
	ProcessStatus_PROCESSING ProcessStatus = "PROCESSING"
	ProcessStatus_SUCCESS    ProcessStatus = "SUCCESS"
	ProcessStatus_ERROR_     ProcessStatus = "ERROR"
	ProcessStatus_FROZEN     ProcessStatus = "FROZEN"
	ProcessStatus_THAWING    ProcessStatus = "THAWING"
	ProcessStatus_LOCKING    ProcessStatus = "LOCKING"
)