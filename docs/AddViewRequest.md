# AddViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewTitle** | Pointer to **string** |  | [optional] 
**ViewType** | Pointer to **int32** |  | [optional] 

## Methods

### NewAddViewRequest

`func NewAddViewRequest() *AddViewRequest`

NewAddViewRequest instantiates a new AddViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddViewRequestWithDefaults

`func NewAddViewRequestWithDefaults() *AddViewRequest`

NewAddViewRequestWithDefaults instantiates a new AddViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewTitle

`func (o *AddViewRequest) GetViewTitle() string`

GetViewTitle returns the ViewTitle field if non-nil, zero value otherwise.

### GetViewTitleOk

`func (o *AddViewRequest) GetViewTitleOk() (*string, bool)`

GetViewTitleOk returns a tuple with the ViewTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewTitle

`func (o *AddViewRequest) SetViewTitle(v string)`

SetViewTitle sets ViewTitle field to given value.

### HasViewTitle

`func (o *AddViewRequest) HasViewTitle() bool`

HasViewTitle returns a boolean if a field has been set.

### GetViewType

`func (o *AddViewRequest) GetViewType() int32`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *AddViewRequest) GetViewTypeOk() (*int32, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *AddViewRequest) SetViewType(v int32)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *AddViewRequest) HasViewType() bool`

HasViewType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


