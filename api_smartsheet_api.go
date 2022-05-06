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


// SmartsheetAPIApiService SmartsheetAPIApi service
type SmartsheetAPIApiService service

type ApiSmartsheetAPIEditSmartsheetSubSheetRequest struct {
	ctx context.Context
	ApiService *SmartsheetAPIApiService
	fileID string
	accessToken *string
	clientId *string
	openId *string
	editSubSheetRequest *EditSubSheetRequest
}

// 访问令牌，用于标识用户和接口鉴权
func (r ApiSmartsheetAPIEditSmartsheetSubSheetRequest) AccessToken(accessToken string) ApiSmartsheetAPIEditSmartsheetSubSheetRequest {
	r.accessToken = &accessToken
	return r
}

// 应用 ID，用于标识应用和接口鉴权
func (r ApiSmartsheetAPIEditSmartsheetSubSheetRequest) ClientId(clientId string) ApiSmartsheetAPIEditSmartsheetSubSheetRequest {
	r.clientId = &clientId
	return r
}

// 开放平台用户 ID，用于标识用户和接口鉴权
func (r ApiSmartsheetAPIEditSmartsheetSubSheetRequest) OpenId(openId string) ApiSmartsheetAPIEditSmartsheetSubSheetRequest {
	r.openId = &openId
	return r
}

func (r ApiSmartsheetAPIEditSmartsheetSubSheetRequest) EditSubSheetRequest(editSubSheetRequest EditSubSheetRequest) ApiSmartsheetAPIEditSmartsheetSubSheetRequest {
	r.editSubSheetRequest = &editSubSheetRequest
	return r
}

func (r ApiSmartsheetAPIEditSmartsheetSubSheetRequest) Execute() (*SmartsheetAPIResponse1, *http.Response, error) {
	return r.ApiService.SmartsheetAPIEditSmartsheetSubSheetExecute(r)
}

/*
SmartsheetAPIEditSmartsheetSubSheet Method for SmartsheetAPIEditSmartsheetSubSheet

Three required req headers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileID
 @return ApiSmartsheetAPIEditSmartsheetSubSheetRequest
*/
func (a *SmartsheetAPIApiService) SmartsheetAPIEditSmartsheetSubSheet(ctx context.Context, fileID string) ApiSmartsheetAPIEditSmartsheetSubSheetRequest {
	return ApiSmartsheetAPIEditSmartsheetSubSheetRequest{
		ApiService: a,
		ctx: ctx,
		fileID: fileID,
	}
}

