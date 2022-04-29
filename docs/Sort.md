# Sort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldTitle** | Pointer to **string** |  | [optional] 
**Desc** | Pointer to **bool** |  | [optional] 

## Methods

### NewSort

`func NewSort() *Sort`

NewSort instantiates a new Sort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortWithDefaults

`func NewSortWithDefaults() *Sort`

NewSortWithDefaults instantiates a new Sort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldTitle

`func (o *Sort) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *Sort) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *Sort) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.

### HasFieldTitle

`func (o *Sort) HasFieldTitle() bool`

HasFieldTitle returns a boolean if a field has been set.

### GetDesc

`func (o *Sort) GetDesc() bool`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Sort) GetDescOk() (*bool, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Sort) SetDesc(v bool)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Sort) HasDesc() bool`

HasDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


