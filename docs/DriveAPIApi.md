# \DriveAPIApi

All URIs are relative to *https://docs.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveAPICreateFile**](DriveAPIApi.md#DriveAPICreateFile) | **Post** /openapi/drive/v2/files | 



## DriveAPICreateFile

> DriveapiDriveAPIResponse1 DriveAPICreateFile(ctx).AccessToken(accessToken).ClientId(clientId).OpenId(openId).Title(title).Type_(type_).TemplateID(templateID).FolderID(folderID).Execute()





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
    title := "title_example" // string | 
    type_ := "type__example" // string | 
    templateID := "templateID_example" // string |  (optional)
    folderID := "folderID_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DriveAPIApi.DriveAPICreateFile(context.Background()).AccessToken(accessToken).ClientId(clientId).OpenId(openId).Title(title).Type_(type_).TemplateID(templateID).FolderID(folderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveAPIApi.DriveAPICreateFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveAPICreateFile`: DriveapiDriveAPIResponse1
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
 **title** | **string** |  | 
 **type_** | **string** |  | 
 **templateID** | **string** |  | 
 **folderID** | **string** |  | 

### Return type

[**DriveapiDriveAPIResponse1**](DriveapiDriveAPIResponse1.md)

### Authorization

[OAuth2AuthCode](../README.md#OAuth2AuthCode)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

