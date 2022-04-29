# EditSubSheetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileID** | Pointer to **string** |  | [optional] 
**AddSheet** | Pointer to [**SmartsheetAddSheetReq**](SmartsheetAddSheetReq.md) |  | [optional] 
**DeleteSheet** | Pointer to [**SmartsheetDeleteSheetReq**](SmartsheetDeleteSheetReq.md) |  | [optional] 

## Methods

### NewEditSubSheetRequest

`func NewEditSubSheetRequest() *EditSubSheetRequest`

NewEditSubSheetRequest instantiates a new EditSubSheetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditSubSheetRequestWithDefaults

`func NewEditSubSheetRequestWithDefaults() *EditSubSheetRequest`

NewEditSubSheetRequestWithDefaults instantiates a new EditSubSheetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileID

`func (o *EditSubSheetRequest) GetFileID() string`

GetFileID returns the FileID field if non-nil, zero value otherwise.

### GetFileIDOk

`func (o *EditSubSheetRequest) GetFileIDOk() (*string, bool)`

GetFileIDOk returns a tuple with the FileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileID

`func (o *EditSubSheetRequest) SetFileID(v string)`

SetFileID sets FileID field to given value.

### HasFileID

`func (o *EditSubSheetRequest) HasFileID() bool`

HasFileID returns a boolean if a field has been set.

### GetAddSheet

`func (o *EditSubSheetRequest) GetAddSheet() SmartsheetAddSheetReq`

GetAddSheet returns the AddSheet field if non-nil, zero value otherwise.

### GetAddSheetOk

`func (o *EditSubSheetRequest) GetAddSheetOk() (*SmartsheetAddSheetReq, bool)`

GetAddSheetOk returns a tuple with the AddSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSheet

`func (o *EditSubSheetRequest) SetAddSheet(v SmartsheetAddSheetReq)`

SetAddSheet sets AddSheet field to given value.

### HasAddSheet

`func (o *EditSubSheetRequest) HasAddSheet() bool`

HasAddSheet returns a boolean if a field has been set.

### GetDeleteSheet

`func (o *EditSubSheetRequest) GetDeleteSheet() SmartsheetDeleteSheetReq`

GetDeleteSheet returns the DeleteSheet field if non-nil, zero value otherwise.

### GetDeleteSheetOk

`func (o *EditSubSheetRequest) GetDeleteSheetOk() (*SmartsheetDeleteSheetReq, bool)`

GetDeleteSheetOk returns a tuple with the DeleteSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSheet

`func (o *EditSubSheetRequest) SetDeleteSheet(v SmartsheetDeleteSheetReq)`

SetDeleteSheet sets DeleteSheet field to given value.

### HasDeleteSheet

`func (o *EditSubSheetRequest) HasDeleteSheet() bool`

HasDeleteSheet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


