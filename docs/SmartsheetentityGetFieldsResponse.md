# SmartsheetentityGetFieldsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**Next** | Pointer to **int32** |  | [optional] 
**Fields** | Pointer to [**[]SmartsheetentityField**](SmartsheetentityField.md) |  | [optional] 

## Methods

### NewSmartsheetentityGetFieldsResponse

`func NewSmartsheetentityGetFieldsResponse() *SmartsheetentityGetFieldsResponse`

NewSmartsheetentityGetFieldsResponse instantiates a new SmartsheetentityGetFieldsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetentityGetFieldsResponseWithDefaults

`func NewSmartsheetentityGetFieldsResponseWithDefaults() *SmartsheetentityGetFieldsResponse`

NewSmartsheetentityGetFieldsResponseWithDefaults instantiates a new SmartsheetentityGetFieldsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *SmartsheetentityGetFieldsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SmartsheetentityGetFieldsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SmartsheetentityGetFieldsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SmartsheetentityGetFieldsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetHasMore

`func (o *SmartsheetentityGetFieldsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *SmartsheetentityGetFieldsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *SmartsheetentityGetFieldsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *SmartsheetentityGetFieldsResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNext

`func (o *SmartsheetentityGetFieldsResponse) GetNext() int32`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SmartsheetentityGetFieldsResponse) GetNextOk() (*int32, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SmartsheetentityGetFieldsResponse) SetNext(v int32)`

SetNext sets Next field to given value.

### HasNext

`func (o *SmartsheetentityGetFieldsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetFields

`func (o *SmartsheetentityGetFieldsResponse) GetFields() []SmartsheetentityField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SmartsheetentityGetFieldsResponse) GetFieldsOk() (*[]SmartsheetentityField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SmartsheetentityGetFieldsResponse) SetFields(v []SmartsheetentityField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SmartsheetentityGetFieldsResponse) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


