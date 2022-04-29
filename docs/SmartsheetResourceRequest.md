# SmartsheetResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileID** | **string** |  | 
**SheetID** | **string** |  | 
**GetViews** | Pointer to [**GetViewsRequest**](GetViewsRequest.md) |  | [optional] 
**GetRecords** | Pointer to [**GetRecordsRequest**](GetRecordsRequest.md) |  | [optional] 
**GetFields** | Pointer to [**GetFieldsRequest**](GetFieldsRequest.md) |  | [optional] 
**AddView** | Pointer to [**AddViewRequest**](AddViewRequest.md) |  | [optional] 
**DeleteViews** | Pointer to [**DeleteViewsRequest**](DeleteViewsRequest.md) |  | [optional] 
**AddRecords** | Pointer to [**AddRecordsRequest**](AddRecordsRequest.md) |  | [optional] 
**UpdateRecords** | Pointer to [**UpdateRecordsRequest**](UpdateRecordsRequest.md) |  | [optional] 
**DeleteRecords** | Pointer to [**DeleteRecordsRequest**](DeleteRecordsRequest.md) |  | [optional] 
**AddFields** | Pointer to [**AddFieldsRequest**](AddFieldsRequest.md) |  | [optional] 
**UpdateFields** | Pointer to [**UpdateFieldsRequest**](UpdateFieldsRequest.md) |  | [optional] 
**DeleteFields** | Pointer to [**DeleteFieldsRequest**](DeleteFieldsRequest.md) |  | [optional] 

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

`func (o *SmartsheetResourceRequest) GetGetViews() GetViewsRequest`

GetGetViews returns the GetViews field if non-nil, zero value otherwise.

### GetGetViewsOk

`func (o *SmartsheetResourceRequest) GetGetViewsOk() (*GetViewsRequest, bool)`

GetGetViewsOk returns a tuple with the GetViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetViews

`func (o *SmartsheetResourceRequest) SetGetViews(v GetViewsRequest)`

SetGetViews sets GetViews field to given value.

### HasGetViews

`func (o *SmartsheetResourceRequest) HasGetViews() bool`

HasGetViews returns a boolean if a field has been set.

### GetGetRecords

`func (o *SmartsheetResourceRequest) GetGetRecords() GetRecordsRequest`

GetGetRecords returns the GetRecords field if non-nil, zero value otherwise.

### GetGetRecordsOk

`func (o *SmartsheetResourceRequest) GetGetRecordsOk() (*GetRecordsRequest, bool)`

GetGetRecordsOk returns a tuple with the GetRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRecords

`func (o *SmartsheetResourceRequest) SetGetRecords(v GetRecordsRequest)`

SetGetRecords sets GetRecords field to given value.

### HasGetRecords

`func (o *SmartsheetResourceRequest) HasGetRecords() bool`

HasGetRecords returns a boolean if a field has been set.

### GetGetFields

`func (o *SmartsheetResourceRequest) GetGetFields() GetFieldsRequest`

GetGetFields returns the GetFields field if non-nil, zero value otherwise.

### GetGetFieldsOk

`func (o *SmartsheetResourceRequest) GetGetFieldsOk() (*GetFieldsRequest, bool)`

GetGetFieldsOk returns a tuple with the GetFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetFields

`func (o *SmartsheetResourceRequest) SetGetFields(v GetFieldsRequest)`

SetGetFields sets GetFields field to given value.

### HasGetFields

`func (o *SmartsheetResourceRequest) HasGetFields() bool`

HasGetFields returns a boolean if a field has been set.

### GetAddView

`func (o *SmartsheetResourceRequest) GetAddView() AddViewRequest`

GetAddView returns the AddView field if non-nil, zero value otherwise.

### GetAddViewOk

`func (o *SmartsheetResourceRequest) GetAddViewOk() (*AddViewRequest, bool)`

GetAddViewOk returns a tuple with the AddView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddView

`func (o *SmartsheetResourceRequest) SetAddView(v AddViewRequest)`

SetAddView sets AddView field to given value.

### HasAddView

`func (o *SmartsheetResourceRequest) HasAddView() bool`

HasAddView returns a boolean if a field has been set.

### GetDeleteViews

`func (o *SmartsheetResourceRequest) GetDeleteViews() DeleteViewsRequest`

GetDeleteViews returns the DeleteViews field if non-nil, zero value otherwise.

### GetDeleteViewsOk

`func (o *SmartsheetResourceRequest) GetDeleteViewsOk() (*DeleteViewsRequest, bool)`

GetDeleteViewsOk returns a tuple with the DeleteViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteViews

`func (o *SmartsheetResourceRequest) SetDeleteViews(v DeleteViewsRequest)`

