# Field

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**Property** | Pointer to [**FieldProperty**](FieldProperty.md) |  | [optional] 

## Methods

### NewField

`func NewField() *Field`

NewField instantiates a new Field object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldWithDefaults

`func NewFieldWithDefaults() *Field`

NewFieldWithDefaults instantiates a new Field object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Field) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Field) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Field) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Field) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *Field) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Field) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Field) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Field) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProperty

`func (o *Field) GetProperty() FieldProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *Field) GetPropertyOk() (*FieldProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *Field) SetProperty(v FieldProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *Field) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


