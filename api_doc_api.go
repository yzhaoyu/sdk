/*
Tencent Docs SDK

Tencent docs sdk contains DocAPI, SmartsheetAPI, SheetAPI, DriveAPI and OAuthAPI

API version: 0.0.1
Contact: tencentdocs@tencent.com
*/

// Code generated by Tencent Docs (https://docs.qq.com); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// DocAPIApiService DocAPIApi service
type DocAPIApiService service

type ApiDocAPIBatchUpdateDocDataRequest struct {
	ctx context.Context
	ApiService *DocAPIApiService
	fileID string
	accessToken *string
	clientId *string
	openId *string
	batchUpdateDocDataReq *BatchUpdateDocDataReq
}

// 访问令牌，用于标识用户和接口鉴权
func (r ApiDocAPIBatchUpdateDocDataRequest) AccessToken(accessToken string) ApiDocAPIBatchUpdateDocDataRequest {
	r.accessToken = &accessToken
	return r
}

// 应用 ID，用于标识应用和接口鉴权
func (r ApiDocAPIBatchUpdateDocDataRequest) ClientId(clientId string) ApiDocAPIBatchUpdateDocDataRequest {
	r.clientId = &clientId
	return r
}

// 开放平台用户 ID，用于标识用户和接口鉴权
func (r ApiDocAPIBatchUpdateDocDataRequest) OpenId(openId string) ApiDocAPIBatchUpdateDocDataRequest {
	r.openId = &openId
	return r
}

func (r ApiDocAPIBatchUpdateDocDataRequest) BatchUpdateDocDataReq(batchUpdateDocDataReq BatchUpdateDocDataReq) ApiDocAPIBatchUpdateDocDataRequest {
	r.batchUpdateDocDataReq = &batchUpdateDocDataReq
	return r
}

func (r ApiDocAPIBatchUpdateDocDataRequest) Execute() (*DocAPIResponse1, *http.Response, error) {
	return r.ApiService.DocAPIBatchUpdateDocDataExecute(r)
}

/*
DocAPIBatchUpdateDocData Method for DocAPIBatchUpdateDocData

BatchUpdateDocData 接口用于更新 Doc 文档内容，支持批量更新

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileID
 @return ApiDocAPIBatchUpdateDocDataRequest
*/
func (a *DocAPIApiService) DocAPIBatchUpdateDocData(ctx context.Context, fileID string) ApiDocAPIBatchUpdateDocDataRequest {
	return ApiDocAPIBatchUpdateDocDataRequest{
		ApiService: a,
		ctx: ctx,
		fileID: fileID,
	}
}

// Execute executes the request
//  @return DocAPIResponse1
func (a *DocAPIApiService) DocAPIBatchUpdateDocDataExecute(r ApiDocAPIBatchUpdateDocDataRequest) (*DocAPIResponse1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocAPIResponse1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocAPIApiService.DocAPIBatchUpdateDocData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openapi/doc/v3/{fileID}/batchUpdate"
	localVarPath = strings.Replace(localVarPath, "{"+"fileID"+"}", url.PathEscape(parameterToString(r.fileID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accessToken == nil {
		return localVarReturnValue, nil, reportError("accessToken is required and must be specified")
	}
	if r.clientId == nil {
		return localVarReturnValue, nil, reportError("clientId is required and must be specified")
	}
	if r.openId == nil {
		return localVarReturnValue, nil, reportError("openId is required and must be specified")
	}
	if r.batchUpdateDocDataReq == nil {
		return localVarReturnValue, nil, reportError("batchUpdateDocDataReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Access-Token"] = parameterToString(*r.accessToken, "")
	localVarHeaderParams["Client-Id"] = parameterToString(*r.clientId, "")
	localVarHeaderParams["Open-Id"] = parameterToString(*r.openId, "")
	// body params
	localVarPostBody = r.batchUpdateDocDataReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDocAPIGetDocFullTextRequest struct {
	ctx context.Context
	ApiService *DocAPIApiService
	fileID string
	accessToken *string
	clientId *string
	openId *string
}

// 访问令牌，用于标识用户和接口鉴权
func (r ApiDocAPIGetDocFullTextRequest) AccessToken(accessToken string) ApiDocAPIGetDocFullTextRequest {
	r.accessToken = &accessToken
	return r
}

// 应用 ID，用于标识应用和接口鉴权
func (r ApiDocAPIGetDocFullTextRequest) ClientId(clientId string) ApiDocAPIGetDocFullTextRequest {
	r.clientId = &clientId
	return r
}

// 开放平台用户 ID，用于标识用户和接口鉴权
func (r ApiDocAPIGetDocFullTextRequest) OpenId(openId string) ApiDocAPIGetDocFullTextRequest {
	r.openId = &openId
	return r
}

func (r ApiDocAPIGetDocFullTextRequest) Execute() (*DocAPIResponse2, *http.Response, error) {
	return r.ApiService.DocAPIGetDocFullTextExecute(r)
}

/*
DocAPIGetDocFullText Method for DocAPIGetDocFullText

GetDocFullText 接口用于获取 Doc 文档内容

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileID 文档 ID
 @return ApiDocAPIGetDocFullTextRequest
*/
func (a *DocAPIApiService) DocAPIGetDocFullText(ctx context.Context, fileID string) ApiDocAPIGetDocFullTextRequest {
	return ApiDocAPIGetDocFullTextRequest{
		ApiService: a,
		ctx: ctx,
		fileID: fileID,
	}
}

// Execute executes the request
//  @return DocAPIResponse2
func (a *DocAPIApiService) DocAPIGetDocFullTextExecute(r ApiDocAPIGetDocFullTextRequest) (*DocAPIResponse2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocAPIResponse2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocAPIApiService.DocAPIGetDocFullText")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openapi/doc/v3/{fileID}"
	localVarPath = strings.Replace(localVarPath, "{"+"fileID"+"}", url.PathEscape(parameterToString(r.fileID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accessToken == nil {
		return localVarReturnValue, nil, reportError("accessToken is required and must be specified")
	}
	if r.clientId == nil {
		return localVarReturnValue, nil, reportError("clientId is required and must be specified")
	}
	if r.openId == nil {
		return localVarReturnValue, nil, reportError("openId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Access-Token"] = parameterToString(*r.accessToken, "")
	localVarHeaderParams["Client-Id"] = parameterToString(*r.clientId, "")
	localVarHeaderParams["Open-Id"] = parameterToString(*r.openId, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
