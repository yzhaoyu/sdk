# \SmartsheetAPIApi

All URIs are relative to *https://docs.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmartsheetAPIEditSmartsheetSubSheet**](SmartsheetAPIApi.md#SmartsheetAPIEditSmartsheetSubSheet) | **Post** /openapi/smartbook/v2/files/{fileID}/sheets | 
[**SmartsheetAPIGetSmartsheetSubSheet**](SmartsheetAPIApi.md#SmartsheetAPIGetSmartsheetSubSheet) | **Get** /openapi/smartbook/v2/files/{fileID}/sheets | 
[**SmartsheetAPIHandleSmartsheetResource**](SmartsheetAPIApi.md#SmartsheetAPIHandleSmartsheetResource) | **Post** /openapi/smartbook/v2/files/{fileID}/sheets/{sheetID} | 



## SmartsheetAPIEditSmartsheetSubSheet

> SmartsheetAPIResponse1 SmartsheetAPIEditSmartsheetSubSheet(ctx, fileID).AccessToken(accessToken).ClientId(clientId).OpenId(openId).EditSubSheetRequest(editSubSheetRequest).Execute()





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
    editSubSheetRequest := *openapiclient.NewEditSubSheetRequest("FileID_example") // EditSubSheetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartsheetAPIApi.SmartsheetAPIEditSmartsheetSubSheet(context.Background(), fileID).AccessToken(accessToken).ClientId(clientId).OpenId(openId).EditSubSheetRequest(editSubSheetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartsheetAPIApi.SmartsheetAPIEditSmartsheetSubSheet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SmartsheetAPIEditSmartsheetSubSheet`: SmartsheetAPIResponse1
    fmt.Fprintf(os.Stdout, "Response from `SmartsheetAPIApi.SmartsheetAPIEditSmartsheetSubSheet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmartsheetAPIEditSmartsheetSubSheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **string** | 访问令牌，用于标识用户和接口鉴权 | 
 **clientId** | **string** | 应用 ID，用于标识应用和接口鉴权 | 
 **openId** | **string** | 开放平台用户 ID，用于标识用户和接口鉴权 | 
 **editSubSheetRequest** | [**EditSubSheetRequest**](EditSubSheetRequest.md) |  | 

### Return type

[**SmartsheetAPIResponse1**](SmartsheetAPIResponse1.md)

### Authorization

[OAuth2AuthCode](../README.md#OAuth2AuthCode)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartsheetAPIGetSmartsheetSubSheet

> SmartsheetAPIResponse2 SmartsheetAPIGetSmartsheetSubSheet(ctx, fileID).AccessToken(accessToken).ClientId(clientId).OpenId(openId).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartsheetAPIApi.SmartsheetAPIGetSmartsheetSubSheet(context.Background(), fileID).AccessToken(accessToken).ClientId(clientId).OpenId(openId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartsheetAPIApi.SmartsheetAPIGetSmartsheetSubSheet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SmartsheetAPIGetSmartsheetSubSheet`: SmartsheetAPIResponse2
    fmt.Fprintf(os.Stdout, "Response from `SmartsheetAPIApi.SmartsheetAPIGetSmartsheetSubSheet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmartsheetAPIGetSmartsheetSubSheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **string** | 访问令牌，用于标识用户和接口鉴权 | 
 **clientId** | **string** | 应用 ID，用于标识应用和接口鉴权 | 
 **openId** | **string** | 开放平台用户 ID，用于标识用户和接口鉴权 | 

### Return type

[**SmartsheetAPIResponse2**](SmartsheetAPIResponse2.md)

### Authorization

[OAuth2AuthCode](../README.md#OAuth2AuthCode)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartsheetAPIHandleSmartsheetResource

> SmartsheetAPIResponse3 SmartsheetAPIHandleSmartsheetResource(ctx, fileID, sheetID).AccessToken(accessToken).ClientId(clientId).OpenId(openId).SmartsheetResourceRequest(smartsheetResourceRequest).Execute()





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
    sheetID := "sheetID_example" // string | 
    accessToken := "accessToken_example" // string | 访问令牌，用于标识用户和接口鉴权
    clientId := "clientId_example" // string | 应用 ID，用于标识应用和接口鉴权
    openId := "openId_example" // string | 开放平台用户 ID，用于标识用户和接口鉴权
    smartsheetResourceRequest := *openapiclient.NewSmartsheetResourceRequest("FileID_example", "SheetID_example") // SmartsheetResourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartsheetAPIApi.SmartsheetAPIHandleSmartsheetResource(context.Background(), fileID, sheetID).AccessToken(accessToken).ClientId(clientId).OpenId(openId).SmartsheetResourceRequest(smartsheetResourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartsheetAPIApi.SmartsheetAPIHandleSmartsheetResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SmartsheetAPIHandleSmartsheetResource`: SmartsheetAPIResponse3
    fmt.Fprintf(os.Stdout, "Response from `SmartsheetAPIApi.SmartsheetAPIHandleSmartsheetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string** |  | 
**sheetID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmartsheetAPIHandleSmartsheetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessToken** | **string** | 访问令牌，用于标识用户和接口鉴权 | 
 **clientId** | **string** | 应用 ID，用于标识应用和接口鉴权 | 
 **openId** | **string** | 开放平台用户 ID，用于标识用户和接口鉴权 | 
 **smartsheetResourceRequest** | [**SmartsheetResourceRequest**](SmartsheetResourceRequest.md) |  | 

### Return type

[**SmartsheetAPIResponse3**](SmartsheetAPIResponse3.md)

### Authorization

[OAuth2AuthCode](../README.md#OAuth2AuthCode)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

