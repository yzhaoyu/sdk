# InsertWebBlockRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**Location**](Location.md) |  | 
**AddonID** | **string** | 插件 ID | 
**AnchorID** | **string** | 文档插件元素的唯一ID(6位的数字小写字母组合) | 
**DisplayText** | Pointer to **string** | 展示文本 | [optional] 
**Width** | Pointer to **int32** | 插件宽度，单位为 emu | [optional] 
**Height** | Pointer to **int32** | 插件高度，单位为 emu | [optional] 
**OriginalURL** | **string** | 原始URL | 
**EmbedURL** | **string** | 插件展示的URL | 
**AddonType** | **int32** |  | 
**ExtraData** | Pointer to **string** | 插件扩展字段 | [optional] 
**LockAspectRatio** | Pointer to **bool** | 锁定宽高比, Default: true | [optional] 

## Methods

### NewInsertWebBlockRequest

`func NewInsertWebBlockRequest(location Location, addonID string, anchorID string, originalURL string, embedURL string, addonType int32, ) *InsertWebBlockRequest`

NewInsertWebBlockRequest instantiates a new InsertWebBlockRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsertWebBlockRequestWithDefaults

`func NewInsertWebBlockRequestWithDefaults() *InsertWebBlockRequest`

NewInsertWebBlockRequestWithDefaults instantiates a new InsertWebBlockRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *InsertWebBlockRequest) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InsertWebBlockRequest) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InsertWebBlockRequest) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetAddonID

`func (o *InsertWebBlockRequest) GetAddonID() string`

GetAddonID returns the AddonID field if non-nil, zero value otherwise.

### GetAddonIDOk

`func (o *InsertWebBlockRequest) GetAddonIDOk() (*string, bool)`

GetAddonIDOk returns a tuple with the AddonID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonID

`func (o *InsertWebBlockRequest) SetAddonID(v string)`

SetAddonID sets AddonID field to given value.


### GetAnchorID

`func (o *InsertWebBlockRequest) GetAnchorID() string`

GetAnchorID returns the AnchorID field if non-nil, zero value otherwise.

### GetAnchorIDOk

`func (o *InsertWebBlockRequest) GetAnchorIDOk() (*string, bool)`

GetAnchorIDOk returns a tuple with the AnchorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorID

`func (o *InsertWebBlockRequest) SetAnchorID(v string)`

SetAnchorID sets AnchorID field to given value.


### GetDisplayText

`func (o *InsertWebBlockRequest) GetDisplayText() string`

GetDisplayText returns the DisplayText field if non-nil, zero value otherwise.

### GetDisplayTextOk

`func (o *InsertWebBlockRequest) GetDisplayTextOk() (*string, bool)`

GetDisplayTextOk returns a tuple with the DisplayText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayText

`func (o *InsertWebBlockRequest) SetDisplayText(v string)`

SetDisplayText sets DisplayText field to given value.

### HasDisplayText

`func (o *InsertWebBlockRequest) HasDisplayText() bool`

HasDisplayText returns a boolean if a field has been set.

### GetWidth

`func (o *InsertWebBlockRequest) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *InsertWebBlockRequest) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *InsertWebBlockRequest) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *InsertWebBlockRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *InsertWebBlockRequest) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *InsertWebBlockRequest) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *InsertWebBlockRequest) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *InsertWebBlockRequest) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetOriginalURL

`func (o *InsertWebBlockRequest) GetOriginalURL() string`

GetOriginalURL returns the OriginalURL field if non-nil, zero value otherwise.

### GetOriginalURLOk

`func (o *InsertWebBlockRequest) GetOriginalURLOk() (*string, bool)`

GetOriginalURLOk returns a tuple with the OriginalURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalURL

`func (o *InsertWebBlockRequest) SetOriginalURL(v string)`

SetOriginalURL sets OriginalURL field to given value.


### GetEmbedURL

`func (o *InsertWebBlockRequest) GetEmbedURL() string`

GetEmbedURL returns the EmbedURL field if non-nil, zero value otherwise.

### GetEmbedURLOk

`func (o *InsertWebBlockRequest) GetEmbedURLOk() (*string, bool)`

GetEmbedURLOk returns a tuple with the EmbedURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedURL

`func (o *InsertWebBlockRequest) SetEmbedURL(v string)`

SetEmbedURL sets EmbedURL field to given value.


### GetAddonType

`func (o *InsertWebBlockRequest) GetAddonType() int32`

GetAddonType returns the AddonType field if non-nil, zero value otherwise.

### GetAddonTypeOk

`func (o *InsertWebBlockRequest) GetAddonTypeOk() (*int32, bool)`

GetAddonTypeOk returns a tuple with the AddonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonType

`func (o *InsertWebBlockRequest) SetAddonType(v int32)`

SetAddonType sets AddonType field to given value.


### GetExtraData

`func (o *InsertWebBlockRequest) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *InsertWebBlockRequest) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *InsertWebBlockRequest) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *InsertWebBlockRequest) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetLockAspectRatio

`func (o *InsertWebBlockRequest) GetLockAspectRatio() bool`

GetLockAspectRatio returns the LockAspectRatio field if non-nil, zero value otherwise.

### GetLockAspectRatioOk

`func (o *InsertWebBlockRequest) GetLockAspectRatioOk() (*bool, bool)`

GetLockAspectRatioOk returns a tuple with the LockAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockAspectRatio

`func (o *InsertWebBlockRequest) SetLockAspectRatio(v bool)`

SetLockAspectRatio sets LockAspectRatio field to given value.

### HasLockAspectRatio

`func (o *InsertWebBlockRequest) HasLockAspectRatio() bool`

HasLockAspectRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


