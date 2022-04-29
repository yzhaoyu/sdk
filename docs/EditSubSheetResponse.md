# EditSubSheetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddSheet** | Pointer to [**SmartsheetAddSheetRsp**](SmartsheetAddSheetRsp.md) |  | [optional] 
**DeleteSheet** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEditSubSheetResponse

`func NewEditSubSheetResponse() *EditSubSheetResponse`

NewEditSubSheetResponse instantiates a new EditSubSheetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditSubSheetResponseWithDefaults

`func NewEditSubSheetResponseWithDefaults() *EditSubSheetResponse`

NewEditSubSheetResponseWithDefaults instantiates a new EditSubSheetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddSheet

`func (o *EditSubSheetResponse) GetAddSheet() SmartsheetAddSheetRsp`

GetAddSheet returns the AddSheet field if non-nil, zero value otherwise.

### GetAddSheetOk

`func (o *EditSubSheetResponse) GetAddSheetOk() (*SmartsheetAddSheetRsp, bool)`

GetAddSheetOk returns a tuple with the AddSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSheet

`func (o *EditSubSheetResponse) SetAddSheet(v SmartsheetAddSheetRsp)`

SetAddSheet sets AddSheet field to given value.

### HasAddSheet

`func (o *EditSubSheetResponse) HasAddSheet() bool`

HasAddSheet returns a boolean if a field has been set.

### GetDeleteSheet

`func (o *EditSubSheetResponse) GetDeleteSheet() map[string]interface{}`

GetDeleteSheet returns the DeleteSheet field if non-nil, zero value otherwise.

### GetDeleteSheetOk

`func (o *EditSubSheetResponse) GetDeleteSheetOk() (*map[string]interface{}, bool)`

GetDeleteSheetOk returns a tuple with the DeleteSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSheet

`func (o *EditSubSheetResponse) SetDeleteSheet(v map[string]interface{})`

SetDeleteSheet sets DeleteSheet field to given value.

### HasDeleteSheet

`func (o *EditSubSheetResponse) HasDeleteSheet() bool`

HasDeleteSheet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


