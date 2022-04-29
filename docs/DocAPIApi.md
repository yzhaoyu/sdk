# \DocAPIApi

All URIs are relative to *https://docs.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocAPIBatchUpdateDocData**](DocAPIApi.md#DocAPIBatchUpdateDocData) | **Post** /openapi/doc/v3/{fileID}/batchUpdate | 
[**DocAPIGetDocFullText**](DocAPIApi.md#DocAPIGetDocFullText) | **Get** /openapi/doc/v3/{fileID} | 



## DocAPIBatchUpdateDocData

> OpenResponse1 DocAPIBatchUpdateDocData(ctx, fileID).AccessToken(accessToken).ClientId(clientId).OpenId(openId).BatchUpdateDocDataReq(batchUpdateDocDataReq).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fileID := "fileID_example" // string | 
    accessToken := "accessToken_example" // string | 访问令牌，用于标识用户和接口鉴权
    clientId := "clientId_example" // string | 应用 ID，用于标识应用和接口鉴权
    openId := "openId_example" // string | 开放平台用户 ID，用于标识用户和接口鉴权
    batchUpdateDocDataReq := *openapiclient.NewBatchUpdateDocDataReq("FileID_example", []openapiclient.Request{*openapiclient.NewRequest()}) // BatchUpdateDocDataReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocAPIApi.DocAPIBatchUpdateDocData(context.Background(), fileID).AccessToken(accessToken).ClientId(clientId).OpenId(openId).BatchUpdateDocDataReq(batchUpdateDocDataReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocAPIApi.DocAPIBatchUpdateDocData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocAPIBatchUpdateDocData`: OpenResponse1
    fmt.Fprintf(os.Stdout, "Response from `DocAPIApi.DocAPIBatchUpdateDocData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocAPIBatchUpdateDocDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **string** | 访问令牌，用于标识用户和接口鉴权 | 
 **clientId** | **string** | 应用 ID，用于标识应用和接口鉴权 | 
 **openId** | **string** | 开放平台用户 ID，用于标识用户和接口鉴权 | 
 **batchUpdateDocDataReq** | [**BatchUpdateDocDataReq**](BatchUpdateDocDataReq.md) |  | 

### Return type

[**OpenResponse1**](OpenResponse1.md)

### Authorization

[OAuth2AuthCode](../README.md#OAuth2AuthCode)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocAPIGetDocFullText

> OpenResponse2 DocAPIGetDocFullText(ctx, fileID).AccessToken(accessToken).ClientId(clientId).OpenId(openId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fileID := "fileID_example" // string | 文档 ID
    accessToken := "accessToken_example" // string | 访问令牌，用于标识用户和接口鉴权
    clientId := "clientId_example" // string | 应用 ID，用于标识应用和接口鉴权
    openId := "openId_example" // string | 开放平台用户 ID，用于标识用户和接口鉴权

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocAPIApi.DocAPIGetDocFullText(context.Background(), fileID).AccessToken(accessToken).ClientId(clientId).OpenId(openId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocAPIApi.DocAPIGetDocFullText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocAPIGetDocFullText`: OpenResponse2
    fmt.Fprintf(os.Stdout, "Response from `DocAPIApi.DocAPIGetDocFullText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string** | 文档 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocAPIGetDocFullTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **string** | 访问令牌，用于标识用户和接口鉴权 | 
 **clientId** | **string** | 应用 ID，用于标识应用和接口鉴权 | 
 **openId** | **string** | 开放平台用户 ID，用于标识用户和接口鉴权 | 

### Return type

[**OpenResponse2**](OpenResponse2.md)

### Authorization

[OAuth2AuthCode](../README.md#OAuth2AuthCode)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

