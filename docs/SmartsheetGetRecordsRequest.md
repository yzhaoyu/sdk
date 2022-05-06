# SmartsheetGetRecordsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewID** | Pointer to **string** |  | [optional] 
**RecordIDs** | Pointer to **[]string** |  | [optional] 
**FieldTitles** | Pointer to **[]string** |  | [optional] 
**Sort** | Pointer to [**[]SmartsheetSort**](SmartsheetSort.md) |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewSmartsheetGetRecordsRequest

`func NewSmartsheetGetRecordsRequest() *SmartsheetGetRecordsRequest`

NewSmartsheetGetRecordsRequest instantiates a new SmartsheetGetRecordsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetGetRecordsRequestWithDefaults

`func NewSmartsheetGetRecordsRequestWithDefaults() *SmartsheetGetRecordsRequest`

NewSmartsheetGetRecordsRequestWithDefaults instantiates a new SmartsheetGetRecordsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewID

`func (o *SmartsheetGetRecordsRequest) GetViewID() string`

GetViewID returns the ViewID field if non-nil, zero value otherwise.

### GetViewIDOk

`func (o *SmartsheetGetRecordsRequest) GetViewIDOk() (*string, bool)`

GetViewIDOk returns a tuple with the ViewID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewID

`func (o *SmartsheetGetRecordsRequest) SetViewID(v string)`

SetViewID sets ViewID field to given value.

### HasViewID

`func (o *SmartsheetGetRecordsRequest) HasViewID() bool`

HasViewID returns a boolean if a field has been set.

### GetRecordIDs

`func (o *SmartsheetGetRecordsRequest) GetRecordIDs() []string`

GetRecordIDs returns the RecordIDs field if non-nil, zero value otherwise.

### GetRecordIDsOk

`func (o *SmartsheetGetRecordsRequest) GetRecordIDsOk() (*[]string, bool)`

GetRecordIDsOk returns a tuple with the RecordIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIDs

`func (o *SmartsheetGetRecordsRequest) SetRecordIDs(v []string)`

SetRecordIDs sets RecordIDs field to given value.

### HasRecordIDs

`func (o *SmartsheetGetRecordsRequest) HasRecordIDs() bool`

HasRecordIDs returns a boolean if a field has been set.

### GetFieldTitles

`func (o *SmartsheetGetRecordsRequest) GetFieldTitles() []string`

GetFieldTitles returns the FieldTitles field if non-nil, zero value otherwise.

### GetFieldTitlesOk

`func (o *SmartsheetGetRecordsRequest) GetFieldTitlesOk() (*[]string, bool)`

GetFieldTitlesOk returns a tuple with the FieldTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitles

`func (o *SmartsheetGetRecordsRequest) SetFieldTitles(v []string)`

SetFieldTitles sets FieldTitles field to given value.

### HasFieldTitles

`func (o *SmartsheetGetRecordsRequest) HasFieldTitles() bool`

HasFieldTitles returns a boolean if a field has been set.

### GetSort

`func (o *SmartsheetGetRecordsRequest) GetSort() []SmartsheetSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SmartsheetGetRecordsRequest) GetSortOk() (*[]SmartsheetSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SmartsheetGetRecordsRequest) SetSort(v []SmartsheetSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SmartsheetGetRecordsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetOffset

`func (o *SmartsheetGetRecordsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SmartsheetGetRecordsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SmartsheetGetRecordsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SmartsheetGetRecordsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SmartsheetGetRecordsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SmartsheetGetRecordsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SmartsheetGetRecordsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SmartsheetGetRecordsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


