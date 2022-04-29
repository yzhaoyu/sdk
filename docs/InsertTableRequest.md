# InsertTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Rows** | Pointer to **int32** |  | [optional] 
**Cols** | Pointer to **int32** |  | [optional] 

## Methods

### NewInsertTableRequest

`func NewInsertTableRequest() *InsertTableRequest`

NewInsertTableRequest instantiates a new InsertTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsertTableRequestWithDefaults

`func NewInsertTableRequestWithDefaults() *InsertTableRequest`

NewInsertTableRequestWithDefaults instantiates a new InsertTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *InsertTableRequest) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InsertTableRequest) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InsertTableRequest) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InsertTableRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRows

`func (o *InsertTableRequest) GetRows() int32`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *InsertTableRequest) GetRowsOk() (*int32, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *InsertTableRequest) SetRows(v int32)`

SetRows sets Rows field to given value.

### HasRows

`func (o *InsertTableRequest) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetCols

`func (o *InsertTableRequest) GetCols() int32`

GetCols returns the Cols field if non-nil, zero value otherwise.

### GetColsOk

`func (o *InsertTableRequest) GetColsOk() (*int32, bool)`

GetColsOk returns a tuple with the Cols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCols

`func (o *InsertTableRequest) SetCols(v int32)`

SetCols sets Cols field to given value.

### HasCols

`func (o *InsertTableRequest) HasCols() bool`

HasCols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