// Execute executes the request
//  @return SmartsheetAPIResponse1
func (a *SmartsheetAPIApiService) SmartsheetAPIEditSmartsheetSubSheetExecute(r ApiSmartsheetAPIEditSmartsheetSubSheetRequest) (*SmartsheetAPIResponse1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmartsheetAPIResponse1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartsheetAPIApiService.SmartsheetAPIEditSmartsheetSubSheet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openapi/smartbook/v2/files/{fileID}/sheets"
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
	if r.editSubSheetRequest == nil {
		return localVarReturnValue, nil, reportError("editSubSheetRequest is required and must be specified")
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
	localVarPostBody = r.editSubSheetRequest
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

type ApiSmartsheetAPIGetSmartsheetSubSheetRequest struct {
	ctx context.Context
	ApiService *SmartsheetAPIApiService
	fileID string
	accessToken *string
	clientId *string
	openId *string
}

// 访问令牌，用于标识用户和接口鉴权
func (r ApiSmartsheetAPIGetSmartsheetSubSheetRequest) AccessToken(accessToken string) ApiSmartsheetAPIGetSmartsheetSubSheetRequest {
	r.accessToken = &accessToken
	return r
}

// 应用 ID，用于标识应用和接口鉴权
func (r ApiSmartsheetAPIGetSmartsheetSubSheetRequest) ClientId(clientId string) ApiSmartsheetAPIGetSmartsheetSubSheetRequest {
	r.clientId = &clientId
	return r
}

// 开放平台用户 ID，用于标识用户和接口鉴权
func (r ApiSmartsheetAPIGetSmartsheetSubSheetRequest) OpenId(openId string) ApiSmartsheetAPIGetSmartsheetSubSheetRequest {
	r.openId = &openId
	return r
}

func (r ApiSmartsheetAPIGetSmartsheetSubSheetRequest) Execute() (*SmartsheetAPIResponse2, *http.Response, error) {
	return r.ApiService.SmartsheetAPIGetSmartsheetSubSheetExecute(r)
}

/*
SmartsheetAPIGetSmartsheetSubSheet Method for SmartsheetAPIGetSmartsheetSubSheet

Three required req headers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileID
 @return ApiSmartsheetAPIGetSmartsheetSubSheetRequest
*/
func (a *SmartsheetAPIApiService) SmartsheetAPIGetSmartsheetSubSheet(ctx context.Context, fileID string) ApiSmartsheetAPIGetSmartsheetSubSheetRequest {
	return ApiSmartsheetAPIGetSmartsheetSubSheetRequest{
		ApiService: a,
		ctx: ctx,
		fileID: fileID,
	}
}

// Execute executes the request
//  @return SmartsheetAPIResponse2
func (a *SmartsheetAPIApiService) SmartsheetAPIGetSmartsheetSubSheetExecute(r ApiSmartsheetAPIGetSmartsheetSubSheetRequest) (*SmartsheetAPIResponse2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmartsheetAPIResponse2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartsheetAPIApiService.SmartsheetAPIGetSmartsheetSubSheet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openapi/smartbook/v2/files/{fileID}/sheets"
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

type ApiSmartsheetAPIHandleSmartsheetResourceRequest struct {
	ctx context.Context
	ApiService *SmartsheetAPIApiService
	fileID string
	sheetID string
	accessToken *string
	clientId *string
	openId *string
	smartsheetResourceRequest *SmartsheetResourceRequest
}

// 访问令牌，用于标识用户和接口鉴权
func (r ApiSmartsheetAPIHandleSmartsheetResourceRequest) AccessToken(accessToken string) ApiSmartsheetAPIHandleSmartsheetResourceRequest {
	r.accessToken = &accessToken
	return r
}

// 应用 ID，用于标识应用和接口鉴权
func (r ApiSmartsheetAPIHandleSmartsheetResourceRequest) ClientId(clientId string) ApiSmartsheetAPIHandleSmartsheetResourceRequest {
	r.clientId = &clientId
	return r
}

// 开放平台用户 ID，用于标识用户和接口鉴权
func (r ApiSmartsheetAPIHandleSmartsheetResourceRequest) OpenId(openId string) ApiSmartsheetAPIHandleSmartsheetResourceRequest {
	r.openId = &openId
	return r
}

func (r ApiSmartsheetAPIHandleSmartsheetResourceRequest) SmartsheetResourceRequest(smartsheetResourceRequest SmartsheetResourceRequest) ApiSmartsheetAPIHandleSmartsheetResourceRequest {
	r.smartsheetResourceRequest = &smartsheetResourceRequest
	return r
}

func (r ApiSmartsheetAPIHandleSmartsheetResourceRequest) Execute() (*SmartsheetAPIResponse3, *http.Response, error) {
	return r.ApiService.SmartsheetAPIHandleSmartsheetResourceExecute(r)
}

/*
SmartsheetAPIHandleSmartsheetResource Method for SmartsheetAPIHandleSmartsheetResource

Three required req headers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileID
 @param sheetID
 @return ApiSmartsheetAPIHandleSmartsheetResourceRequest
*/
func (a *SmartsheetAPIApiService) SmartsheetAPIHandleSmartsheetResource(ctx context.Context, fileID string, sheetID string) ApiSmartsheetAPIHandleSmartsheetResourceRequest {
	return ApiSmartsheetAPIHandleSmartsheetResourceRequest{
		ApiService: a,
		ctx: ctx,
		fileID: fileID,
		sheetID: sheetID,
	}
}

// Execute executes the request
//  @return SmartsheetAPIResponse3
func (a *SmartsheetAPIApiService) SmartsheetAPIHandleSmartsheetResourceExecute(r ApiSmartsheetAPIHandleSmartsheetResourceRequest) (*SmartsheetAPIResponse3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmartsheetAPIResponse3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartsheetAPIApiService.SmartsheetAPIHandleSmartsheetResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openapi/smartbook/v2/files/{fileID}/sheets/{sheetID}"
	localVarPath = strings.Replace(localVarPath, "{"+"fileID"+"}", url.PathEscape(parameterToString(r.fileID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetID"+"}", url.PathEscape(parameterToString(r.sheetID, "")), -1)

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
	if r.smartsheetResourceRequest == nil {
		return localVarReturnValue, nil, reportError("smartsheetResourceRequest is required and must be specified")
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
	localVarPostBody = r.smartsheetResourceRequest
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