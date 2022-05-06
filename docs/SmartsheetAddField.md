# SmartsheetAddField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldTitle** | Pointer to **string** |  | [optional] 
**FieldType** | Pointer to **string** |  | [optional] 
**PropertyText** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyNumber** | Pointer to [**SmartsheetNumberFieldProperty**](SmartsheetNumberFieldProperty.md) |  | [optional] 
**PropertyCheckbox** | Pointer to [**SmartsheetCheckboxFieldProperty**](SmartsheetCheckboxFieldProperty.md) |  | [optional] 
**PropertyDateTime** | Pointer to [**SmartsheetDateTimeFieldProperty**](SmartsheetDateTimeFieldProperty.md) |  | [optional] 
**PropertyImage** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyAttachment** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyUser** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyUrl** | Pointer to [**SmartsheetUrlFieldProperty**](SmartsheetUrlFieldProperty.md) |  | [optional] 
**PropertySelect** | Pointer to [**SmartsheetSelectFieldProperty**](SmartsheetSelectFieldProperty.md) |  | [optional] 
**PropertyCreatedUser** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyModifiedUser** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyCreatedTime** | Pointer to [**SmartsheetCreatedTimeFieldProperty**](SmartsheetCreatedTimeFieldProperty.md) |  | [optional] 
**PropertyModifiedTime** | Pointer to [**SmartsheetModifiedTimeFieldProperty**](SmartsheetModifiedTimeFieldProperty.md) |  | [optional] 
**PropertyProgress** | Pointer to [**SmartsheetProgressFieldProperty**](SmartsheetProgressFieldProperty.md) |  | [optional] 
**PropertyPhoneNumber** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertyEmail** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertySingleSelect** | Pointer to [**SmartsheetSingleSelectFieldProperty**](SmartsheetSingleSelectFieldProperty.md) |  | [optional] 

## Methods

### NewSmartsheetAddField

`func NewSmartsheetAddField() *SmartsheetAddField`

NewSmartsheetAddField instantiates a new SmartsheetAddField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetAddFieldWithDefaults

`func NewSmartsheetAddFieldWithDefaults() *SmartsheetAddField`

NewSmartsheetAddFieldWithDefaults instantiates a new SmartsheetAddField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldTitle

`func (o *SmartsheetAddField) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *SmartsheetAddField) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *SmartsheetAddField) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.

### HasFieldTitle

`func (o *SmartsheetAddField) HasFieldTitle() bool`

HasFieldTitle returns a boolean if a field has been set.

### GetFieldType

`func (o *SmartsheetAddField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *SmartsheetAddField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *SmartsheetAddField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *SmartsheetAddField) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetPropertyText

`func (o *SmartsheetAddField) GetPropertyText() map[string]interface{}`

GetPropertyText returns the PropertyText field if non-nil, zero value otherwise.

### GetPropertyTextOk

`func (o *SmartsheetAddField) GetPropertyTextOk() (*map[string]interface{}, bool)`

GetPropertyTextOk returns a tuple with the PropertyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyText

`func (o *SmartsheetAddField) SetPropertyText(v map[string]interface{})`

SetPropertyText sets PropertyText field to given value.

### HasPropertyText

`func (o *SmartsheetAddField) HasPropertyText() bool`

HasPropertyText returns a boolean if a field has been set.

### GetPropertyNumber

`func (o *SmartsheetAddField) GetPropertyNumber() SmartsheetNumberFieldProperty`

GetPropertyNumber returns the PropertyNumber field if non-nil, zero value otherwise.

### GetPropertyNumberOk

`func (o *SmartsheetAddField) GetPropertyNumberOk() (*SmartsheetNumberFieldProperty, bool)`

GetPropertyNumberOk returns a tuple with the PropertyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNumber

`func (o *SmartsheetAddField) SetPropertyNumber(v SmartsheetNumberFieldProperty)`

SetPropertyNumber sets PropertyNumber field to given value.

### HasPropertyNumber

`func (o *SmartsheetAddField) HasPropertyNumber() bool`

HasPropertyNumber returns a boolean if a field has been set.

