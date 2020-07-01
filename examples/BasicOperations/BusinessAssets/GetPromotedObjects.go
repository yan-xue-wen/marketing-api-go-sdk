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

type PromotedObjectsGetExample struct {
	TAds                   *ads.SDKClient
	AccessToken            string
	AccountId              int64
	PromotedObjectsGetOpts *api.PromotedObjectsGetOpts
}

func (e *PromotedObjectsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.PromotedObjectsGetOpts = &api.PromotedObjectsGetOpts{

		Fields: optional.NewInterface([]string{"promoted_object_name", "promoted_object_id", "promoted_object_type", "promoted_object_spec", "created_time", "last_modified_time"}),
	}
}

func (e *PromotedObjectsGetExample) RunExample() (interface{}, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.PromotedObjects().Get(ctx, e.AccountId, e.PromotedObjectsGetOpts)
}

func main() {
	e := &PromotedObjectsGetExample{}
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