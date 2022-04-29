# GetRecordsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**Next** | Pointer to **int32** |  | [optional] 
**Records** | Pointer to [**[]RecordResource**](RecordResource.md) |  | [optional] 

## Methods

### NewGetRecordsResponse

`func NewGetRecordsResponse() *GetRecordsResponse`

NewGetRecordsResponse instantiates a new GetRecordsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordsResponseWithDefaults

`func NewGetRecordsResponseWithDefaults() *GetRecordsResponse`

NewGetRecordsResponseWithDefaults instantiates a new GetRecordsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetRecordsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetRecordsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetRecordsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetRecordsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetHasMore

`func (o *GetRecordsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetRecordsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetRecordsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetRecordsResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNext

`func (o *GetRecordsResponse) GetNext() int32`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetRecordsResponse) GetNextOk() (*int32, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetRecordsResponse) SetNext(v int32)`

SetNext sets Next field to given value.

### HasNext

`func (o *GetRecordsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetRecords

`func (o *GetRecordsResponse) GetRecords() []RecordResource`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *GetRecordsResponse) GetRecordsOk() (*[]RecordResource, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *GetRecordsResponse) SetRecords(v []RecordResource)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *GetRecordsResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


