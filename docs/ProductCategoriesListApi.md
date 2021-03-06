# TencentAds\ProductCategoriesListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductCategoriesListGet**](ProductCategoriesListApi.md#ProductCategoriesListGet) | **Get** /product_categories_list/get | 获取商品类目


# **ProductCategoriesListGet**
> ProductCategoriesListGetResponse ProductCategoriesListGet(ctx, accountId, productCatalogId, optional)
获取商品类目

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **productCatalogId** | **int64**|  | 
 **optional** | ***ProductCategoriesListApiProductCategoriesListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCategoriesListApiProductCategoriesListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **categoryName** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProductCategoriesListGetResponse**](ProductCategoriesListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

