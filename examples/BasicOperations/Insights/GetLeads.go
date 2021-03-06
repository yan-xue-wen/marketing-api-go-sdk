/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
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
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type LeadsGetExample struct {
	TAds         *ads.SDKClient
	AccessToken  string
	AccountId    int64
	TimeRange    model.TimeRange
	PositionType string
	LeadsGetOpts *api.LeadsGetOpts
}

func (e *LeadsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.TimeRange = model.TimeRange{
		StartTime: int64(0),
		EndTime:   int64(0),
	}
	e.PositionType = "POSITION_TYPE_WECHAT_MOMENTS"
	e.LeadsGetOpts = &api.LeadsGetOpts{

		Fields: optional.NewInterface([]string{"campaign_id", "campaign_name", "adgroup_id", "adgroup_name", "wechat_adgroup_id", "lead_spec_list", "wechat_campaign_id", "wechat_campaign_name", "wechat_adgroup_name", "wechat_agency_id", "wechat_agency_name", "click_id"}),
	}
}

func (e *LeadsGetExample) RunExample() (model.LeadsGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Leads().Get(ctx, e.AccountId, e.TimeRange, e.PositionType, e.LeadsGetOpts)
}

func main() {
	e := &LeadsGetExample{}
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
