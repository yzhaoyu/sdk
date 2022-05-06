# SmartsheetResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileID** | **string** |  | 
**SheetID** | **string** |  | 
**GetViews** | Pointer to [**SmartsheetGetViewsRequest**](SmartsheetGetViewsRequest.md) |  | [optional] 
**GetRecords** | Pointer to [**SmartsheetGetRecordsRequest**](SmartsheetGetRecordsRequest.md) |  | [optional] 
**GetFields** | Pointer to [**SmartsheetGetFieldsRequest**](SmartsheetGetFieldsRequest.md) |  | [optional] 
**AddView** | Pointer to [**SmartsheetAddViewRequest**](SmartsheetAddViewRequest.md) |  | [optional] 
**DeleteViews** | Pointer to [**SmartsheetDeleteViewsRequest**](SmartsheetDeleteViewsRequest.md) |  | [optional] 
**AddRecords** | Pointer to [**SmartsheetAddRecordsRequest**](SmartsheetAddRecordsRequest.md) |  | [optional] 
**UpdateRecords** | Pointer to [**SmartsheetUpdateRecordsRequest**](SmartsheetUpdateRecordsRequest.md) |  | [optional] 
**DeleteRecords** | Pointer to [**SmartsheetDeleteRecordsRequest**](SmartsheetDeleteRecordsRequest.md) |  | [optional] 
**AddFields** | Pointer to [**SmartsheetAddFieldsRequest**](SmartsheetAddFieldsRequest.md) |  | [optional] 
**UpdateFields** | Pointer to [**SmartsheetUpdateFieldsRequest**](SmartsheetUpdateFieldsRequest.md) |  | [optional] 
**DeleteFields** | Pointer to [**SmartsheetDeleteFieldsRequest**](SmartsheetDeleteFieldsRequest.md) |  | [optional] 

## Methods

### NewSmartsheetResourceRequest

`func NewSmartsheetResourceRequest(fileID string, sheetID string, ) *SmartsheetResourceRequest`

NewSmartsheetResourceRequest instantiates a new SmartsheetResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetResourceRequestWithDefaults

`func NewSmartsheetResourceRequestWithDefaults() *SmartsheetResourceRequest`

NewSmartsheetResourceRequestWithDefaults instantiates a new SmartsheetResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileID

`func (o *SmartsheetResourceRequest) GetFileID() string`

GetFileID returns the FileID field if non-nil, zero value otherwise.

### GetFileIDOk

`func (o *SmartsheetResourceRequest) GetFileIDOk() (*string, bool)`

GetFileIDOk returns a tuple with the FileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileID

`func (o *SmartsheetResourceRequest) SetFileID(v string)`

SetFileID sets FileID field to given value.


### GetSheetID

`func (o *SmartsheetResourceRequest) GetSheetID() string`

GetSheetID returns the SheetID field if non-nil, zero value otherwise.

### GetSheetIDOk

`func (o *SmartsheetResourceRequest) GetSheetIDOk() (*string, bool)`

GetSheetIDOk returns a tuple with the SheetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetID

`func (o *SmartsheetResourceRequest) SetSheetID(v string)`

SetSheetID sets SheetID field to given value.


### GetGetViews

`func (o *SmartsheetResourceRequest) GetGetViews() SmartsheetGetViewsRequest`

GetGetViews returns the GetViews field if non-nil, zero value otherwise.

### GetGetViewsOk

`func (o *SmartsheetResourceRequest) GetGetViewsOk() (*SmartsheetGetViewsRequest, bool)`

GetGetViewsOk returns a tuple with the GetViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetViews

`func (o *SmartsheetResourceRequest) SetGetViews(v SmartsheetGetViewsRequest)`

SetGetViews sets GetViews field to given value.

### HasGetViews

`func (o *SmartsheetResourceRequest) HasGetViews() bool`

HasGetViews returns a boolean if a field has been set.

### GetGetRecords

`func (o *SmartsheetResourceRequest) GetGetRecords() SmartsheetGetRecordsRequest`

GetGetRecords returns the GetRecords field if non-nil, zero value otherwise.

### GetGetRecordsOk

`func (o *SmartsheetResourceRequest) GetGetRecordsOk() (*SmartsheetGetRecordsRequest, bool)`

GetGetRecordsOk returns a tuple with the GetRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRecords

`func (o *SmartsheetResourceRequest) SetGetRecords(v SmartsheetGetRecordsRequest)`

SetGetRecords sets GetRecords field to given value.

### HasGetRecords

`func (o *SmartsheetResourceRequest) HasGetRecords() bool`

HasGetRecords returns a boolean if a field has been set.

### GetGetFields

`func (o *SmartsheetResourceRequest) GetGetFields() SmartsheetGetFieldsRequest`

GetGetFields returns the GetFields field if non-nil, zero value otherwise.

### GetGetFieldsOk

`func (o *SmartsheetResourceRequest) GetGetFieldsOk() (*SmartsheetGetFieldsRequest, bool)`

GetGetFieldsOk returns a tuple with the GetFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetFields

`func (o *SmartsheetResourceRequest) SetGetFields(v SmartsheetGetFieldsRequest)`

SetGetFields sets GetFields field to given value.

### HasGetFields

`func (o *SmartsheetResourceRequest) HasGetFields() bool`

HasGetFields returns a boolean if a field has been set.

### GetAddView

`func (o *SmartsheetResourceRequest) GetAddView() SmartsheetAddViewRequest`

GetAddView returns the AddView field if non-nil, zero value otherwise.

### GetAddViewOk

`func (o *SmartsheetResourceRequest) GetAddViewOk() (*SmartsheetAddViewRequest, bool)`

GetAddViewOk returns a tuple with the AddView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddView

`func (o *SmartsheetResourceRequest) SetAddView(v SmartsheetAddViewRequest)`

SetAddView sets AddView field to given value.

### HasAddView

`func (o *SmartsheetResourceRequest) HasAddView() bool`

HasAddView returns a boolean if a field has been set.

### GetDeleteViews

`func (o *SmartsheetResourceRequest) GetDeleteViews() SmartsheetDeleteViewsRequest`

GetDeleteViews returns the DeleteViews field if non-nil, zero value otherwise.

### GetDeleteViewsOk

`func (o *SmartsheetResourceRequest) GetDeleteViewsOk() (*SmartsheetDeleteViewsRequest, bool)`

GetDeleteViewsOk returns a tuple with the DeleteViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteViews

`func (o *SmartsheetResourceRequest) SetDeleteViews(v SmartsheetDeleteViewsRequest)`

SetDeleteViews sets DeleteViews field to given value.

### HasDeleteViews

`func (o *SmartsheetResourceRequest) HasDeleteViews() bool`

HasDeleteViews returns a boolean if a field has been set.

### GetAddRecords

`func (o *SmartsheetResourceRequest) GetAddRecords() SmartsheetAddRecordsRequest`

GetAddRecords returns the AddRecords field if non-nil, zero value otherwise.

### GetAddRecordsOk

`func (o *SmartsheetResourceRequest) GetAddRecordsOk() (*SmartsheetAddRecordsRequest, bool)`

GetAddRecordsOk returns a tuple with the AddRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRecords

`func (o *SmartsheetResourceRequest) SetAddRecords(v SmartsheetAddRecordsRequest)`

SetAddRecords sets AddRecords field to given value.

### HasAddRecords

`func (o *SmartsheetResourceRequest) HasAddRecords() bool`

HasAddRecords returns a boolean if a field has been set.

### GetUpdateRecords

`func (o *SmartsheetResourceRequest) GetUpdateRecords() SmartsheetUpdateRecordsRequest`

GetUpdateRecords returns the UpdateRecords field if non-nil, zero value otherwise.

### GetUpdateRecordsOk

`func (o *SmartsheetResourceRequest) GetUpdateRecordsOk() (*SmartsheetUpdateRecordsRequest, bool)`

GetUpdateRecordsOk returns a tuple with the UpdateRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRecords

`func (o *SmartsheetResourceRequest) SetUpdateRecords(v SmartsheetUpdateRecordsRequest)`

SetUpdateRecords sets UpdateRecords field to given value.

### HasUpdateRecords

`func (o *SmartsheetResourceRequest) HasUpdateRecords() bool`

HasUpdateRecords returns a boolean if a field has been set.

### GetDeleteRecords

`func (o *SmartsheetResourceRequest) GetDeleteRecords() SmartsheetDeleteRecordsRequest`

GetDeleteRecords returns the DeleteRecords field if non-nil, zero value otherwise.

### GetDeleteRecordsOk

`func (o *SmartsheetResourceRequest) GetDeleteRecordsOk() (*SmartsheetDeleteRecordsRequest, bool)`

GetDeleteRecordsOk returns a tuple with the DeleteRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRecords

`func (o *SmartsheetResourceRequest) SetDeleteRecords(v SmartsheetDeleteRecordsRequest)`

SetDeleteRecords sets DeleteRecords field to given value.

### HasDeleteRecords

`func (o *SmartsheetResourceRequest) HasDeleteRecords() bool`

HasDeleteRecords returns a boolean if a field has been set.

### GetAddFields

`func (o *SmartsheetResourceRequest) GetAddFields() SmartsheetAddFieldsRequest`

GetAddFields returns the AddFields field if non-nil, zero value otherwise.

### GetAddFieldsOk

`func (o *SmartsheetResourceRequest) GetAddFieldsOk() (*SmartsheetAddFieldsRequest, bool)`

GetAddFieldsOk returns a tuple with the AddFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFields

`func (o *SmartsheetResourceRequest) SetAddFields(v SmartsheetAddFieldsRequest)`

SetAddFields sets AddFields field to given value.

### HasAddFields

`func (o *SmartsheetResourceRequest) HasAddFields() bool`

HasAddFields returns a boolean if a field has been set.

### GetUpdateFields

`func (o *SmartsheetResourceRequest) GetUpdateFields() SmartsheetUpdateFieldsRequest`

GetUpdateFields returns the UpdateFields field if non-nil, zero value otherwise.

### GetUpdateFieldsOk

`func (o *SmartsheetResourceRequest) GetUpdateFieldsOk() (*SmartsheetUpdateFieldsRequest, bool)`

GetUpdateFieldsOk returns a tuple with the UpdateFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFields

`func (o *SmartsheetResourceRequest) SetUpdateFields(v SmartsheetUpdateFieldsRequest)`

SetUpdateFields sets UpdateFields field to given value.

### HasUpdateFields

`func (o *SmartsheetResourceRequest) HasUpdateFields() bool`

HasUpdateFields returns a boolean if a field has been set.

### GetDeleteFields

`func (o *SmartsheetResourceRequest) GetDeleteFields() SmartsheetDeleteFieldsRequest`

GetDeleteFields returns the DeleteFields field if non-nil, zero value otherwise.

### GetDeleteFieldsOk

`func (o *SmartsheetResourceRequest) GetDeleteFieldsOk() (*SmartsheetDeleteFieldsRequest, bool)`

GetDeleteFieldsOk returns a tuple with the DeleteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFields

`func (o *SmartsheetResourceRequest) SetDeleteFields(v SmartsheetDeleteFieldsRequest)`

SetDeleteFields sets DeleteFields field to given value.

### HasDeleteFields

`func (o *SmartsheetResourceRequest) HasDeleteFields() bool`

HasDeleteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


