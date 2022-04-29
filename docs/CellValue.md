# CellValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to [**Value**](Value.md) |  | [optional] 
**ModifiedUser** | Pointer to **string** |  | [optional] 
**ModifiedTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewCellValue

`func NewCellValue() *CellValue`

NewCellValue instantiates a new CellValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellValueWithDefaults

`func NewCellValueWithDefaults() *CellValue`

NewCellValueWithDefaults instantiates a new CellValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CellValue) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CellValue) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CellValue) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CellValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *CellValue) GetValue() Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CellValue) GetValueOk() (*Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CellValue) SetValue(v Value)`

SetValue sets Value field to given value.

### HasValue

`func (o *CellValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetModifiedUser

`func (o *CellValue) GetModifiedUser() string`

GetModifiedUser returns the ModifiedUser field if non-nil, zero value otherwise.

### GetModifiedUserOk

`func (o *CellValue) GetModifiedUserOk() (*string, bool)`

GetModifiedUserOk returns a tuple with the ModifiedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedUser

`func (o *CellValue) SetModifiedUser(v string)`

SetModifiedUser sets ModifiedUser field to given value.

### HasModifiedUser

`func (o *CellValue) HasModifiedUser() bool`

HasModifiedUser returns a boolean if a field has been set.

### GetModifiedTime

`func (o *CellValue) GetModifiedTime() int64`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *CellValue) GetModifiedTimeOk() (*int64, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *CellValue) SetModifiedTime(v int64)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *CellValue) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


