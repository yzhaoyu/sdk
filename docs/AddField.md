# AddField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldTitle** | Pointer to **string** |  | [optional] 
**FieldType** | Pointer to **int32** |  | [optional] 
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

### NewAddField

`func NewAddField() *AddField`

NewAddField instantiates a new AddField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFieldWithDefaults

`func NewAddFieldWithDefaults() *AddField`

NewAddFieldWithDefaults instantiates a new AddField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldTitle

`func (o *AddField) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *AddField) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *AddField) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.

### HasFieldTitle

`func (o *AddField) HasFieldTitle() bool`

HasFieldTitle returns a boolean if a field has been set.

### GetFieldType

`func (o *AddField) GetFieldType() int32`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *AddField) GetFieldTypeOk() (*int32, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *AddField) SetFieldType(v int32)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *AddField) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetPropertyText

`func (o *AddField) GetPropertyText() map[string]interface{}`

GetPropertyText returns the PropertyText field if non-nil, zero value otherwise.

### GetPropertyTextOk

`func (o *AddField) GetPropertyTextOk() (*map[string]interface{}, bool)`

GetPropertyTextOk returns a tuple with the PropertyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyText

`func (o *AddField) SetPropertyText(v map[string]interface{})`

SetPropertyText sets PropertyText field to given value.

### HasPropertyText

`func (o *AddField) HasPropertyText() bool`

HasPropertyText returns a boolean if a field has been set.

### GetPropertyNumber

`func (o *AddField) GetPropertyNumber() NumberFieldProperty`

GetPropertyNumber returns the PropertyNumber field if non-nil, zero value otherwise.

### GetPropertyNumberOk

`func (o *AddField) GetPropertyNumberOk() (*NumberFieldProperty, bool)`

GetPropertyNumberOk returns a tuple with the PropertyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNumber

`func (o *AddField) SetPropertyNumber(v NumberFieldProperty)`

SetPropertyNumber sets PropertyNumber field to given value.

### HasPropertyNumber

`func (o *AddField) HasPropertyNumber() bool`

HasPropertyNumber returns a boolean if a field has been set.

### GetPropertyCheckbox

`func (o *AddField) GetPropertyCheckbox() CheckboxFieldProperty`

GetPropertyCheckbox returns the PropertyCheckbox field if non-nil, zero value otherwise.

### GetPropertyCheckboxOk

`func (o *AddField) GetPropertyCheckboxOk() (*CheckboxFieldProperty, bool)`

GetPropertyCheckboxOk returns a tuple with the PropertyCheckbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCheckbox

`func (o *AddField) SetPropertyCheckbox(v CheckboxFieldProperty)`

SetPropertyCheckbox sets PropertyCheckbox field to given value.

### HasPropertyCheckbox

`func (o *AddField) HasPropertyCheckbox() bool`

HasPropertyCheckbox returns a boolean if a field has been set.

### GetPropertyDateTime

`func (o *AddField) GetPropertyDateTime() DateTimeFieldProperty`

GetPropertyDateTime returns the PropertyDateTime field if non-nil, zero value otherwise.

### GetPropertyDateTimeOk

`func (o *AddField) GetPropertyDateTimeOk() (*DateTimeFieldProperty, bool)`

GetPropertyDateTimeOk returns a tuple with the PropertyDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDateTime

`func (o *AddField) SetPropertyDateTime(v DateTimeFieldProperty)`

SetPropertyDateTime sets PropertyDateTime field to given value.

### HasPropertyDateTime

`func (o *AddField) HasPropertyDateTime() bool`

HasPropertyDateTime returns a boolean if a field has been set.

### GetPropertyImage

`func (o *AddField) GetPropertyImage() map[string]interface{}`

GetPropertyImage returns the PropertyImage field if non-nil, zero value otherwise.

### GetPropertyImageOk

`func (o *AddField) GetPropertyImageOk() (*map[string]interface{}, bool)`

GetPropertyImageOk returns a tuple with the PropertyImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyImage

`func (o *AddField) SetPropertyImage(v map[string]interface{})`

SetPropertyImage sets PropertyImage field to given value.

### HasPropertyImage

`func (o *AddField) HasPropertyImage() bool`

HasPropertyImage returns a boolean if a field has been set.

### GetPropertyAttachment

`func (o *AddField) GetPropertyAttachment() map[string]interface{}`

GetPropertyAttachment returns the PropertyAttachment field if non-nil, zero value otherwise.

### GetPropertyAttachmentOk

`func (o *AddField) GetPropertyAttachmentOk() (*map[string]interface{}, bool)`

GetPropertyAttachmentOk returns a tuple with the PropertyAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyAttachment

`func (o *AddField) SetPropertyAttachment(v map[string]interface{})`

SetPropertyAttachment sets PropertyAttachment field to given value.

### HasPropertyAttachment

`func (o *AddField) HasPropertyAttachment() bool`

HasPropertyAttachment returns a boolean if a field has been set.

### GetPropertyUser

`func (o *AddField) GetPropertyUser() map[string]interface{}`

GetPropertyUser returns the PropertyUser field if non-nil, zero value otherwise.

### GetPropertyUserOk

`func (o *AddField) GetPropertyUserOk() (*map[string]interface{}, bool)`

GetPropertyUserOk returns a tuple with the PropertyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUser

`func (o *AddField) SetPropertyUser(v map[string]interface{})`

SetPropertyUser sets PropertyUser field to given value.

### HasPropertyUser

`func (o *AddField) HasPropertyUser() bool`

HasPropertyUser returns a boolean if a field has been set.

### GetPropertyUrl

`func (o *AddField) GetPropertyUrl() UrlFieldProperty`

GetPropertyUrl returns the PropertyUrl field if non-nil, zero value otherwise.

### GetPropertyUrlOk

`func (o *AddField) GetPropertyUrlOk() (*UrlFieldProperty, bool)`

GetPropertyUrlOk returns a tuple with the PropertyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUrl

`func (o *AddField) SetPropertyUrl(v UrlFieldProperty)`

SetPropertyUrl sets PropertyUrl field to given value.

### HasPropertyUrl

`func (o *AddField) HasPropertyUrl() bool`

HasPropertyUrl returns a boolean if a field has been set.

### GetPropertySelect

`func (o *AddField) GetPropertySelect() SelectFieldProperty`

GetPropertySelect returns the PropertySelect field if non-nil, zero value otherwise.

### GetPropertySelectOk

`func (o *AddField) GetPropertySelectOk() (*SelectFieldProperty, bool)`

GetPropertySelectOk returns a tuple with the PropertySelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySelect

`func (o *AddField) SetPropertySelect(v SelectFieldProperty)`

SetPropertySelect sets PropertySelect field to given value.

### HasPropertySelect

`func (o *AddField) HasPropertySelect() bool`

HasPropertySelect returns a boolean if a field has been set.

### GetPropertyCreatedUser

`func (o *AddField) GetPropertyCreatedUser() map[string]interface{}`

GetPropertyCreatedUser returns the PropertyCreatedUser field if non-nil, zero value otherwise.

### GetPropertyCreatedUserOk

`func (o *AddField) GetPropertyCreatedUserOk() (*map[string]interface{}, bool)`

GetPropertyCreatedUserOk returns a tuple with the PropertyCreatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCreatedUser

`func (o *AddField) SetPropertyCreatedUser(v map[string]interface{})`

SetPropertyCreatedUser sets PropertyCreatedUser field to given value.

### HasPropertyCreatedUser

`func (o *AddField) HasPropertyCreatedUser() bool`

HasPropertyCreatedUser returns a boolean if a field has been set.

### GetPropertyModifiedUser

`func (o *AddField) GetPropertyModifiedUser() map[string]interface{}`

GetPropertyModifiedUser returns the PropertyModifiedUser field if non-nil, zero value otherwise.

### GetPropertyModifiedUserOk

`func (o *AddField) GetPropertyModifiedUserOk() (*map[string]interface{}, bool)`

GetPropertyModifiedUserOk returns a tuple with the PropertyModifiedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyModifiedUser

`func (o *AddField) SetPropertyModifiedUser(v map[string]interface{})`

SetPropertyModifiedUser sets PropertyModifiedUser field to given value.

### HasPropertyModifiedUser

`func (o *AddField) HasPropertyModifiedUser() bool`

HasPropertyModifiedUser returns a boolean if a field has been set.

### GetPropertyCreatedTime

`func (o *AddField) GetPropertyCreatedTime() CreatedTimeFieldProperty`

GetPropertyCreatedTime returns the PropertyCreatedTime field if non-nil, zero value otherwise.

### GetPropertyCreatedTimeOk

`func (o *AddField) GetPropertyCreatedTimeOk() (*CreatedTimeFieldProperty, bool)`

GetPropertyCreatedTimeOk returns a tuple with the PropertyCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCreatedTime

`func (o *AddField) SetPropertyCreatedTime(v CreatedTimeFieldProperty)`

SetPropertyCreatedTime sets PropertyCreatedTime field to given value.

### HasPropertyCreatedTime

`func (o *AddField) HasPropertyCreatedTime() bool`

HasPropertyCreatedTime returns a boolean if a field has been set.

### GetPropertyModifiedTime

`func (o *AddField) GetPropertyModifiedTime() ModifiedTimeFieldProperty`

GetPropertyModifiedTime returns the PropertyModifiedTime field if non-nil, zero value otherwise.

### GetPropertyModifiedTimeOk

`func (o *AddField) GetPropertyModifiedTimeOk() (*ModifiedTimeFieldProperty, bool)`

GetPropertyModifiedTimeOk returns a tuple with the PropertyModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyModifiedTime

`func (o *AddField) SetPropertyModifiedTime(v ModifiedTimeFieldProperty)`

SetPropertyModifiedTime sets PropertyModifiedTime field to given value.

### HasPropertyModifiedTime

`func (o *AddField) HasPropertyModifiedTime() bool`

HasPropertyModifiedTime returns a boolean if a field has been set.

### GetPropertyProgress

`func (o *AddField) GetPropertyProgress() ProgressFieldProperty`

GetPropertyProgress returns the PropertyProgress field if non-nil, zero value otherwise.

### GetPropertyProgressOk

`func (o *AddField) GetPropertyProgressOk() (*ProgressFieldProperty, bool)`

GetPropertyProgressOk returns a tuple with the PropertyProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyProgress

`func (o *AddField) SetPropertyProgress(v ProgressFieldProperty)`

SetPropertyProgress sets PropertyProgress field to given value.

### HasPropertyProgress

`func (o *AddField) HasPropertyProgress() bool`

HasPropertyProgress returns a boolean if a field has been set.

### GetPropertyPhoneNumber

`func (o *AddField) GetPropertyPhoneNumber() map[string]interface{}`

GetPropertyPhoneNumber returns the PropertyPhoneNumber field if non-nil, zero value otherwise.

### GetPropertyPhoneNumberOk

`func (o *AddField) GetPropertyPhoneNumberOk() (*map[string]interface{}, bool)`

GetPropertyPhoneNumberOk returns a tuple with the PropertyPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyPhoneNumber

`func (o *AddField) SetPropertyPhoneNumber(v map[string]interface{})`

SetPropertyPhoneNumber sets PropertyPhoneNumber field to given value.

### HasPropertyPhoneNumber

`func (o *AddField) HasPropertyPhoneNumber() bool`

HasPropertyPhoneNumber returns a boolean if a field has been set.

### GetPropertyEmail

`func (o *AddField) GetPropertyEmail() map[string]interface{}`

GetPropertyEmail returns the PropertyEmail field if non-nil, zero value otherwise.

### GetPropertyEmailOk

`func (o *AddField) GetPropertyEmailOk() (*map[string]interface{}, bool)`

GetPropertyEmailOk returns a tuple with the PropertyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyEmail

`func (o *AddField) SetPropertyEmail(v map[string]interface{})`

SetPropertyEmail sets PropertyEmail field to given value.

### HasPropertyEmail

`func (o *AddField) HasPropertyEmail() bool`

HasPropertyEmail returns a boolean if a field has been set.

### GetPropertySingleSelect

`func (o *AddField) GetPropertySingleSelect() SingleSelectFieldProperty`

GetPropertySingleSelect returns the PropertySingleSelect field if non-nil, zero value otherwise.

### GetPropertySingleSelectOk

`func (o *AddField) GetPropertySingleSelectOk() (*SingleSelectFieldProperty, bool)`

GetPropertySingleSelectOk returns a tuple with the PropertySingleSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySingleSelect

`func (o *AddField) SetPropertySingleSelect(v SingleSelectFieldProperty)`

SetPropertySingleSelect sets PropertySingleSelect field to given value.

### HasPropertySingleSelect

`func (o *AddField) HasPropertySingleSelect() bool`

HasPropertySingleSelect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


