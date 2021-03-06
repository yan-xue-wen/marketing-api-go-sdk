/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// LearningStatus : 学习状态
type LearningStatus string

// List of LearningStatus
const (
	LearningStatus_PROCESSING       LearningStatus = "LEARNING_STATUS_PROCESSING"
	LearningStatus_SUGGEST_CONTINUE LearningStatus = "LEARNING_STATUS_SUGGEST_CONTINUE"
	LearningStatus_SUGGEST_STOP     LearningStatus = "LEARNING_STATUS_SUGGEST_STOP"
	LearningStatus_SUGGEST_IMPROVE  LearningStatus = "LEARNING_STATUS_SUGGEST_IMPROVE"
	LearningStatus_UNKNOWN          LearningStatus = "LEARNING_STATUS_UNKNOWN"
	LearningStatus_WIP              LearningStatus = "LEARNING_STATUS_WIP"
	LearningStatus_FINISHED         LearningStatus = "LEARNING_STATUS_FINISHED"
	LearningStatus_FAILED           LearningStatus = "LEARNING_STATUS_FAILED"
)
