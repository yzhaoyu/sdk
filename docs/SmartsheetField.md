# SmartsheetField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldID** | Pointer to **string** |  | [optional] 
**FieldTitle** | Pointer to [**GoogleProtobufStringValue**](GoogleProtobufStringValue.md) |  | [optional] 
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

### NewSmartsheetField

`func NewSmartsheetField() *SmartsheetField`

NewSmartsheetField instantiates a new SmartsheetField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetFieldWithDefaults

`func NewSmartsheetFieldWithDefaults() *SmartsheetField`

NewSmartsheetFieldWithDefaults instantiates a new SmartsheetField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldID

`func (o *SmartsheetField) GetFieldID() string`

GetFieldID returns the FieldID field if non-nil, zero value otherwise.

### GetFieldIDOk

`func (o *SmartsheetField) GetFieldIDOk() (*string, bool)`

GetFieldIDOk returns a tuple with the FieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldID

`func (o *SmartsheetField) SetFieldID(v string)`

SetFieldID sets FieldID field to given value.

### HasFieldID

`func (o *SmartsheetField) HasFieldID() bool`

HasFieldID returns a boolean if a field has been set.

### GetFieldTitle

`func (o *SmartsheetField) GetFieldTitle() GoogleProtobufStringValue`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *SmartsheetField) GetFieldTitleOk() (*GoogleProtobufStringValue, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *SmartsheetField) SetFieldTitle(v GoogleProtobufStringValue)`

SetFieldTitle sets FieldTitle field to given value.

### HasFieldTitle

`func (o *SmartsheetField) HasFieldTitle() bool`

HasFieldTitle returns a boolean if a field has been set.

### GetFieldType

`func (o *SmartsheetField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *SmartsheetField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *SmartsheetField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *SmartsheetField) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetPropertyText

`func (o *SmartsheetField) GetPropertyText() map[string]interface{}`

GetPropertyText returns the PropertyText field if non-nil, zero value otherwise.

### GetPropertyTextOk

`func (o *SmartsheetField) GetPropertyTextOk() (*map[string]interface{}, bool)`

GetPropertyTextOk returns a tuple with the PropertyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyText

`func (o *SmartsheetField) SetPropertyText(v map[string]interface{})`

SetPropertyText sets PropertyText field to given value.

### HasPropertyText

`func (o *SmartsheetField) HasPropertyText() bool`

HasPropertyText returns a boolean if a field has been set.

### GetPropertyNumber

`func (o *SmartsheetField) GetPropertyNumber() SmartsheetNumberFieldProperty`

GetPropertyNumber returns the PropertyNumber field if non-nil, zero value otherwise.

### GetPropertyNumberOk

`func (o *SmartsheetField) GetPropertyNumberOk() (*SmartsheetNumberFieldProperty, bool)`

GetPropertyNumberOk returns a tuple with the PropertyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNumber

`func (o *SmartsheetField) SetPropertyNumber(v SmartsheetNumberFieldProperty)`

SetPropertyNumber sets PropertyNumber field to given value.

### HasPropertyNumber

`func (o *SmartsheetField) HasPropertyNumber() bool`

HasPropertyNumber returns a boolean if a field has been set.

### GetPropertyCheckbox

`func (o *SmartsheetField) GetPropertyCheckbox() SmartsheetCheckboxFieldProperty`

GetPropertyCheckbox returns the PropertyCheckbox field if non-nil, zero value otherwise.

### GetPropertyCheckboxOk

`func (o *SmartsheetField) GetPropertyCheckboxOk() (*SmartsheetCheckboxFieldProperty, bool)`

GetPropertyCheckboxOk returns a tuple with the PropertyCheckbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCheckbox

`func (o *SmartsheetField) SetPropertyCheckbox(v SmartsheetCheckboxFieldProperty)`

SetPropertyCheckbox sets PropertyCheckbox field to given value.

### HasPropertyCheckbox

`func (o *SmartsheetField) HasPropertyCheckbox() bool`

HasPropertyCheckbox returns a boolean if a field has been set.

### GetPropertyDateTime

`func (o *SmartsheetField) GetPropertyDateTime() SmartsheetDateTimeFieldProperty`

GetPropertyDateTime returns the PropertyDateTime field if non-nil, zero value otherwise.

### GetPropertyDateTimeOk

`func (o *SmartsheetField) GetPropertyDateTimeOk() (*SmartsheetDateTimeFieldProperty, bool)`

GetPropertyDateTimeOk returns a tuple with the PropertyDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDateTime

`func (o *SmartsheetField) SetPropertyDateTime(v SmartsheetDateTimeFieldProperty)`

SetPropertyDateTime sets PropertyDateTime field to given value.

### HasPropertyDateTime

`func (o *SmartsheetField) HasPropertyDateTime() bool`

HasPropertyDateTime returns a boolean if a field has been set.

### GetPropertyImage

`func (o *SmartsheetField) GetPropertyImage() map[string]interface{}`

GetPropertyImage returns the PropertyImage field if non-nil, zero value otherwise.

