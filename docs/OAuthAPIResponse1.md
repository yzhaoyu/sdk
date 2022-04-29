# OAuthAPIResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ret** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**GetTokenRsp**](GetTokenRsp.md) |  | [optional] 

## Methods

### NewOAuthAPIResponse1

`func NewOAuthAPIResponse1() *OAuthAPIResponse1`

NewOAuthAPIResponse1 instantiates a new OAuthAPIResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthAPIResponse1WithDefaults

`func NewOAuthAPIResponse1WithDefaults() *OAuthAPIResponse1`

NewOAuthAPIResponse1WithDefaults instantiates a new OAuthAPIResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRet

`func (o *OAuthAPIResponse1) GetRet() int32`

GetRet returns the Ret field if non-nil, zero value otherwise.

### GetRetOk

`func (o *OAuthAPIResponse1) GetRetOk() (*int32, bool)`

GetRetOk returns a tuple with the Ret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRet

`func (o *OAuthAPIResponse1) SetRet(v int32)`

SetRet sets Ret field to given value.

### HasRet

`func (o *OAuthAPIResponse1) HasRet() bool`

HasRet returns a boolean if a field has been set.

### GetMsg

`func (o *OAuthAPIResponse1) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OAuthAPIResponse1) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OAuthAPIResponse1) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OAuthAPIResponse1) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *OAuthAPIResponse1) GetData() GetTokenRsp`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OAuthAPIResponse1) GetDataOk() (*GetTokenRsp, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OAuthAPIResponse1) SetData(v GetTokenRsp)`

SetData sets Data field to given value.

### HasData

`func (o *OAuthAPIResponse1) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