### GetPropertyCheckbox

`func (o *SmartsheetAddField) GetPropertyCheckbox() SmartsheetCheckboxFieldProperty`

GetPropertyCheckbox returns the PropertyCheckbox field if non-nil, zero value otherwise.

### GetPropertyCheckboxOk

`func (o *SmartsheetAddField) GetPropertyCheckboxOk() (*SmartsheetCheckboxFieldProperty, bool)`

GetPropertyCheckboxOk returns a tuple with the PropertyCheckbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCheckbox

`func (o *SmartsheetAddField) SetPropertyCheckbox(v SmartsheetCheckboxFieldProperty)`

SetPropertyCheckbox sets PropertyCheckbox field to given value.

### HasPropertyCheckbox

`func (o *SmartsheetAddField) HasPropertyCheckbox() bool`

HasPropertyCheckbox returns a boolean if a field has been set.

### GetPropertyDateTime

`func (o *SmartsheetAddField) GetPropertyDateTime() SmartsheetDateTimeFieldProperty`

GetPropertyDateTime returns the PropertyDateTime field if non-nil, zero value otherwise.

### GetPropertyDateTimeOk

`func (o *SmartsheetAddField) GetPropertyDateTimeOk() (*SmartsheetDateTimeFieldProperty, bool)`

GetPropertyDateTimeOk returns a tuple with the PropertyDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDateTime

`func (o *SmartsheetAddField) SetPropertyDateTime(v SmartsheetDateTimeFieldProperty)`

SetPropertyDateTime sets PropertyDateTime field to given value.

### HasPropertyDateTime

`func (o *SmartsheetAddField) HasPropertyDateTime() bool`

HasPropertyDateTime returns a boolean if a field has been set.

### GetPropertyImage

`func (o *SmartsheetAddField) GetPropertyImage() map[string]interface{}`

GetPropertyImage returns the PropertyImage field if non-nil, zero value otherwise.

### GetPropertyImageOk

`func (o *SmartsheetAddField) GetPropertyImageOk() (*map[string]interface{}, bool)`

GetPropertyImageOk returns a tuple with the PropertyImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyImage

`func (o *SmartsheetAddField) SetPropertyImage(v map[string]interface{})`

SetPropertyImage sets PropertyImage field to given value.

### HasPropertyImage

`func (o *SmartsheetAddField) HasPropertyImage() bool`

HasPropertyImage returns a boolean if a field has been set.

### GetPropertyAttachment

`func (o *SmartsheetAddField) GetPropertyAttachment() map[string]interface{}`

GetPropertyAttachment returns the PropertyAttachment field if non-nil, zero value otherwise.

### GetPropertyAttachmentOk

`func (o *SmartsheetAddField) GetPropertyAttachmentOk() (*map[string]interface{}, bool)`

GetPropertyAttachmentOk returns a tuple with the PropertyAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyAttachment

`func (o *SmartsheetAddField) SetPropertyAttachment(v map[string]interface{})`

SetPropertyAttachment sets PropertyAttachment field to given value.

### HasPropertyAttachment

`func (o *SmartsheetAddField) HasPropertyAttachment() bool`

HasPropertyAttachment returns a boolean if a field has been set.

### GetPropertyUser

`func (o *SmartsheetAddField) GetPropertyUser() map[string]interface{}`

GetPropertyUser returns the PropertyUser field if non-nil, zero value otherwise.

### GetPropertyUserOk

`func (o *SmartsheetAddField) GetPropertyUserOk() (*map[string]interface{}, bool)`

GetPropertyUserOk returns a tuple with the PropertyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUser

`func (o *SmartsheetAddField) SetPropertyUser(v map[string]interface{})`

SetPropertyUser sets PropertyUser field to given value.

### HasPropertyUser

`func (o *SmartsheetAddField) HasPropertyUser() bool`

HasPropertyUser returns a boolean if a field has been set.

### GetPropertyUrl

`func (o *SmartsheetAddField) GetPropertyUrl() SmartsheetUrlFieldProperty`

