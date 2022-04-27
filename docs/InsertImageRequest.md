# InsertImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageID** | **string** | 图片 ID | 
**Location** | [**Location**](Location.md) |  | 
**Width** | Pointer to **int32** | 图片宽度, 单位为 emu | [optional] 
**Height** | Pointer to **int32** | 图片高度, 单位为 emu | [optional] 
**AddonSource** | Pointer to **int32** | 业务来源 | [optional] 
**AddonID** | Pointer to **string** | 插件ID | [optional] 
**AnchorID** | Pointer to **string** | 文档插件元素的唯一 ID (6位的数字小写字母组合) | [optional] 
**ExtraData** | Pointer to **string** | 插件扩展字段 | [optional] 

## Methods

### NewInsertImageRequest

`func NewInsertImageRequest(imageID string, location Location, ) *InsertImageRequest`

NewInsertImageRequest instantiates a new InsertImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsertImageRequestWithDefaults

`func NewInsertImageRequestWithDefaults() *InsertImageRequest`

NewInsertImageRequestWithDefaults instantiates a new InsertImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageID

`func (o *InsertImageRequest) GetImageID() string`

GetImageID returns the ImageID field if non-nil, zero value otherwise.

### GetImageIDOk

`func (o *InsertImageRequest) GetImageIDOk() (*string, bool)`

GetImageIDOk returns a tuple with the ImageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageID

`func (o *InsertImageRequest) SetImageID(v string)`

SetImageID sets ImageID field to given value.


### GetLocation

`func (o *InsertImageRequest) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InsertImageRequest) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InsertImageRequest) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetWidth

`func (o *InsertImageRequest) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *InsertImageRequest) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *InsertImageRequest) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *InsertImageRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *InsertImageRequest) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *InsertImageRequest) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *InsertImageRequest) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *InsertImageRequest) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetAddonSource

`func (o *InsertImageRequest) GetAddonSource() int32`

GetAddonSource returns the AddonSource field if non-nil, zero value otherwise.

### GetAddonSourceOk

`func (o *InsertImageRequest) GetAddonSourceOk() (*int32, bool)`

GetAddonSourceOk returns a tuple with the AddonSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonSource

`func (o *InsertImageRequest) SetAddonSource(v int32)`

SetAddonSource sets AddonSource field to given value.

### HasAddonSource

`func (o *InsertImageRequest) HasAddonSource() bool`

HasAddonSource returns a boolean if a field has been set.

### GetAddonID

`func (o *InsertImageRequest) GetAddonID() string`

GetAddonID returns the AddonID field if non-nil, zero value otherwise.

### GetAddonIDOk

`func (o *InsertImageRequest) GetAddonIDOk() (*string, bool)`

GetAddonIDOk returns a tuple with the AddonID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonID

`func (o *InsertImageRequest) SetAddonID(v string)`

SetAddonID sets AddonID field to given value.

### HasAddonID

`func (o *InsertImageRequest) HasAddonID() bool`

HasAddonID returns a boolean if a field has been set.

### GetAnchorID

`func (o *InsertImageRequest) GetAnchorID() string`

GetAnchorID returns the AnchorID field if non-nil, zero value otherwise.

### GetAnchorIDOk

`func (o *InsertImageRequest) GetAnchorIDOk() (*string, bool)`

GetAnchorIDOk returns a tuple with the AnchorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorID

`func (o *InsertImageRequest) SetAnchorID(v string)`

SetAnchorID sets AnchorID field to given value.

### HasAnchorID

`func (o *InsertImageRequest) HasAnchorID() bool`

HasAnchorID returns a boolean if a field has been set.

### GetExtraData

`func (o *InsertImageRequest) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *InsertImageRequest) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *InsertImageRequest) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *InsertImageRequest) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


