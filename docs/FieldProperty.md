# FieldProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyText** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyNumber** | Pointer to [**NumberFieldProperty**](NumberFieldProperty.md) |  | [optional] 
**PropertyCheckbox** | Pointer to [**CheckboxFieldProperty**](CheckboxFieldProperty.md) |  | [optional] 
**PropertyDateTime** | Pointer to [**DateTimeFieldProperty**](DateTimeFieldProperty.md) |  | [optional] 
**PropertyImage** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyAttachment** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyUser** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyUrl** | Pointer to [**UrlFieldProperty**](UrlFieldProperty.md) |  | [optional] 
**PropertySelect** | Pointer to [**SelectFieldProperty**](SelectFieldProperty.md) |  | [optional] 
**PropertyCreatedUser** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyModifiedUser** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyCreatedTime** | Pointer to [**CreatedTimeFieldProperty**](CreatedTimeFieldProperty.md) |  | [optional] 
**PropertyModifiedTime** | Pointer to [**ModifiedTimeFieldProperty**](ModifiedTimeFieldProperty.md) |  | [optional] 
**PropertyProgress** | Pointer to [**ProgressFieldProperty**](ProgressFieldProperty.md) |  | [optional] 
**PropertyPhoneNumber** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyEmail** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertySingleSelect** | Pointer to [**SingleSelectFieldProperty**](SingleSelectFieldProperty.md) |  | [optional] 

## Methods

### NewFieldProperty

`func NewFieldProperty() *FieldProperty`

NewFieldProperty instantiates a new FieldProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldPropertyWithDefaults

`func NewFieldPropertyWithDefaults() *FieldProperty`

NewFieldPropertyWithDefaults instantiates a new FieldProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyText

`func (o *FieldProperty) GetPropertyText() map[string]interface{}`

GetPropertyText returns the PropertyText field if non-nil, zero value otherwise.

### GetPropertyTextOk

`func (o *FieldProperty) GetPropertyTextOk() (*map[string]interface{}, bool)`

GetPropertyTextOk returns a tuple with the PropertyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyText

`func (o *FieldProperty) SetPropertyText(v map[string]interface{})`

SetPropertyText sets PropertyText field to given value.

### HasPropertyText

`func (o *FieldProperty) HasPropertyText() bool`

HasPropertyText returns a boolean if a field has been set.

### GetPropertyNumber

`func (o *FieldProperty) GetPropertyNumber() NumberFieldProperty`

GetPropertyNumber returns the PropertyNumber field if non-nil, zero value otherwise.

### GetPropertyNumberOk

`func (o *FieldProperty) GetPropertyNumberOk() (*NumberFieldProperty, bool)`

GetPropertyNumberOk returns a tuple with the PropertyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNumber

`func (o *FieldProperty) SetPropertyNumber(v NumberFieldProperty)`

SetPropertyNumber sets PropertyNumber field to given value.

### HasPropertyNumber

`func (o *FieldProperty) HasPropertyNumber() bool`

HasPropertyNumber returns a boolean if a field has been set.

### GetPropertyCheckbox

`func (o *FieldProperty) GetPropertyCheckbox() CheckboxFieldProperty`

GetPropertyCheckbox returns the PropertyCheckbox field if non-nil, zero value otherwise.

### GetPropertyCheckboxOk

`func (o *FieldProperty) GetPropertyCheckboxOk() (*CheckboxFieldProperty, bool)`

GetPropertyCheckboxOk returns a tuple with the PropertyCheckbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCheckbox

`func (o *FieldProperty) SetPropertyCheckbox(v CheckboxFieldProperty)`

SetPropertyCheckbox sets PropertyCheckbox field to given value.

### HasPropertyCheckbox

`func (o *FieldProperty) HasPropertyCheckbox() bool`

HasPropertyCheckbox returns a boolean if a field has been set.

### GetPropertyDateTime

`func (o *FieldProperty) GetPropertyDateTime() DateTimeFieldProperty`

GetPropertyDateTime returns the PropertyDateTime field if non-nil, zero value otherwise.

### GetPropertyDateTimeOk

`func (o *FieldProperty) GetPropertyDateTimeOk() (*DateTimeFieldProperty, bool)`

GetPropertyDateTimeOk returns a tuple with the PropertyDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDateTime

`func (o *FieldProperty) SetPropertyDateTime(v DateTimeFieldProperty)`

SetPropertyDateTime sets PropertyDateTime field to given value.

### HasPropertyDateTime

`func (o *FieldProperty) HasPropertyDateTime() bool`

HasPropertyDateTime returns a boolean if a field has been set.

### GetPropertyImage

`func (o *FieldProperty) GetPropertyImage() map[string]interface{}`

GetPropertyImage returns the PropertyImage field if non-nil, zero value otherwise.

### GetPropertyImageOk

