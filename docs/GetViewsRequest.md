# GetViewsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewIDs** | Pointer to **[]string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetViewsRequest

`func NewGetViewsRequest() *GetViewsRequest`

NewGetViewsRequest instantiates a new GetViewsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetViewsRequestWithDefaults

`func NewGetViewsRequestWithDefaults() *GetViewsRequest`

NewGetViewsRequestWithDefaults instantiates a new GetViewsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewIDs

`func (o *GetViewsRequest) GetViewIDs() []string`

GetViewIDs returns the ViewIDs field if non-nil, zero value otherwise.

### GetViewIDsOk

`func (o *GetViewsRequest) GetViewIDsOk() (*[]string, bool)`

GetViewIDsOk returns a tuple with the ViewIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewIDs

`func (o *GetViewsRequest) SetViewIDs(v []string)`

SetViewIDs sets ViewIDs field to given value.

### HasViewIDs

`func (o *GetViewsRequest) HasViewIDs() bool`

HasViewIDs returns a boolean if a field has been set.

### GetOffset

`func (o *GetViewsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetViewsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetViewsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetViewsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *GetViewsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetViewsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetViewsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetViewsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


