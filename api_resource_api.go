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
)


// ResourceAPIApiService ResourceAPIApi service
type ResourceAPIApiService service

type ApiResourceAPIUploadImageRequest struct {
	ctx context.Context
	ApiService *ResourceAPIApiService
	accessToken *string
	clientId *string
	openId *string
	body *string
}

// 访问令牌，用于标识用户和接口鉴权
func (r ApiResourceAPIUploadImageRequest) AccessToken(accessToken string) ApiResourceAPIUploadImageRequest {
	r.accessToken = &accessToken
	return r
}

// 应用 ID，用于标识应用和接口鉴权
func (r ApiResourceAPIUploadImageRequest) ClientId(clientId string) ApiResourceAPIUploadImageRequest {
	r.clientId = &clientId
	return r
}

// 开放平台用户 ID，用于标识用户和接口鉴权
func (r ApiResourceAPIUploadImageRequest) OpenId(openId string) ApiResourceAPIUploadImageRequest {
	r.openId = &openId
	return r
}

func (r ApiResourceAPIUploadImageRequest) Body(body string) ApiResourceAPIUploadImageRequest {
	r.body = &body
	return r
}

func (r ApiResourceAPIUploadImageRequest) Execute() (*ResourcesAPIResponse1, *http.Response, error) {
	return r.ApiService.ResourceAPIUploadImageExecute(r)
}

/*
ResourceAPIUploadImage Method for ResourceAPIUploadImage

Three required req headers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiResourceAPIUploadImageRequest
*/
func (a *ResourceAPIApiService) ResourceAPIUploadImage(ctx context.Context) ApiResourceAPIUploadImageRequest {
	return ApiResourceAPIUploadImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResourcesAPIResponse1
func (a *ResourceAPIApiService) ResourceAPIUploadImageExecute(r ApiResourceAPIUploadImageRequest) (*ResourcesAPIResponse1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourcesAPIResponse1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceAPIApiService.ResourceAPIUploadImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openapi/resources/v2/images"

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
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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