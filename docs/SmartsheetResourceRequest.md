# SmartsheetResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileID** | **string** |  | 
**SheetID** | **string** |  | 
**GetViews** | Pointer to [**SmartsheetentityGetViewsRequest**](SmartsheetentityGetViewsRequest.md) |  | [optional] 
**GetRecords** | Pointer to [**SmartsheetentityGetRecordsRequest**](SmartsheetentityGetRecordsRequest.md) |  | [optional] 
**GetFields** | Pointer to [**SmartsheetentityGetFieldsRequest**](SmartsheetentityGetFieldsRequest.md) |  | [optional] 
**AddView** | Pointer to [**SmartsheetentityAddViewRequest**](SmartsheetentityAddViewRequest.md) |  | [optional] 
**DeleteViews** | Pointer to [**SmartsheetentityDeleteViewsRequest**](SmartsheetentityDeleteViewsRequest.md) |  | [optional] 
**AddRecords** | Pointer to [**SmartsheetentityAddRecordsRequest**](SmartsheetentityAddRecordsRequest.md) |  | [optional] 
**UpdateRecords** | Pointer to [**SmartsheetentityUpdateRecordsRequest**](SmartsheetentityUpdateRecordsRequest.md) |  | [optional] 
**DeleteRecords** | Pointer to [**SmartsheetentityDeleteRecordsRequest**](SmartsheetentityDeleteRecordsRequest.md) |  | [optional] 
**AddFields** | Pointer to [**SmartsheetentityAddFieldsRequest**](SmartsheetentityAddFieldsRequest.md) |  | [optional] 
**UpdateFields** | Pointer to [**SmartsheetentityUpdateFieldsRequest**](SmartsheetentityUpdateFieldsRequest.md) |  | [optional] 
**DeleteFields** | Pointer to [**SmartsheetentityDeleteFieldsRequest**](SmartsheetentityDeleteFieldsRequest.md) |  | [optional] 

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

`func (o *SmartsheetResourceRequest) GetGetViews() SmartsheetentityGetViewsRequest`

GetGetViews returns the GetViews field if non-nil, zero value otherwise.

### GetGetViewsOk

`func (o *SmartsheetResourceRequest) GetGetViewsOk() (*SmartsheetentityGetViewsRequest, bool)`

GetGetViewsOk returns a tuple with the GetViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetViews

`func (o *SmartsheetResourceRequest) SetGetViews(v SmartsheetentityGetViewsRequest)`

SetGetViews sets GetViews field to given value.

### HasGetViews

`func (o *SmartsheetResourceRequest) HasGetViews() bool`

HasGetViews returns a boolean if a field has been set.

### GetGetRecords

`func (o *SmartsheetResourceRequest) GetGetRecords() SmartsheetentityGetRecordsRequest`

GetGetRecords returns the GetRecords field if non-nil, zero value otherwise.

### GetGetRecordsOk

`func (o *SmartsheetResourceRequest) GetGetRecordsOk() (*SmartsheetentityGetRecordsRequest, bool)`

GetGetRecordsOk returns a tuple with the GetRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRecords

`func (o *SmartsheetResourceRequest) SetGetRecords(v SmartsheetentityGetRecordsRequest)`

SetGetRecords sets GetRecords field to given value.

### HasGetRecords

`func (o *SmartsheetResourceRequest) HasGetRecords() bool`

HasGetRecords returns a boolean if a field has been set.

### GetGetFields

`func (o *SmartsheetResourceRequest) GetGetFields() SmartsheetentityGetFieldsRequest`

GetGetFields returns the GetFields field if non-nil, zero value otherwise.

### GetGetFieldsOk

`func (o *SmartsheetResourceRequest) GetGetFieldsOk() (*SmartsheetentityGetFieldsRequest, bool)`

GetGetFieldsOk returns a tuple with the GetFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetFields

`func (o *SmartsheetResourceRequest) SetGetFields(v SmartsheetentityGetFieldsRequest)`

SetGetFields sets GetFields field to given value.

### HasGetFields

`func (o *SmartsheetResourceRequest) HasGetFields() bool`

HasGetFields returns a boolean if a field has been set.

### GetAddView

`func (o *SmartsheetResourceRequest) GetAddView() SmartsheetentityAddViewRequest`

GetAddView returns the AddView field if non-nil, zero value otherwise.

### GetAddViewOk

`func (o *SmartsheetResourceRequest) GetAddViewOk() (*SmartsheetentityAddViewRequest, bool)`

GetAddViewOk returns a tuple with the AddView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddView

`func (o *SmartsheetResourceRequest) SetAddView(v SmartsheetentityAddViewRequest)`

SetAddView sets AddView field to given value.

### HasAddView

`func (o *SmartsheetResourceRequest) HasAddView() bool`

HasAddView returns a boolean if a field has been set.

### GetDeleteViews

`func (o *SmartsheetResourceRequest) GetDeleteViews() SmartsheetentityDeleteViewsRequest`

GetDeleteViews returns the DeleteViews field if non-nil, zero value otherwise.

### GetDeleteViewsOk

`func (o *SmartsheetResourceRequest) GetDeleteViewsOk() (*SmartsheetentityDeleteViewsRequest, bool)`

