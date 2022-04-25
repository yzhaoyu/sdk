# {{classname}}

All URIs are relative to *https://docs.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocAPIBatchUpdateDocData**](DocAPIApi.md#DocAPIBatchUpdateDocData) | **Post** /openapi/doc/v3/{fileID}/batchUpdate | 
[**DocAPIGetDocFullText**](DocAPIApi.md#DocAPIGetDocFullText) | **Get** /openapi/doc/v3/{fileID} | 

# **DocAPIBatchUpdateDocData**
> UpdateDocResponse DocAPIBatchUpdateDocData(ctx, body, fileID)


BatchUpdateDocData 接口用于更新 Doc 文档内容，支持批量更新

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchUpdateDocDataReq**](BatchUpdateDocDataReq.md)|  | 
  **fileID** | **string**|  | 

### Return type

[**UpdateDocResponse**](UpdateDocResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocAPIGetDocFullText**
> GetDocResponse DocAPIGetDocFullText(ctx, fileID)


GetDocFullText 接口用于获取 Doc 文档内容

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileID** | **string**| 文档 ID | 

### Return type

[**GetDocResponse**](GetDocResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