SetDeleteViews sets DeleteViews field to given value.

### HasDeleteViews

`func (o *SmartsheetResourceRequest) HasDeleteViews() bool`

HasDeleteViews returns a boolean if a field has been set.

### GetAddRecords

`func (o *SmartsheetResourceRequest) GetAddRecords() AddRecordsRequest`

GetAddRecords returns the AddRecords field if non-nil, zero value otherwise.

### GetAddRecordsOk

`func (o *SmartsheetResourceRequest) GetAddRecordsOk() (*AddRecordsRequest, bool)`

GetAddRecordsOk returns a tuple with the AddRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRecords

`func (o *SmartsheetResourceRequest) SetAddRecords(v AddRecordsRequest)`

SetAddRecords sets AddRecords field to given value.

### HasAddRecords

`func (o *SmartsheetResourceRequest) HasAddRecords() bool`

HasAddRecords returns a boolean if a field has been set.

### GetUpdateRecords

`func (o *SmartsheetResourceRequest) GetUpdateRecords() UpdateRecordsRequest`

GetUpdateRecords returns the UpdateRecords field if non-nil, zero value otherwise.

### GetUpdateRecordsOk

`func (o *SmartsheetResourceRequest) GetUpdateRecordsOk() (*UpdateRecordsRequest, bool)`

GetUpdateRecordsOk returns a tuple with the UpdateRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRecords

`func (o *SmartsheetResourceRequest) SetUpdateRecords(v UpdateRecordsRequest)`

SetUpdateRecords sets UpdateRecords field to given value.

### HasUpdateRecords

`func (o *SmartsheetResourceRequest) HasUpdateRecords() bool`

HasUpdateRecords returns a boolean if a field has been set.

### GetDeleteRecords

`func (o *SmartsheetResourceRequest) GetDeleteRecords() DeleteRecordsRequest`

GetDeleteRecords returns the DeleteRecords field if non-nil, zero value otherwise.

### GetDeleteRecordsOk

`func (o *SmartsheetResourceRequest) GetDeleteRecordsOk() (*DeleteRecordsRequest, bool)`

GetDeleteRecordsOk returns a tuple with the DeleteRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRecords

`func (o *SmartsheetResourceRequest) SetDeleteRecords(v DeleteRecordsRequest)`

SetDeleteRecords sets DeleteRecords field to given value.

### HasDeleteRecords

`func (o *SmartsheetResourceRequest) HasDeleteRecords() bool`

HasDeleteRecords returns a boolean if a field has been set.

### GetAddFields

`func (o *SmartsheetResourceRequest) GetAddFields() AddFieldsRequest`

GetAddFields returns the AddFields field if non-nil, zero value otherwise.

### GetAddFieldsOk

`func (o *SmartsheetResourceRequest) GetAddFieldsOk() (*AddFieldsRequest, bool)`

GetAddFieldsOk returns a tuple with the AddFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFields

`func (o *SmartsheetResourceRequest) SetAddFields(v AddFieldsRequest)`

SetAddFields sets AddFields field to given value.

### HasAddFields

`func (o *SmartsheetResourceRequest) HasAddFields() bool`

HasAddFields returns a boolean if a field has been set.

### GetUpdateFields

`func (o *SmartsheetResourceRequest) GetUpdateFields() UpdateFieldsRequest`

GetUpdateFields returns the UpdateFields field if non-nil, zero value otherwise.

### GetUpdateFieldsOk

`func (o *SmartsheetResourceRequest) GetUpdateFieldsOk() (*UpdateFieldsRequest, bool)`

GetUpdateFieldsOk returns a tuple with the UpdateFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFields

`func (o *SmartsheetResourceRequest) SetUpdateFields(v UpdateFieldsRequest)`

SetUpdateFields sets UpdateFields field to given value.

### HasUpdateFields

`func (o *SmartsheetResourceRequest) HasUpdateFields() bool`

HasUpdateFields returns a boolean if a field has been set.

### GetDeleteFields

`func (o *SmartsheetResourceRequest) GetDeleteFields() DeleteFieldsRequest`

GetDeleteFields returns the DeleteFields field if non-nil, zero value otherwise.

### GetDeleteFieldsOk

`func (o *SmartsheetResourceRequest) GetDeleteFieldsOk() (*DeleteFieldsRequest, bool)`

GetDeleteFieldsOk returns a tuple with the DeleteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFields

`func (o *SmartsheetResourceRequest) SetDeleteFields(v DeleteFieldsRequest)`

SetDeleteFields sets DeleteFields field to given value.

### HasDeleteFields

`func (o *SmartsheetResourceRequest) HasDeleteFields() bool`

HasDeleteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


