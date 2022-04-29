# SelectFieldProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMultiple** | Pointer to **bool** |  | [optional] 
**IsQuickAdd** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to [**[]Option**](Option.md) |  | [optional] 

## Methods

### NewSelectFieldProperty

`func NewSelectFieldProperty() *SelectFieldProperty`

NewSelectFieldProperty instantiates a new SelectFieldProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectFieldPropertyWithDefaults

`func NewSelectFieldPropertyWithDefaults() *SelectFieldProperty`

NewSelectFieldPropertyWithDefaults instantiates a new SelectFieldProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMultiple

`func (o *SelectFieldProperty) GetIsMultiple() bool`

GetIsMultiple returns the IsMultiple field if non-nil, zero value otherwise.

### GetIsMultipleOk

`func (o *SelectFieldProperty) GetIsMultipleOk() (*bool, bool)`

GetIsMultipleOk returns a tuple with the IsMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiple

`func (o *SelectFieldProperty) SetIsMultiple(v bool)`

SetIsMultiple sets IsMultiple field to given value.

### HasIsMultiple

`func (o *SelectFieldProperty) HasIsMultiple() bool`

HasIsMultiple returns a boolean if a field has been set.

### GetIsQuickAdd

`func (o *SelectFieldProperty) GetIsQuickAdd() bool`

GetIsQuickAdd returns the IsQuickAdd field if non-nil, zero value otherwise.

### GetIsQuickAddOk

`func (o *SelectFieldProperty) GetIsQuickAddOk() (*bool, bool)`

GetIsQuickAddOk returns a tuple with the IsQuickAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuickAdd

`func (o *SelectFieldProperty) SetIsQuickAdd(v bool)`

SetIsQuickAdd sets IsQuickAdd field to given value.

### HasIsQuickAdd

`func (o *SelectFieldProperty) HasIsQuickAdd() bool`

HasIsQuickAdd returns a boolean if a field has been set.

### GetOptions

`func (o *SelectFieldProperty) GetOptions() []Option`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SelectFieldProperty) GetOptionsOk() (*[]Option, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SelectFieldProperty) SetOptions(v []Option)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SelectFieldProperty) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


