# SmartsheetGetRecordsRsp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**Next** | Pointer to **int32** |  | [optional] 
**Records** | Pointer to [**[]SmartsheetRecord**](SmartsheetRecord.md) |  | [optional] 

## Methods

### NewSmartsheetGetRecordsRsp

`func NewSmartsheetGetRecordsRsp() *SmartsheetGetRecordsRsp`

NewSmartsheetGetRecordsRsp instantiates a new SmartsheetGetRecordsRsp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetGetRecordsRspWithDefaults

`func NewSmartsheetGetRecordsRspWithDefaults() *SmartsheetGetRecordsRsp`

NewSmartsheetGetRecordsRspWithDefaults instantiates a new SmartsheetGetRecordsRsp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *SmartsheetGetRecordsRsp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SmartsheetGetRecordsRsp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SmartsheetGetRecordsRsp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SmartsheetGetRecordsRsp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetHasMore

`func (o *SmartsheetGetRecordsRsp) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *SmartsheetGetRecordsRsp) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *SmartsheetGetRecordsRsp) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *SmartsheetGetRecordsRsp) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNext

`func (o *SmartsheetGetRecordsRsp) GetNext() int32`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SmartsheetGetRecordsRsp) GetNextOk() (*int32, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SmartsheetGetRecordsRsp) SetNext(v int32)`

SetNext sets Next field to given value.

### HasNext

`func (o *SmartsheetGetRecordsRsp) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetRecords

`func (o *SmartsheetGetRecordsRsp) GetRecords() []SmartsheetRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *SmartsheetGetRecordsRsp) GetRecordsOk() (*[]SmartsheetRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *SmartsheetGetRecordsRsp) SetRecords(v []SmartsheetRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *SmartsheetGetRecordsRsp) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