### GetPropertyImageOk

`func (o *SmartsheetField) GetPropertyImageOk() (*map[string]interface{}, bool)`

GetPropertyImageOk returns a tuple with the PropertyImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyImage

`func (o *SmartsheetField) SetPropertyImage(v map[string]interface{})`

SetPropertyImage sets PropertyImage field to given value.

### HasPropertyImage

`func (o *SmartsheetField) HasPropertyImage() bool`

HasPropertyImage returns a boolean if a field has been set.

### GetPropertyAttachment

`func (o *SmartsheetField) GetPropertyAttachment() map[string]interface{}`

GetPropertyAttachment returns the PropertyAttachment field if non-nil, zero value otherwise.

### GetPropertyAttachmentOk

`func (o *SmartsheetField) GetPropertyAttachmentOk() (*map[string]interface{}, bool)`

GetPropertyAttachmentOk returns a tuple with the PropertyAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyAttachment

`func (o *SmartsheetField) SetPropertyAttachment(v map[string]interface{})`

SetPropertyAttachment sets PropertyAttachment field to given value.

### HasPropertyAttachment

`func (o *SmartsheetField) HasPropertyAttachment() bool`

HasPropertyAttachment returns a boolean if a field has been set.

### GetPropertyUser

`func (o *SmartsheetField) GetPropertyUser() map[string]interface{}`

GetPropertyUser returns the PropertyUser field if non-nil, zero value otherwise.

### GetPropertyUserOk

`func (o *SmartsheetField) GetPropertyUserOk() (*map[string]interface{}, bool)`

GetPropertyUserOk returns a tuple with the PropertyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUser

`func (o *SmartsheetField) SetPropertyUser(v map[string]interface{})`

SetPropertyUser sets PropertyUser field to given value.

### HasPropertyUser

`func (o *SmartsheetField) HasPropertyUser() bool`

HasPropertyUser returns a boolean if a field has been set.

### GetPropertyUrl

`func (o *SmartsheetField) GetPropertyUrl() SmartsheetUrlFieldProperty`

GetPropertyUrl returns the PropertyUrl field if non-nil, zero value otherwise.

### GetPropertyUrlOk

`func (o *SmartsheetField) GetPropertyUrlOk() (*SmartsheetUrlFieldProperty, bool)`

GetPropertyUrlOk returns a tuple with the PropertyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUrl

`func (o *SmartsheetField) SetPropertyUrl(v SmartsheetUrlFieldProperty)`

SetPropertyUrl sets PropertyUrl field to given value.

### HasPropertyUrl

`func (o *SmartsheetField) HasPropertyUrl() bool`

HasPropertyUrl returns a boolean if a field has been set.

### GetPropertySelect

`func (o *SmartsheetField) GetPropertySelect() SmartsheetSelectFieldProperty`

GetPropertySelect returns the PropertySelect field if non-nil, zero value otherwise.

### GetPropertySelectOk

`func (o *SmartsheetField) GetPropertySelectOk() (*SmartsheetSelectFieldProperty, bool)`

GetPropertySelectOk returns a tuple with the PropertySelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySelect

`func (o *SmartsheetField) SetPropertySelect(v SmartsheetSelectFieldProperty)`

SetPropertySelect sets PropertySelect field to given value.

### HasPropertySelect

`func (o *SmartsheetField) HasPropertySelect() bool`

HasPropertySelect returns a boolean if a field has been set.

### GetPropertyCreatedUser

`func (o *SmartsheetField) GetPropertyCreatedUser() map[string]interface{}`

GetPropertyCreatedUser returns the PropertyCreatedUser field if non-nil, zero value otherwise.

### GetPropertyCreatedUserOk

`func (o *SmartsheetField) GetPropertyCreatedUserOk() (*map[string]interface{}, bool)`

GetPropertyCreatedUserOk returns a tuple with the PropertyCreatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCreatedUser

`func (o *SmartsheetField) SetPropertyCreatedUser(v map[string]interface{})`

SetPropertyCreatedUser sets PropertyCreatedUser field to given value.

### HasPropertyCreatedUser

`func (o *SmartsheetField) HasPropertyCreatedUser() bool`

HasPropertyCreatedUser returns a boolean if a field has been set.

### GetPropertyModifiedUser

`func (o *SmartsheetField) GetPropertyModifiedUser() map[string]interface{}`

GetPropertyModifiedUser returns the PropertyModifiedUser field if non-nil, zero value otherwise.

### GetPropertyModifiedUserOk

`func (o *SmartsheetField) GetPropertyModifiedUserOk() (*map[string]interface{}, bool)`

GetPropertyModifiedUserOk returns a tuple with the PropertyModifiedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyModifiedUser

`func (o *SmartsheetField) SetPropertyModifiedUser(v map[string]interface{})`

SetPropertyModifiedUser sets PropertyModifiedUser field to given value.

### HasPropertyModifiedUser

`func (o *SmartsheetField) HasPropertyModifiedUser() bool`

