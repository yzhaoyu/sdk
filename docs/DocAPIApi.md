# {{classname}}

All URIs are relative to *https://docs.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocAPIBatchUpdateDocData**](DocAPIApi.md#DocAPIBatchUpdateDocData) | **Post** /openapi/doc/v3/{fileID}/batchUpdate | 
[**DocAPIGetDocFullText**](DocAPIApi.md#DocAPIGetDocFullText) | **Get** /openapi/doc/v3/{fileID} | 

# **DocAPIBatchUpdateDocData**
> UserResponse DocAPIBatchUpdateDocData(ctx, body, fileID, accessToken, clientId, openId)


Three required req headers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchUpdateDocDataReq**](BatchUpdateDocDataReq.md)|  | 
  **fileID** | **string**|  | 
  **accessToken** | **string**| 访问令牌，用于标识用户和接口鉴权 | 
  **clientId** | **string**| 应用 ID，用于标识应用和接口鉴权 | 
  **openId** | **string**| 开放平台用户 ID，用于标识用户和接口鉴权 | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocAPIGetDocFullText**
> UserResponse DocAPIGetDocFullText(ctx, fileID, accessToken, clientId, openId)


Three required req headers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileID** | **string**| 文档 ID | 
  **accessToken** | **string**| 访问令牌，用于标识用户和接口鉴权 | 
  **clientId** | **string**| 应用 ID，用于标识应用和接口鉴权 | 
  **openId** | **string**| 开放平台用户 ID，用于标识用户和接口鉴权 | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

