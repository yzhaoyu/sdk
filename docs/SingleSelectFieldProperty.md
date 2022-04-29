# SingleSelectFieldProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMultiple** | Pointer to **bool** |  | [optional] 
**IsQuickAdd** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to [**[]Option**](Option.md) |  | [optional] 

## Methods

### NewSingleSelectFieldProperty

`func NewSingleSelectFieldProperty() *SingleSelectFieldProperty`

NewSingleSelectFieldProperty instantiates a new SingleSelectFieldProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleSelectFieldPropertyWithDefaults

`func NewSingleSelectFieldPropertyWithDefaults() *SingleSelectFieldProperty`

NewSingleSelectFieldPropertyWithDefaults instantiates a new SingleSelectFieldProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMultiple

`func (o *SingleSelectFieldProperty) GetIsMultiple() bool`

GetIsMultiple returns the IsMultiple field if non-nil, zero value otherwise.

### GetIsMultipleOk

`func (o *SingleSelectFieldProperty) GetIsMultipleOk() (*bool, bool)`

GetIsMultipleOk returns a tuple with the IsMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiple

`func (o *SingleSelectFieldProperty) SetIsMultiple(v bool)`

SetIsMultiple sets IsMultiple field to given value.

### HasIsMultiple

`func (o *SingleSelectFieldProperty) HasIsMultiple() bool`

HasIsMultiple returns a boolean if a field has been set.

### GetIsQuickAdd

`func (o *SingleSelectFieldProperty) GetIsQuickAdd() bool`

GetIsQuickAdd returns the IsQuickAdd field if non-nil, zero value otherwise.

### GetIsQuickAddOk

`func (o *SingleSelectFieldProperty) GetIsQuickAddOk() (*bool, bool)`

GetIsQuickAddOk returns a tuple with the IsQuickAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuickAdd

`func (o *SingleSelectFieldProperty) SetIsQuickAdd(v bool)`

SetIsQuickAdd sets IsQuickAdd field to given value.

### HasIsQuickAdd

`func (o *SingleSelectFieldProperty) HasIsQuickAdd() bool`

HasIsQuickAdd returns a boolean if a field has been set.

### GetOptions

`func (o *SingleSelectFieldProperty) GetOptions() []Option`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SingleSelectFieldProperty) GetOptionsOk() (*[]Option, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SingleSelectFieldProperty) SetOptions(v []Option)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SingleSelectFieldProperty) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


