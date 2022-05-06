# SmartsheetentityGetRecordsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewID** | Pointer to **string** |  | [optional] 
**RecordIDs** | Pointer to **[]string** |  | [optional] 
**FieldTitles** | Pointer to **[]string** |  | [optional] 
**Sort** | Pointer to [**[]SmartsheetentitySort**](SmartsheetentitySort.md) |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewSmartsheetentityGetRecordsRequest

`func NewSmartsheetentityGetRecordsRequest() *SmartsheetentityGetRecordsRequest`

NewSmartsheetentityGetRecordsRequest instantiates a new SmartsheetentityGetRecordsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetentityGetRecordsRequestWithDefaults

`func NewSmartsheetentityGetRecordsRequestWithDefaults() *SmartsheetentityGetRecordsRequest`

NewSmartsheetentityGetRecordsRequestWithDefaults instantiates a new SmartsheetentityGetRecordsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewID

`func (o *SmartsheetentityGetRecordsRequest) GetViewID() string`

GetViewID returns the ViewID field if non-nil, zero value otherwise.

### GetViewIDOk

`func (o *SmartsheetentityGetRecordsRequest) GetViewIDOk() (*string, bool)`

GetViewIDOk returns a tuple with the ViewID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewID

`func (o *SmartsheetentityGetRecordsRequest) SetViewID(v string)`

SetViewID sets ViewID field to given value.

### HasViewID

`func (o *SmartsheetentityGetRecordsRequest) HasViewID() bool`

HasViewID returns a boolean if a field has been set.

### GetRecordIDs

`func (o *SmartsheetentityGetRecordsRequest) GetRecordIDs() []string`

GetRecordIDs returns the RecordIDs field if non-nil, zero value otherwise.

### GetRecordIDsOk

`func (o *SmartsheetentityGetRecordsRequest) GetRecordIDsOk() (*[]string, bool)`

GetRecordIDsOk returns a tuple with the RecordIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIDs

`func (o *SmartsheetentityGetRecordsRequest) SetRecordIDs(v []string)`

SetRecordIDs sets RecordIDs field to given value.

### HasRecordIDs

`func (o *SmartsheetentityGetRecordsRequest) HasRecordIDs() bool`

HasRecordIDs returns a boolean if a field has been set.

### GetFieldTitles

`func (o *SmartsheetentityGetRecordsRequest) GetFieldTitles() []string`

GetFieldTitles returns the FieldTitles field if non-nil, zero value otherwise.

### GetFieldTitlesOk

`func (o *SmartsheetentityGetRecordsRequest) GetFieldTitlesOk() (*[]string, bool)`

GetFieldTitlesOk returns a tuple with the FieldTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitles

`func (o *SmartsheetentityGetRecordsRequest) SetFieldTitles(v []string)`

SetFieldTitles sets FieldTitles field to given value.

### HasFieldTitles

`func (o *SmartsheetentityGetRecordsRequest) HasFieldTitles() bool`

HasFieldTitles returns a boolean if a field has been set.

### GetSort

`func (o *SmartsheetentityGetRecordsRequest) GetSort() []SmartsheetentitySort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SmartsheetentityGetRecordsRequest) GetSortOk() (*[]SmartsheetentitySort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SmartsheetentityGetRecordsRequest) SetSort(v []SmartsheetentitySort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SmartsheetentityGetRecordsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetOffset

`func (o *SmartsheetentityGetRecordsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SmartsheetentityGetRecordsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SmartsheetentityGetRecordsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SmartsheetentityGetRecordsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SmartsheetentityGetRecordsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SmartsheetentityGetRecordsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SmartsheetentityGetRecordsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SmartsheetentityGetRecordsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