HasPropertyModifiedUser returns a boolean if a field has been set.

### GetPropertyCreatedTime

`func (o *SmartsheetField) GetPropertyCreatedTime() SmartsheetCreatedTimeFieldProperty`

GetPropertyCreatedTime returns the PropertyCreatedTime field if non-nil, zero value otherwise.

### GetPropertyCreatedTimeOk

`func (o *SmartsheetField) GetPropertyCreatedTimeOk() (*SmartsheetCreatedTimeFieldProperty, bool)`

GetPropertyCreatedTimeOk returns a tuple with the PropertyCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCreatedTime

`func (o *SmartsheetField) SetPropertyCreatedTime(v SmartsheetCreatedTimeFieldProperty)`

SetPropertyCreatedTime sets PropertyCreatedTime field to given value.

### HasPropertyCreatedTime

`func (o *SmartsheetField) HasPropertyCreatedTime() bool`

HasPropertyCreatedTime returns a boolean if a field has been set.

### GetPropertyModifiedTime

`func (o *SmartsheetField) GetPropertyModifiedTime() SmartsheetModifiedTimeFieldProperty`

GetPropertyModifiedTime returns the PropertyModifiedTime field if non-nil, zero value otherwise.

### GetPropertyModifiedTimeOk

`func (o *SmartsheetField) GetPropertyModifiedTimeOk() (*SmartsheetModifiedTimeFieldProperty, bool)`

GetPropertyModifiedTimeOk returns a tuple with the PropertyModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyModifiedTime

`func (o *SmartsheetField) SetPropertyModifiedTime(v SmartsheetModifiedTimeFieldProperty)`

SetPropertyModifiedTime sets PropertyModifiedTime field to given value.

### HasPropertyModifiedTime

`func (o *SmartsheetField) HasPropertyModifiedTime() bool`

HasPropertyModifiedTime returns a boolean if a field has been set.

### GetPropertyProgress

`func (o *SmartsheetField) GetPropertyProgress() SmartsheetProgressFieldProperty`

GetPropertyProgress returns the PropertyProgress field if non-nil, zero value otherwise.

### GetPropertyProgressOk

`func (o *SmartsheetField) GetPropertyProgressOk() (*SmartsheetProgressFieldProperty, bool)`

GetPropertyProgressOk returns a tuple with the PropertyProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyProgress

`func (o *SmartsheetField) SetPropertyProgress(v SmartsheetProgressFieldProperty)`

SetPropertyProgress sets PropertyProgress field to given value.

### HasPropertyProgress

`func (o *SmartsheetField) HasPropertyProgress() bool`

HasPropertyProgress returns a boolean if a field has been set.

### GetPropertyPhoneNumber

`func (o *SmartsheetField) GetPropertyPhoneNumber() map[string]interface{}`

GetPropertyPhoneNumber returns the PropertyPhoneNumber field if non-nil, zero value otherwise.

### GetPropertyPhoneNumberOk

`func (o *SmartsheetField) GetPropertyPhoneNumberOk() (*map[string]interface{}, bool)`

GetPropertyPhoneNumberOk returns a tuple with the PropertyPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyPhoneNumber

`func (o *SmartsheetField) SetPropertyPhoneNumber(v map[string]interface{})`

SetPropertyPhoneNumber sets PropertyPhoneNumber field to given value.

### HasPropertyPhoneNumber

`func (o *SmartsheetField) HasPropertyPhoneNumber() bool`

HasPropertyPhoneNumber returns a boolean if a field has been set.

### GetPropertyEmail

`func (o *SmartsheetField) GetPropertyEmail() map[string]interface{}`

GetPropertyEmail returns the PropertyEmail field if non-nil, zero value otherwise.

### GetPropertyEmailOk

`func (o *SmartsheetField) GetPropertyEmailOk() (*map[string]interface{}, bool)`

GetPropertyEmailOk returns a tuple with the PropertyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyEmail

`func (o *SmartsheetField) SetPropertyEmail(v map[string]interface{})`

SetPropertyEmail sets PropertyEmail field to given value.

### HasPropertyEmail

`func (o *SmartsheetField) HasPropertyEmail() bool`

HasPropertyEmail returns a boolean if a field has been set.

### GetPropertySingleSelect

`func (o *SmartsheetField) GetPropertySingleSelect() SmartsheetSingleSelectFieldProperty`

GetPropertySingleSelect returns the PropertySingleSelect field if non-nil, zero value otherwise.

### GetPropertySingleSelectOk

`func (o *SmartsheetField) GetPropertySingleSelectOk() (*SmartsheetSingleSelectFieldProperty, bool)`

GetPropertySingleSelectOk returns a tuple with the PropertySingleSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySingleSelect

`func (o *SmartsheetField) SetPropertySingleSelect(v SmartsheetSingleSelectFieldProperty)`

SetPropertySingleSelect sets PropertySingleSelect field to given value.

### HasPropertySingleSelect

`func (o *SmartsheetField) HasPropertySingleSelect() bool`

HasPropertySingleSelect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


