# \DriveAPIApi

All URIs are relative to *https://docs.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveAPICreateFile**](DriveAPIApi.md#DriveAPICreateFile) | **Post** /openapi/drive/v2/files | 



## DriveAPICreateFile

> DriveAPIResponse1 DriveAPICreateFile(ctx).AccessToken(accessToken).ClientId(clientId).OpenId(openId).CreateFileReq(createFileReq).Execute()





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
    accessToken := "accessToken_example" // string | 访问令牌，用于标识用户和接口鉴权
    clientId := "clientId_example" // string | 应用 ID，用于标识应用和接口鉴权
    openId := "openId_example" // string | 开放平台用户 ID，用于标识用户和接口鉴权
    createFileReq := *openapiclient.NewCreateFileReq("Title_example", "Type_example") // CreateFileReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DriveAPIApi.DriveAPICreateFile(context.Background()).AccessToken(accessToken).ClientId(clientId).OpenId(openId).CreateFileReq(createFileReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveAPIApi.DriveAPICreateFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveAPICreateFile`: DriveAPIResponse1
    fmt.Fprintf(os.Stdout, "Response from `DriveAPIApi.DriveAPICreateFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveAPICreateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string** | 访问令牌，用于标识用户和接口鉴权 | 
 **clientId** | **string** | 应用 ID，用于标识应用和接口鉴权 | 
 **openId** | **string** | 开放平台用户 ID，用于标识用户和接口鉴权 | 
 **createFileReq** | [**CreateFileReq**](CreateFileReq.md) |  | 

### Return type

[**DriveAPIResponse1**](DriveAPIResponse1.md)

### Authorization

[OAuth2AuthCode](../README.md#OAuth2AuthCode)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

