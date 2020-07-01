/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
)

type CustomAudienceFilesGetExample struct {
	TAds                       *ads.SDKClient
	AccessToken                string
	AccountId                  int64
	CustomAudienceFilesGetOpts *api.CustomAudienceFilesGetOpts
}

func (e *CustomAudienceFilesGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.CustomAudienceFilesGetOpts = &api.CustomAudienceFilesGetOpts{

		AudienceId: optional.NewInt64(int64(0)),

		Fields: optional.NewInterface([]string{"audience_id", "custom_audience_file_id", "name", "user_id_type", "operation_type", "open_app_id", "salt_id", "process_status", "process_code", "error_message", "line_count", "valid_line_count", "user_count", "size", "created_time"}),
	}
}

func (e *CustomAudienceFilesGetExample) RunExample() (interface{}, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.CustomAudienceFiles().Get(ctx, e.AccountId, e.CustomAudienceFilesGetOpts)
}

func main() {
	e := &CustomAudienceFilesGetExample{}
	e.Init()
	response, httpResponse, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Http response:", httpResponse)
}