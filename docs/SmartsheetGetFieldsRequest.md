# SmartsheetGetFieldsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewID** | Pointer to **string** |  | [optional] 
**FieldIDs** | Pointer to **[]string** |  | [optional] 
**FieldTitles** | Pointer to **[]string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewSmartsheetGetFieldsRequest

`func NewSmartsheetGetFieldsRequest() *SmartsheetGetFieldsRequest`

NewSmartsheetGetFieldsRequest instantiates a new SmartsheetGetFieldsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetGetFieldsRequestWithDefaults

`func NewSmartsheetGetFieldsRequestWithDefaults() *SmartsheetGetFieldsRequest`

NewSmartsheetGetFieldsRequestWithDefaults instantiates a new SmartsheetGetFieldsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewID

`func (o *SmartsheetGetFieldsRequest) GetViewID() string`

GetViewID returns the ViewID field if non-nil, zero value otherwise.

### GetViewIDOk

`func (o *SmartsheetGetFieldsRequest) GetViewIDOk() (*string, bool)`

GetViewIDOk returns a tuple with the ViewID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewID

`func (o *SmartsheetGetFieldsRequest) SetViewID(v string)`

SetViewID sets ViewID field to given value.

### HasViewID

`func (o *SmartsheetGetFieldsRequest) HasViewID() bool`

HasViewID returns a boolean if a field has been set.

### GetFieldIDs

`func (o *SmartsheetGetFieldsRequest) GetFieldIDs() []string`

GetFieldIDs returns the FieldIDs field if non-nil, zero value otherwise.

### GetFieldIDsOk

`func (o *SmartsheetGetFieldsRequest) GetFieldIDsOk() (*[]string, bool)`

GetFieldIDsOk returns a tuple with the FieldIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIDs

`func (o *SmartsheetGetFieldsRequest) SetFieldIDs(v []string)`

SetFieldIDs sets FieldIDs field to given value.

### HasFieldIDs

`func (o *SmartsheetGetFieldsRequest) HasFieldIDs() bool`

HasFieldIDs returns a boolean if a field has been set.

### GetFieldTitles

`func (o *SmartsheetGetFieldsRequest) GetFieldTitles() []string`

GetFieldTitles returns the FieldTitles field if non-nil, zero value otherwise.

### GetFieldTitlesOk

`func (o *SmartsheetGetFieldsRequest) GetFieldTitlesOk() (*[]string, bool)`

GetFieldTitlesOk returns a tuple with the FieldTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitles

`func (o *SmartsheetGetFieldsRequest) SetFieldTitles(v []string)`

SetFieldTitles sets FieldTitles field to given value.

### HasFieldTitles

`func (o *SmartsheetGetFieldsRequest) HasFieldTitles() bool`

HasFieldTitles returns a boolean if a field has been set.

### GetOffset

`func (o *SmartsheetGetFieldsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SmartsheetGetFieldsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SmartsheetGetFieldsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SmartsheetGetFieldsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SmartsheetGetFieldsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SmartsheetGetFieldsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SmartsheetGetFieldsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SmartsheetGetFieldsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