`func (o *FieldProperty) GetPropertyImageOk() (*map[string]interface{}, bool)`

GetPropertyImageOk returns a tuple with the PropertyImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyImage

`func (o *FieldProperty) SetPropertyImage(v map[string]interface{})`

SetPropertyImage sets PropertyImage field to given value.

### HasPropertyImage

`func (o *FieldProperty) HasPropertyImage() bool`

HasPropertyImage returns a boolean if a field has been set.

### GetPropertyAttachment

`func (o *FieldProperty) GetPropertyAttachment() map[string]interface{}`

GetPropertyAttachment returns the PropertyAttachment field if non-nil, zero value otherwise.

### GetPropertyAttachmentOk

`func (o *FieldProperty) GetPropertyAttachmentOk() (*map[string]interface{}, bool)`

GetPropertyAttachmentOk returns a tuple with the PropertyAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyAttachment

`func (o *FieldProperty) SetPropertyAttachment(v map[string]interface{})`

SetPropertyAttachment sets PropertyAttachment field to given value.

### HasPropertyAttachment

`func (o *FieldProperty) HasPropertyAttachment() bool`

HasPropertyAttachment returns a boolean if a field has been set.

### GetPropertyUser

`func (o *FieldProperty) GetPropertyUser() map[string]interface{}`

GetPropertyUser returns the PropertyUser field if non-nil, zero value otherwise.

### GetPropertyUserOk

`func (o *FieldProperty) GetPropertyUserOk() (*map[string]interface{}, bool)`

GetPropertyUserOk returns a tuple with the PropertyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUser

`func (o *FieldProperty) SetPropertyUser(v map[string]interface{})`

SetPropertyUser sets PropertyUser field to given value.

### HasPropertyUser

`func (o *FieldProperty) HasPropertyUser() bool`

HasPropertyUser returns a boolean if a field has been set.

### GetPropertyUrl

`func (o *FieldProperty) GetPropertyUrl() UrlFieldProperty`

GetPropertyUrl returns the PropertyUrl field if non-nil, zero value otherwise.

### GetPropertyUrlOk

`func (o *FieldProperty) GetPropertyUrlOk() (*UrlFieldProperty, bool)`

GetPropertyUrlOk returns a tuple with the PropertyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUrl

`func (o *FieldProperty) SetPropertyUrl(v UrlFieldProperty)`

SetPropertyUrl sets PropertyUrl field to given value.

### HasPropertyUrl

`func (o *FieldProperty) HasPropertyUrl() bool`

HasPropertyUrl returns a boolean if a field has been set.

### GetPropertySelect

`func (o *FieldProperty) GetPropertySelect() SelectFieldProperty`

GetPropertySelect returns the PropertySelect field if non-nil, zero value otherwise.

### GetPropertySelectOk

`func (o *FieldProperty) GetPropertySelectOk() (*SelectFieldProperty, bool)`

GetPropertySelectOk returns a tuple with the PropertySelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySelect

`func (o *FieldProperty) SetPropertySelect(v SelectFieldProperty)`

SetPropertySelect sets PropertySelect field to given value.

### HasPropertySelect

`func (o *FieldProperty) HasPropertySelect() bool`

HasPropertySelect returns a boolean if a field has been set.

### GetPropertyCreatedUser

`func (o *FieldProperty) GetPropertyCreatedUser() map[string]interface{}`

GetPropertyCreatedUser returns the PropertyCreatedUser field if non-nil, zero value otherwise.

### GetPropertyCreatedUserOk

`func (o *FieldProperty) GetPropertyCreatedUserOk() (*map[string]interface{}, bool)`

GetPropertyCreatedUserOk returns a tuple with the PropertyCreatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCreatedUser

`func (o *FieldProperty) SetPropertyCreatedUser(v map[string]interface{})`

SetPropertyCreatedUser sets PropertyCreatedUser field to given value.

### HasPropertyCreatedUser

`func (o *FieldProperty) HasPropertyCreatedUser() bool`

HasPropertyCreatedUser returns a boolean if a field has been set.

### GetPropertyModifiedUser

`func (o *FieldProperty) GetPropertyModifiedUser() map[string]interface{}`

GetPropertyModifiedUser returns the PropertyModifiedUser field if non-nil, zero value otherwise.

### GetPropertyModifiedUserOk

`func (o *FieldProperty) GetPropertyModifiedUserOk() (*map[string]interface{}, bool)`

GetPropertyModifiedUserOk returns a tuple with the PropertyModifiedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyModifiedUser

`func (o *FieldProperty) SetPropertyModifiedUser(v map[string]interface{})`

SetPropertyModifiedUser sets PropertyModifiedUser field to given value.

### HasPropertyModifiedUser

`func (o *FieldProperty) HasPropertyModifiedUser() bool`