GetPropertyUrl returns the PropertyUrl field if non-nil, zero value otherwise.

### GetPropertyUrlOk

`func (o *SmartsheetAddField) GetPropertyUrlOk() (*SmartsheetUrlFieldProperty, bool)`

GetPropertyUrlOk returns a tuple with the PropertyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUrl

`func (o *SmartsheetAddField) SetPropertyUrl(v SmartsheetUrlFieldProperty)`

SetPropertyUrl sets PropertyUrl field to given value.

### HasPropertyUrl

`func (o *SmartsheetAddField) HasPropertyUrl() bool`

HasPropertyUrl returns a boolean if a field has been set.

### GetPropertySelect

`func (o *SmartsheetAddField) GetPropertySelect() SmartsheetSelectFieldProperty`

GetPropertySelect returns the PropertySelect field if non-nil, zero value otherwise.

### GetPropertySelectOk

`func (o *SmartsheetAddField) GetPropertySelectOk() (*SmartsheetSelectFieldProperty, bool)`

GetPropertySelectOk returns a tuple with the PropertySelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySelect

`func (o *SmartsheetAddField) SetPropertySelect(v SmartsheetSelectFieldProperty)`

SetPropertySelect sets PropertySelect field to given value.

### HasPropertySelect

`func (o *SmartsheetAddField) HasPropertySelect() bool`

HasPropertySelect returns a boolean if a field has been set.

### GetPropertyCreatedUser

`func (o *SmartsheetAddField) GetPropertyCreatedUser() map[string]interface{}`

GetPropertyCreatedUser returns the PropertyCreatedUser field if non-nil, zero value otherwise.

### GetPropertyCreatedUserOk

`func (o *SmartsheetAddField) GetPropertyCreatedUserOk() (*map[string]interface{}, bool)`

GetPropertyCreatedUserOk returns a tuple with the PropertyCreatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCreatedUser

`func (o *SmartsheetAddField) SetPropertyCreatedUser(v map[string]interface{})`

SetPropertyCreatedUser sets PropertyCreatedUser field to given value.

### HasPropertyCreatedUser

`func (o *SmartsheetAddField) HasPropertyCreatedUser() bool`

HasPropertyCreatedUser returns a boolean if a field has been set.

### GetPropertyModifiedUser

`func (o *SmartsheetAddField) GetPropertyModifiedUser() map[string]interface{}`

GetPropertyModifiedUser returns the PropertyModifiedUser field if non-nil, zero value otherwise.

### GetPropertyModifiedUserOk

`func (o *SmartsheetAddField) GetPropertyModifiedUserOk() (*map[string]interface{}, bool)`

GetPropertyModifiedUserOk returns a tuple with the PropertyModifiedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyModifiedUser

`func (o *SmartsheetAddField) SetPropertyModifiedUser(v map[string]interface{})`

SetPropertyModifiedUser sets PropertyModifiedUser field to given value.

### HasPropertyModifiedUser

`func (o *SmartsheetAddField) HasPropertyModifiedUser() bool`

HasPropertyModifiedUser returns a boolean if a field has been set.

### GetPropertyCreatedTime

`func (o *SmartsheetAddField) GetPropertyCreatedTime() SmartsheetCreatedTimeFieldProperty`

GetPropertyCreatedTime returns the PropertyCreatedTime field if non-nil, zero value otherwise.

### GetPropertyCreatedTimeOk

`func (o *SmartsheetAddField) GetPropertyCreatedTimeOk() (*SmartsheetCreatedTimeFieldProperty, bool)`

GetPropertyCreatedTimeOk returns a tuple with the PropertyCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCreatedTime

`func (o *SmartsheetAddField) SetPropertyCreatedTime(v SmartsheetCreatedTimeFieldProperty)`

SetPropertyCreatedTime sets PropertyCreatedTime field to given value.

### HasPropertyCreatedTime

`func (o *SmartsheetAddField) HasPropertyCreatedTime() bool`

HasPropertyCreatedTime returns a boolean if a field has been set.

### GetPropertyModifiedTime

`func (o *SmartsheetAddField) GetPropertyModifiedTime() SmartsheetModifiedTimeFieldProperty`

