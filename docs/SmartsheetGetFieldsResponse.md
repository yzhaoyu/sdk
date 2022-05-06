# SmartsheetGetFieldsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**Next** | Pointer to **int32** |  | [optional] 
**Fields** | Pointer to [**[]SmartsheetField**](SmartsheetField.md) |  | [optional] 

## Methods

### NewSmartsheetGetFieldsResponse

`func NewSmartsheetGetFieldsResponse() *SmartsheetGetFieldsResponse`

NewSmartsheetGetFieldsResponse instantiates a new SmartsheetGetFieldsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetGetFieldsResponseWithDefaults

`func NewSmartsheetGetFieldsResponseWithDefaults() *SmartsheetGetFieldsResponse`

NewSmartsheetGetFieldsResponseWithDefaults instantiates a new SmartsheetGetFieldsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *SmartsheetGetFieldsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SmartsheetGetFieldsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SmartsheetGetFieldsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SmartsheetGetFieldsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetHasMore

`func (o *SmartsheetGetFieldsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *SmartsheetGetFieldsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *SmartsheetGetFieldsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *SmartsheetGetFieldsResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNext

`func (o *SmartsheetGetFieldsResponse) GetNext() int32`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SmartsheetGetFieldsResponse) GetNextOk() (*int32, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SmartsheetGetFieldsResponse) SetNext(v int32)`

SetNext sets Next field to given value.

### HasNext

`func (o *SmartsheetGetFieldsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetFields

`func (o *SmartsheetGetFieldsResponse) GetFields() []SmartsheetField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SmartsheetGetFieldsResponse) GetFieldsOk() (*[]SmartsheetField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SmartsheetGetFieldsResponse) SetFields(v []SmartsheetField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SmartsheetGetFieldsResponse) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


