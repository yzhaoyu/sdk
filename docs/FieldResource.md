# FieldResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldID** | Pointer to **string** |  | [optional] 
**FieldTitle** | Pointer to [**StringValue**](StringValue.md) |  | [optional] 
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

### NewFieldResource

`func NewFieldResource() *FieldResource`

NewFieldResource instantiates a new FieldResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldResourceWithDefaults

`func NewFieldResourceWithDefaults() *FieldResource`

NewFieldResourceWithDefaults instantiates a new FieldResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldID

`func (o *FieldResource) GetFieldID() string`

GetFieldID returns the FieldID field if non-nil, zero value otherwise.

### GetFieldIDOk

`func (o *FieldResource) GetFieldIDOk() (*string, bool)`

GetFieldIDOk returns a tuple with the FieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldID

`func (o *FieldResource) SetFieldID(v string)`

SetFieldID sets FieldID field to given value.

### HasFieldID

`func (o *FieldResource) HasFieldID() bool`

HasFieldID returns a boolean if a field has been set.

### GetFieldTitle

`func (o *FieldResource) GetFieldTitle() StringValue`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *FieldResource) GetFieldTitleOk() (*StringValue, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *FieldResource) SetFieldTitle(v StringValue)`

SetFieldTitle sets FieldTitle field to given value.

### HasFieldTitle

`func (o *FieldResource) HasFieldTitle() bool`

HasFieldTitle returns a boolean if a field has been set.

### GetFieldType

`func (o *FieldResource) GetFieldType() int32`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *FieldResource) GetFieldTypeOk() (*int32, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *FieldResource) SetFieldType(v int32)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *FieldResource) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetPropertyText

`func (o *FieldResource) GetPropertyText() map[string]interface{}`

GetPropertyText returns the PropertyText field if non-nil, zero value otherwise.

### GetPropertyTextOk

`func (o *FieldResource) GetPropertyTextOk() (*map[string]interface{}, bool)`

GetPropertyTextOk returns a tuple with the PropertyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyText

`func (o *FieldResource) SetPropertyText(v map[string]interface{})`

SetPropertyText sets PropertyText field to given value.

### HasPropertyText

`func (o *FieldResource) HasPropertyText() bool`

HasPropertyText returns a boolean if a field has been set.

### GetPropertyNumber

`func (o *FieldResource) GetPropertyNumber() NumberFieldProperty`

GetPropertyNumber returns the PropertyNumber field if non-nil, zero value otherwise.

### GetPropertyNumberOk

`func (o *FieldResource) GetPropertyNumberOk() (*NumberFieldProperty, bool)`

GetPropertyNumberOk returns a tuple with the PropertyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNumber

`func (o *FieldResource) SetPropertyNumber(v NumberFieldProperty)`

SetPropertyNumber sets PropertyNumber field to given value.

### HasPropertyNumber

`func (o *FieldResource) HasPropertyNumber() bool`

HasPropertyNumber returns a boolean if a field has been set.

### GetPropertyCheckbox

`func (o *FieldResource) GetPropertyCheckbox() CheckboxFieldProperty`

GetPropertyCheckbox returns the PropertyCheckbox field if non-nil, zero value otherwise.

### GetPropertyCheckboxOk

`func (o *FieldResource) GetPropertyCheckboxOk() (*CheckboxFieldProperty, bool)`

GetPropertyCheckboxOk returns a tuple with the PropertyCheckbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCheckbox

`func (o *FieldResource) SetPropertyCheckbox(v CheckboxFieldProperty)`

SetPropertyCheckbox sets PropertyCheckbox field to given value.

### HasPropertyCheckbox

`func (o *FieldResource) HasPropertyCheckbox() bool`

HasPropertyCheckbox returns a boolean if a field has been set.

### GetPropertyDateTime

`func (o *FieldResource) GetPropertyDateTime() DateTimeFieldProperty`

GetPropertyDateTime returns the PropertyDateTime field if non-nil, zero value otherwise.

### GetPropertyDateTimeOk

`func (o *FieldResource) GetPropertyDateTimeOk() (*DateTimeFieldProperty, bool)`

GetPropertyDateTimeOk returns a tuple with the PropertyDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDateTime

`func (o *FieldResource) SetPropertyDateTime(v DateTimeFieldProperty)`

SetPropertyDateTime sets PropertyDateTime field to given value.

### HasPropertyDateTime

`func (o *FieldResource) HasPropertyDateTime() bool`

HasPropertyDateTime returns a boolean if a field has been set.

### GetPropertyImage

`func (o *FieldResource) GetPropertyImage() map[string]interface{}`

GetPropertyImage returns the PropertyImage field if non-nil, zero value otherwise.

### GetPropertyImageOk

`func (o *FieldResource) GetPropertyImageOk() (*map[string]interface{}, bool)`

GetPropertyImageOk returns a tuple with the PropertyImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyImage

`func (o *FieldResource) SetPropertyImage(v map[string]interface{})`

SetPropertyImage sets PropertyImage field to given value.

### HasPropertyImage

`func (o *FieldResource) HasPropertyImage() bool`

HasPropertyImage returns a boolean if a field has been set.

### GetPropertyAttachment

`func (o *FieldResource) GetPropertyAttachment() map[string]interface{}`

GetPropertyAttachment returns the PropertyAttachment field if non-nil, zero value otherwise.

### GetPropertyAttachmentOk

`func (o *FieldResource) GetPropertyAttachmentOk() (*map[string]interface{}, bool)`

GetPropertyAttachmentOk returns a tuple with the PropertyAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyAttachment

`func (o *FieldResource) SetPropertyAttachment(v map[string]interface{})`

SetPropertyAttachment sets PropertyAttachment field to given value.

### HasPropertyAttachment

`func (o *FieldResource) HasPropertyAttachment() bool`

HasPropertyAttachment returns a boolean if a field has been set.

### GetPropertyUser

`func (o *FieldResource) GetPropertyUser() map[string]interface{}`

GetPropertyUser returns the PropertyUser field if non-nil, zero value otherwise.

### GetPropertyUserOk

`func (o *FieldResource) GetPropertyUserOk() (*map[string]interface{}, bool)`

GetPropertyUserOk returns a tuple with the PropertyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUser

`func (o *FieldResource) SetPropertyUser(v map[string]interface{})`

SetPropertyUser sets PropertyUser field to given value.

### HasPropertyUser

`func (o *FieldResource) HasPropertyUser() bool`

HasPropertyUser returns a boolean if a field has been set.

### GetPropertyUrl

`func (o *FieldResource) GetPropertyUrl() UrlFieldProperty`

GetPropertyUrl returns the PropertyUrl field if non-nil, zero value otherwise.

### GetPropertyUrlOk

`func (o *FieldResource) GetPropertyUrlOk() (*UrlFieldProperty, bool)`

GetPropertyUrlOk returns a tuple with the PropertyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUrl

`func (o *FieldResource) SetPropertyUrl(v UrlFieldProperty)`

SetPropertyUrl sets PropertyUrl field to given value.

### HasPropertyUrl

`func (o *FieldResource) HasPropertyUrl() bool`

HasPropertyUrl returns a boolean if a field has been set.

### GetPropertySelect

`func (o *FieldResource) GetPropertySelect() SelectFieldProperty`

GetPropertySelect returns the PropertySelect field if non-nil, zero value otherwise.

### GetPropertySelectOk

`func (o *FieldResource) GetPropertySelectOk() (*SelectFieldProperty, bool)`

GetPropertySelectOk returns a tuple with the PropertySelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySelect

`func (o *FieldResource) SetPropertySelect(v SelectFieldProperty)`

SetPropertySelect sets PropertySelect field to given value.

### HasPropertySelect

`func (o *FieldResource) HasPropertySelect() bool`

HasPropertySelect returns a boolean if a field has been set.

### GetPropertyCreatedUser

`func (o *FieldResource) GetPropertyCreatedUser() map[string]interface{}`

GetPropertyCreatedUser returns the PropertyCreatedUser field if non-nil, zero value otherwise.

### GetPropertyCreatedUserOk

`func (o *FieldResource) GetPropertyCreatedUserOk() (*map[string]interface{}, bool)`

GetPropertyCreatedUserOk returns a tuple with the PropertyCreatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCreatedUser

`func (o *FieldResource) SetPropertyCreatedUser(v map[string]interface{})`

SetPropertyCreatedUser sets PropertyCreatedUser field to given value.

### HasPropertyCreatedUser

`func (o *FieldResource) HasPropertyCreatedUser() bool`

HasPropertyCreatedUser returns a boolean if a field has been set.

### GetPropertyModifiedUser

`func (o *FieldResource) GetPropertyModifiedUser() map[string]interface{}`

GetPropertyModifiedUser returns the PropertyModifiedUser field if non-nil, zero value otherwise.

### GetPropertyModifiedUserOk

`func (o *FieldResource) GetPropertyModifiedUserOk() (*map[string]interface{}, bool)`

GetPropertyModifiedUserOk returns a tuple with the PropertyModifiedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyModifiedUser

`func (o *FieldResource) SetPropertyModifiedUser(v map[string]interface{})`

SetPropertyModifiedUser sets PropertyModifiedUser field to given value.

### HasPropertyModifiedUser

`func (o *FieldResource) HasPropertyModifiedUser() bool`

HasPropertyModifiedUser returns a boolean if a field has been set.

### GetPropertyCreatedTime

`func (o *FieldResource) GetPropertyCreatedTime() CreatedTimeFieldProperty`

GetPropertyCreatedTime returns the PropertyCreatedTime field if non-nil, zero value otherwise.

### GetPropertyCreatedTimeOk

`func (o *FieldResource) GetPropertyCreatedTimeOk() (*CreatedTimeFieldProperty, bool)`

GetPropertyCreatedTimeOk returns a tuple with the PropertyCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCreatedTime

`func (o *FieldResource) SetPropertyCreatedTime(v CreatedTimeFieldProperty)`

SetPropertyCreatedTime sets PropertyCreatedTime field to given value.

### HasPropertyCreatedTime

`func (o *FieldResource) HasPropertyCreatedTime() bool`

HasPropertyCreatedTime returns a boolean if a field has been set.

### GetPropertyModifiedTime

`func (o *FieldResource) GetPropertyModifiedTime() ModifiedTimeFieldProperty`

GetPropertyModifiedTime returns the PropertyModifiedTime field if non-nil, zero value otherwise.

### GetPropertyModifiedTimeOk

`func (o *FieldResource) GetPropertyModifiedTimeOk() (*ModifiedTimeFieldProperty, bool)`

GetPropertyModifiedTimeOk returns a tuple with the PropertyModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyModifiedTime

`func (o *FieldResource) SetPropertyModifiedTime(v ModifiedTimeFieldProperty)`

SetPropertyModifiedTime sets PropertyModifiedTime field to given value.

### HasPropertyModifiedTime

`func (o *FieldResource) HasPropertyModifiedTime() bool`

HasPropertyModifiedTime returns a boolean if a field has been set.

### GetPropertyProgress

`func (o *FieldResource) GetPropertyProgress() ProgressFieldProperty`

GetPropertyProgress returns the PropertyProgress field if non-nil, zero value otherwise.

### GetPropertyProgressOk

`func (o *FieldResource) GetPropertyProgressOk() (*ProgressFieldProperty, bool)`

GetPropertyProgressOk returns a tuple with the PropertyProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyProgress

`func (o *FieldResource) SetPropertyProgress(v ProgressFieldProperty)`

SetPropertyProgress sets PropertyProgress field to given value.

### HasPropertyProgress

`func (o *FieldResource) HasPropertyProgress() bool`

HasPropertyProgress returns a boolean if a field has been set.

### GetPropertyPhoneNumber

`func (o *FieldResource) GetPropertyPhoneNumber() map[string]interface{}`

GetPropertyPhoneNumber returns the PropertyPhoneNumber field if non-nil, zero value otherwise.

### GetPropertyPhoneNumberOk

`func (o *FieldResource) GetPropertyPhoneNumberOk() (*map[string]interface{}, bool)`

GetPropertyPhoneNumberOk returns a tuple with the PropertyPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyPhoneNumber

`func (o *FieldResource) SetPropertyPhoneNumber(v map[string]interface{})`

SetPropertyPhoneNumber sets PropertyPhoneNumber field to given value.

### HasPropertyPhoneNumber

`func (o *FieldResource) HasPropertyPhoneNumber() bool`

HasPropertyPhoneNumber returns a boolean if a field has been set.

### GetPropertyEmail

`func (o *FieldResource) GetPropertyEmail() map[string]interface{}`

GetPropertyEmail returns the PropertyEmail field if non-nil, zero value otherwise.

### GetPropertyEmailOk

`func (o *FieldResource) GetPropertyEmailOk() (*map[string]interface{}, bool)`

GetPropertyEmailOk returns a tuple with the PropertyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyEmail

`func (o *FieldResource) SetPropertyEmail(v map[string]interface{})`

SetPropertyEmail sets PropertyEmail field to given value.

### HasPropertyEmail

`func (o *FieldResource) HasPropertyEmail() bool`

HasPropertyEmail returns a boolean if a field has been set.

### GetPropertySingleSelect

`func (o *FieldResource) GetPropertySingleSelect() SingleSelectFieldProperty`

GetPropertySingleSelect returns the PropertySingleSelect field if non-nil, zero value otherwise.

### GetPropertySingleSelectOk

`func (o *FieldResource) GetPropertySingleSelectOk() (*SingleSelectFieldProperty, bool)`

GetPropertySingleSelectOk returns a tuple with the PropertySingleSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySingleSelect

`func (o *FieldResource) SetPropertySingleSelect(v SingleSelectFieldProperty)`

SetPropertySingleSelect sets PropertySingleSelect field to given value.

### HasPropertySingleSelect

`func (o *FieldResource) HasPropertySingleSelect() bool`

HasPropertySingleSelect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


