# \OAuthAPIApi

All URIs are relative to *https://docs.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OAuthAPIGetToken**](OAuthAPIApi.md#OAuthAPIGetToken) | **Get** /oauth/v2/token | 
[**OAuthAPIRefreshToken**](OAuthAPIApi.md#OAuthAPIRefreshToken) | **Get** /oauth/v2/userinfo | 



## OAuthAPIGetToken

> OAuthAPIResponse1 OAuthAPIGetToken(ctx).ClientId(clientId).ClientSecret(clientSecret).RedirectUri(redirectUri).GrantType(grantType).Code(code).Execute()



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
    clientId := "clientId_example" // string |  (optional)
    clientSecret := "clientSecret_example" // string |  (optional)
    redirectUri := "redirectUri_example" // string |  (optional)
    grantType := "grantType_example" // string |  (optional)
    code := "code_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthAPIApi.OAuthAPIGetToken(context.Background()).ClientId(clientId).ClientSecret(clientSecret).RedirectUri(redirectUri).GrantType(grantType).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPIApi.OAuthAPIGetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OAuthAPIGetToken`: OAuthAPIResponse1
    fmt.Fprintf(os.Stdout, "Response from `OAuthAPIApi.OAuthAPIGetToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOAuthAPIGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 
 **redirectUri** | **string** |  | 
 **grantType** | **string** |  | 
 **code** | **string** |  | 

### Return type

[**OAuthAPIResponse1**](OAuthAPIResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OAuthAPIRefreshToken

> OAuthAPIResponse3 OAuthAPIRefreshToken(ctx).ClientId(clientId).ClientSecret(clientSecret).GrantType(grantType).RefreshToken(refreshToken).Execute()



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
    clientId := "clientId_example" // string |  (optional)
    clientSecret := "clientSecret_example" // string |  (optional)
    grantType := "grantType_example" // string |  (optional)
    refreshToken := "refreshToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthAPIApi.OAuthAPIRefreshToken(context.Background()).ClientId(clientId).ClientSecret(clientSecret).GrantType(grantType).RefreshToken(refreshToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPIApi.OAuthAPIRefreshToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OAuthAPIRefreshToken`: OAuthAPIResponse3
    fmt.Fprintf(os.Stdout, "Response from `OAuthAPIApi.OAuthAPIRefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOAuthAPIRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 
 **grantType** | **string** |  | 
 **refreshToken** | **string** |  | 

### Return type

[**OAuthAPIResponse3**](OAuthAPIResponse3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