GetPropertyModifiedTime returns the PropertyModifiedTime field if non-nil, zero value otherwise.

### GetPropertyModifiedTimeOk

`func (o *SmartsheetAddField) GetPropertyModifiedTimeOk() (*SmartsheetModifiedTimeFieldProperty, bool)`

GetPropertyModifiedTimeOk returns a tuple with the PropertyModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyModifiedTime

`func (o *SmartsheetAddField) SetPropertyModifiedTime(v SmartsheetModifiedTimeFieldProperty)`

SetPropertyModifiedTime sets PropertyModifiedTime field to given value.

### HasPropertyModifiedTime

`func (o *SmartsheetAddField) HasPropertyModifiedTime() bool`

HasPropertyModifiedTime returns a boolean if a field has been set.

### GetPropertyProgress

`func (o *SmartsheetAddField) GetPropertyProgress() SmartsheetProgressFieldProperty`

GetPropertyProgress returns the PropertyProgress field if non-nil, zero value otherwise.

### GetPropertyProgressOk

`func (o *SmartsheetAddField) GetPropertyProgressOk() (*SmartsheetProgressFieldProperty, bool)`

GetPropertyProgressOk returns a tuple with the PropertyProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyProgress

`func (o *SmartsheetAddField) SetPropertyProgress(v SmartsheetProgressFieldProperty)`

SetPropertyProgress sets PropertyProgress field to given value.

### HasPropertyProgress

`func (o *SmartsheetAddField) HasPropertyProgress() bool`

HasPropertyProgress returns a boolean if a field has been set.

### GetPropertyPhoneNumber

`func (o *SmartsheetAddField) GetPropertyPhoneNumber() map[string]interface{}`

GetPropertyPhoneNumber returns the PropertyPhoneNumber field if non-nil, zero value otherwise.

### GetPropertyPhoneNumberOk

`func (o *SmartsheetAddField) GetPropertyPhoneNumberOk() (*map[string]interface{}, bool)`

GetPropertyPhoneNumberOk returns a tuple with the PropertyPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyPhoneNumber

`func (o *SmartsheetAddField) SetPropertyPhoneNumber(v map[string]interface{})`

SetPropertyPhoneNumber sets PropertyPhoneNumber field to given value.

### HasPropertyPhoneNumber

`func (o *SmartsheetAddField) HasPropertyPhoneNumber() bool`

HasPropertyPhoneNumber returns a boolean if a field has been set.

### GetPropertyEmail

`func (o *SmartsheetAddField) GetPropertyEmail() map[string]interface{}`

GetPropertyEmail returns the PropertyEmail field if non-nil, zero value otherwise.

### GetPropertyEmailOk

`func (o *SmartsheetAddField) GetPropertyEmailOk() (*map[string]interface{}, bool)`

GetPropertyEmailOk returns a tuple with the PropertyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyEmail

`func (o *SmartsheetAddField) SetPropertyEmail(v map[string]interface{})`

SetPropertyEmail sets PropertyEmail field to given value.

### HasPropertyEmail

`func (o *SmartsheetAddField) HasPropertyEmail() bool`

HasPropertyEmail returns a boolean if a field has been set.

### GetPropertySingleSelect

`func (o *SmartsheetAddField) GetPropertySingleSelect() SmartsheetSingleSelectFieldProperty`

GetPropertySingleSelect returns the PropertySingleSelect field if non-nil, zero value otherwise.

### GetPropertySingleSelectOk

`func (o *SmartsheetAddField) GetPropertySingleSelectOk() (*SmartsheetSingleSelectFieldProperty, bool)`

GetPropertySingleSelectOk returns a tuple with the PropertySingleSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySingleSelect

`func (o *SmartsheetAddField) SetPropertySingleSelect(v SmartsheetSingleSelectFieldProperty)`

SetPropertySingleSelect sets PropertySingleSelect field to given value.

### HasPropertySingleSelect

`func (o *SmartsheetAddField) HasPropertySingleSelect() bool`

HasPropertySingleSelect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


