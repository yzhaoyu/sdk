# OAuthAPIResponse3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ret** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**RefreshTokenRsp**](RefreshTokenRsp.md) |  | [optional] 

## Methods

### NewOAuthAPIResponse3

`func NewOAuthAPIResponse3() *OAuthAPIResponse3`

NewOAuthAPIResponse3 instantiates a new OAuthAPIResponse3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthAPIResponse3WithDefaults

`func NewOAuthAPIResponse3WithDefaults() *OAuthAPIResponse3`

NewOAuthAPIResponse3WithDefaults instantiates a new OAuthAPIResponse3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRet

`func (o *OAuthAPIResponse3) GetRet() int32`

GetRet returns the Ret field if non-nil, zero value otherwise.

### GetRetOk

`func (o *OAuthAPIResponse3) GetRetOk() (*int32, bool)`

GetRetOk returns a tuple with the Ret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRet

`func (o *OAuthAPIResponse3) SetRet(v int32)`

SetRet sets Ret field to given value.

### HasRet

`func (o *OAuthAPIResponse3) HasRet() bool`

HasRet returns a boolean if a field has been set.

### GetMsg

`func (o *OAuthAPIResponse3) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OAuthAPIResponse3) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OAuthAPIResponse3) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OAuthAPIResponse3) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *OAuthAPIResponse3) GetData() RefreshTokenRsp`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OAuthAPIResponse3) GetDataOk() (*RefreshTokenRsp, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OAuthAPIResponse3) SetData(v RefreshTokenRsp)`

SetData sets Data field to given value.

### HasData

`func (o *OAuthAPIResponse3) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


