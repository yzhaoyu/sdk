# \ResourceAPIApi

All URIs are relative to *https://docs.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourceAPIUploadImage**](ResourceAPIApi.md#ResourceAPIUploadImage) | **Post** /openapi/resources/v2/images | 



## ResourceAPIUploadImage

> ResourcesAPIResponse1 ResourceAPIUploadImage(ctx).AccessToken(accessToken).ClientId(clientId).OpenId(openId).Image(image).Execute()





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
    image := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceAPIApi.ResourceAPIUploadImage(context.Background()).AccessToken(accessToken).ClientId(clientId).OpenId(openId).Image(image).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceAPIApi.ResourceAPIUploadImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResourceAPIUploadImage`: ResourcesAPIResponse1
    fmt.Fprintf(os.Stdout, "Response from `ResourceAPIApi.ResourceAPIUploadImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResourceAPIUploadImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string** | 访问令牌，用于标识用户和接口鉴权 | 
 **clientId** | **string** | 应用 ID，用于标识应用和接口鉴权 | 
 **openId** | **string** | 开放平台用户 ID，用于标识用户和接口鉴权 | 
 **image** | ***os.File** |  | 

### Return type

[**ResourcesAPIResponse1**](ResourcesAPIResponse1.md)

### Authorization

[OAuth2AuthCode](../README.md#OAuth2AuthCode)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

