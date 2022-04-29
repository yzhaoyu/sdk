# GetUserInfoRsp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenID** | Pointer to **string** |  | [optional] 
**Nick** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUserInfoRsp

`func NewGetUserInfoRsp() *GetUserInfoRsp`

NewGetUserInfoRsp instantiates a new GetUserInfoRsp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserInfoRspWithDefaults

`func NewGetUserInfoRspWithDefaults() *GetUserInfoRsp`

NewGetUserInfoRspWithDefaults instantiates a new GetUserInfoRsp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenID

`func (o *GetUserInfoRsp) GetOpenID() string`

GetOpenID returns the OpenID field if non-nil, zero value otherwise.

### GetOpenIDOk

`func (o *GetUserInfoRsp) GetOpenIDOk() (*string, bool)`

GetOpenIDOk returns a tuple with the OpenID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenID

`func (o *GetUserInfoRsp) SetOpenID(v string)`

SetOpenID sets OpenID field to given value.

### HasOpenID

`func (o *GetUserInfoRsp) HasOpenID() bool`

HasOpenID returns a boolean if a field has been set.

### GetNick

`func (o *GetUserInfoRsp) GetNick() string`

GetNick returns the Nick field if non-nil, zero value otherwise.

### GetNickOk

`func (o *GetUserInfoRsp) GetNickOk() (*string, bool)`

GetNickOk returns a tuple with the Nick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNick

`func (o *GetUserInfoRsp) SetNick(v string)`

SetNick sets Nick field to given value.

### HasNick

`func (o *GetUserInfoRsp) HasNick() bool`

HasNick returns a boolean if a field has been set.

### GetAvatar

`func (o *GetUserInfoRsp) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *GetUserInfoRsp) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *GetUserInfoRsp) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *GetUserInfoRsp) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetSource

`func (o *GetUserInfoRsp) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetUserInfoRsp) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetUserInfoRsp) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetUserInfoRsp) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