GetDeleteViewsOk returns a tuple with the DeleteViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteViews

`func (o *SmartsheetResourceRequest) SetDeleteViews(v SmartsheetentityDeleteViewsRequest)`

SetDeleteViews sets DeleteViews field to given value.

### HasDeleteViews

`func (o *SmartsheetResourceRequest) HasDeleteViews() bool`

HasDeleteViews returns a boolean if a field has been set.

### GetAddRecords

`func (o *SmartsheetResourceRequest) GetAddRecords() SmartsheetentityAddRecordsRequest`

GetAddRecords returns the AddRecords field if non-nil, zero value otherwise.

### GetAddRecordsOk

`func (o *SmartsheetResourceRequest) GetAddRecordsOk() (*SmartsheetentityAddRecordsRequest, bool)`

GetAddRecordsOk returns a tuple with the AddRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRecords

`func (o *SmartsheetResourceRequest) SetAddRecords(v SmartsheetentityAddRecordsRequest)`

SetAddRecords sets AddRecords field to given value.

### HasAddRecords

`func (o *SmartsheetResourceRequest) HasAddRecords() bool`

HasAddRecords returns a boolean if a field has been set.

### GetUpdateRecords

`func (o *SmartsheetResourceRequest) GetUpdateRecords() SmartsheetentityUpdateRecordsRequest`

GetUpdateRecords returns the UpdateRecords field if non-nil, zero value otherwise.

### GetUpdateRecordsOk

`func (o *SmartsheetResourceRequest) GetUpdateRecordsOk() (*SmartsheetentityUpdateRecordsRequest, bool)`

GetUpdateRecordsOk returns a tuple with the UpdateRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRecords

`func (o *SmartsheetResourceRequest) SetUpdateRecords(v SmartsheetentityUpdateRecordsRequest)`

SetUpdateRecords sets UpdateRecords field to given value.

### HasUpdateRecords

`func (o *SmartsheetResourceRequest) HasUpdateRecords() bool`

HasUpdateRecords returns a boolean if a field has been set.

### GetDeleteRecords

`func (o *SmartsheetResourceRequest) GetDeleteRecords() SmartsheetentityDeleteRecordsRequest`

GetDeleteRecords returns the DeleteRecords field if non-nil, zero value otherwise.

### GetDeleteRecordsOk

`func (o *SmartsheetResourceRequest) GetDeleteRecordsOk() (*SmartsheetentityDeleteRecordsRequest, bool)`

GetDeleteRecordsOk returns a tuple with the DeleteRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRecords

`func (o *SmartsheetResourceRequest) SetDeleteRecords(v SmartsheetentityDeleteRecordsRequest)`

SetDeleteRecords sets DeleteRecords field to given value.

### HasDeleteRecords

`func (o *SmartsheetResourceRequest) HasDeleteRecords() bool`

HasDeleteRecords returns a boolean if a field has been set.

### GetAddFields

`func (o *SmartsheetResourceRequest) GetAddFields() SmartsheetentityAddFieldsRequest`

GetAddFields returns the AddFields field if non-nil, zero value otherwise.

### GetAddFieldsOk

`func (o *SmartsheetResourceRequest) GetAddFieldsOk() (*SmartsheetentityAddFieldsRequest, bool)`

GetAddFieldsOk returns a tuple with the AddFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFields

`func (o *SmartsheetResourceRequest) SetAddFields(v SmartsheetentityAddFieldsRequest)`

SetAddFields sets AddFields field to given value.

### HasAddFields

`func (o *SmartsheetResourceRequest) HasAddFields() bool`

HasAddFields returns a boolean if a field has been set.

### GetUpdateFields

`func (o *SmartsheetResourceRequest) GetUpdateFields() SmartsheetentityUpdateFieldsRequest`

GetUpdateFields returns the UpdateFields field if non-nil, zero value otherwise.

### GetUpdateFieldsOk

`func (o *SmartsheetResourceRequest) GetUpdateFieldsOk() (*SmartsheetentityUpdateFieldsRequest, bool)`

GetUpdateFieldsOk returns a tuple with the UpdateFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFields

`func (o *SmartsheetResourceRequest) SetUpdateFields(v SmartsheetentityUpdateFieldsRequest)`

SetUpdateFields sets UpdateFields field to given value.

### HasUpdateFields

`func (o *SmartsheetResourceRequest) HasUpdateFields() bool`

HasUpdateFields returns a boolean if a field has been set.

### GetDeleteFields

`func (o *SmartsheetResourceRequest) GetDeleteFields() SmartsheetentityDeleteFieldsRequest`

GetDeleteFields returns the DeleteFields field if non-nil, zero value otherwise.

### GetDeleteFieldsOk

`func (o *SmartsheetResourceRequest) GetDeleteFieldsOk() (*SmartsheetentityDeleteFieldsRequest, bool)`

GetDeleteFieldsOk returns a tuple with the DeleteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFields

`func (o *SmartsheetResourceRequest) SetDeleteFields(v SmartsheetentityDeleteFieldsRequest)`

SetDeleteFields sets DeleteFields field to given value.

### HasDeleteFields

`func (o *SmartsheetResourceRequest) HasDeleteFields() bool`

HasDeleteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


