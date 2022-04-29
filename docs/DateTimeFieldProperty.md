# DateTimeFieldProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] 
**AutoFill** | Pointer to **bool** |  | [optional] 

## Methods

### NewDateTimeFieldProperty

`func NewDateTimeFieldProperty() *DateTimeFieldProperty`

NewDateTimeFieldProperty instantiates a new DateTimeFieldProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeFieldPropertyWithDefaults

`func NewDateTimeFieldPropertyWithDefaults() *DateTimeFieldProperty`

NewDateTimeFieldPropertyWithDefaults instantiates a new DateTimeFieldProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *DateTimeFieldProperty) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DateTimeFieldProperty) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DateTimeFieldProperty) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DateTimeFieldProperty) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetAutoFill

`func (o *DateTimeFieldProperty) GetAutoFill() bool`

GetAutoFill returns the AutoFill field if non-nil, zero value otherwise.

### GetAutoFillOk

`func (o *DateTimeFieldProperty) GetAutoFillOk() (*bool, bool)`

GetAutoFillOk returns a tuple with the AutoFill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFill

`func (o *DateTimeFieldProperty) SetAutoFill(v bool)`

SetAutoFill sets AutoFill field to given value.

### HasAutoFill

`func (o *DateTimeFieldProperty) HasAutoFill() bool`

HasAutoFill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


