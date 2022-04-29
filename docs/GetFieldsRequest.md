# GetFieldsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewID** | Pointer to **string** |  | [optional] 
**FieldIDs** | Pointer to **[]string** |  | [optional] 
**FieldTitles** | Pointer to **[]string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetFieldsRequest

`func NewGetFieldsRequest() *GetFieldsRequest`

NewGetFieldsRequest instantiates a new GetFieldsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFieldsRequestWithDefaults

`func NewGetFieldsRequestWithDefaults() *GetFieldsRequest`

NewGetFieldsRequestWithDefaults instantiates a new GetFieldsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewID

`func (o *GetFieldsRequest) GetViewID() string`

GetViewID returns the ViewID field if non-nil, zero value otherwise.

### GetViewIDOk

`func (o *GetFieldsRequest) GetViewIDOk() (*string, bool)`

GetViewIDOk returns a tuple with the ViewID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewID

`func (o *GetFieldsRequest) SetViewID(v string)`

SetViewID sets ViewID field to given value.

### HasViewID

`func (o *GetFieldsRequest) HasViewID() bool`

HasViewID returns a boolean if a field has been set.

### GetFieldIDs

`func (o *GetFieldsRequest) GetFieldIDs() []string`

GetFieldIDs returns the FieldIDs field if non-nil, zero value otherwise.

### GetFieldIDsOk

`func (o *GetFieldsRequest) GetFieldIDsOk() (*[]string, bool)`

GetFieldIDsOk returns a tuple with the FieldIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIDs

`func (o *GetFieldsRequest) SetFieldIDs(v []string)`

SetFieldIDs sets FieldIDs field to given value.

### HasFieldIDs

`func (o *GetFieldsRequest) HasFieldIDs() bool`

HasFieldIDs returns a boolean if a field has been set.

### GetFieldTitles

`func (o *GetFieldsRequest) GetFieldTitles() []string`

GetFieldTitles returns the FieldTitles field if non-nil, zero value otherwise.

### GetFieldTitlesOk

`func (o *GetFieldsRequest) GetFieldTitlesOk() (*[]string, bool)`

GetFieldTitlesOk returns a tuple with the FieldTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitles

`func (o *GetFieldsRequest) SetFieldTitles(v []string)`

SetFieldTitles sets FieldTitles field to given value.

### HasFieldTitles

`func (o *GetFieldsRequest) HasFieldTitles() bool`

HasFieldTitles returns a boolean if a field has been set.

### GetOffset

`func (o *GetFieldsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetFieldsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetFieldsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetFieldsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *GetFieldsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetFieldsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetFieldsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetFieldsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