HasPropertyModifiedUser returns a boolean if a field has been set.

### GetPropertyCreatedTime

`func (o *FieldProperty) GetPropertyCreatedTime() CreatedTimeFieldProperty`

GetPropertyCreatedTime returns the PropertyCreatedTime field if non-nil, zero value otherwise.

### GetPropertyCreatedTimeOk

`func (o *FieldProperty) GetPropertyCreatedTimeOk() (*CreatedTimeFieldProperty, bool)`

GetPropertyCreatedTimeOk returns a tuple with the PropertyCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCreatedTime

`func (o *FieldProperty) SetPropertyCreatedTime(v CreatedTimeFieldProperty)`

SetPropertyCreatedTime sets PropertyCreatedTime field to given value.

### HasPropertyCreatedTime

`func (o *FieldProperty) HasPropertyCreatedTime() bool`

HasPropertyCreatedTime returns a boolean if a field has been set.

### GetPropertyModifiedTime

`func (o *FieldProperty) GetPropertyModifiedTime() ModifiedTimeFieldProperty`

GetPropertyModifiedTime returns the PropertyModifiedTime field if non-nil, zero value otherwise.

### GetPropertyModifiedTimeOk

`func (o *FieldProperty) GetPropertyModifiedTimeOk() (*ModifiedTimeFieldProperty, bool)`

GetPropertyModifiedTimeOk returns a tuple with the PropertyModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyModifiedTime

`func (o *FieldProperty) SetPropertyModifiedTime(v ModifiedTimeFieldProperty)`

SetPropertyModifiedTime sets PropertyModifiedTime field to given value.

### HasPropertyModifiedTime

`func (o *FieldProperty) HasPropertyModifiedTime() bool`

HasPropertyModifiedTime returns a boolean if a field has been set.

### GetPropertyProgress

`func (o *FieldProperty) GetPropertyProgress() ProgressFieldProperty`

GetPropertyProgress returns the PropertyProgress field if non-nil, zero value otherwise.

### GetPropertyProgressOk

`func (o *FieldProperty) GetPropertyProgressOk() (*ProgressFieldProperty, bool)`

GetPropertyProgressOk returns a tuple with the PropertyProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyProgress

`func (o *FieldProperty) SetPropertyProgress(v ProgressFieldProperty)`

SetPropertyProgress sets PropertyProgress field to given value.

### HasPropertyProgress

`func (o *FieldProperty) HasPropertyProgress() bool`

HasPropertyProgress returns a boolean if a field has been set.

### GetPropertyPhoneNumber

`func (o *FieldProperty) GetPropertyPhoneNumber() map[string]interface{}`

GetPropertyPhoneNumber returns the PropertyPhoneNumber field if non-nil, zero value otherwise.

### GetPropertyPhoneNumberOk

`func (o *FieldProperty) GetPropertyPhoneNumberOk() (*map[string]interface{}, bool)`

GetPropertyPhoneNumberOk returns a tuple with the PropertyPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyPhoneNumber

`func (o *FieldProperty) SetPropertyPhoneNumber(v map[string]interface{})`

SetPropertyPhoneNumber sets PropertyPhoneNumber field to given value.

### HasPropertyPhoneNumber

`func (o *FieldProperty) HasPropertyPhoneNumber() bool`

HasPropertyPhoneNumber returns a boolean if a field has been set.

### GetPropertyEmail

`func (o *FieldProperty) GetPropertyEmail() map[string]interface{}`

GetPropertyEmail returns the PropertyEmail field if non-nil, zero value otherwise.

### GetPropertyEmailOk

`func (o *FieldProperty) GetPropertyEmailOk() (*map[string]interface{}, bool)`

GetPropertyEmailOk returns a tuple with the PropertyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyEmail

`func (o *FieldProperty) SetPropertyEmail(v map[string]interface{})`

SetPropertyEmail sets PropertyEmail field to given value.

### HasPropertyEmail

`func (o *FieldProperty) HasPropertyEmail() bool`

HasPropertyEmail returns a boolean if a field has been set.

### GetPropertySingleSelect

`func (o *FieldProperty) GetPropertySingleSelect() SingleSelectFieldProperty`

GetPropertySingleSelect returns the PropertySingleSelect field if non-nil, zero value otherwise.

### GetPropertySingleSelectOk

`func (o *FieldProperty) GetPropertySingleSelectOk() (*SingleSelectFieldProperty, bool)`

GetPropertySingleSelectOk returns a tuple with the PropertySingleSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySingleSelect

`func (o *FieldProperty) SetPropertySingleSelect(v SingleSelectFieldProperty)`

SetPropertySingleSelect sets PropertySingleSelect field to given value.

### HasPropertySingleSelect

`func (o *FieldProperty) HasPropertySingleSelect() bool`

HasPropertySingleSelect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


