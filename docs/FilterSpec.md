# FilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var1** | Pointer to **int32** |  | [optional] 
**Var2** | Pointer to [**[]Condition**](Condition.md) |  | [optional] 

## Methods

### NewFilterSpec

`func NewFilterSpec() *FilterSpec`

NewFilterSpec instantiates a new FilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterSpecWithDefaults

`func NewFilterSpecWithDefaults() *FilterSpec`

NewFilterSpecWithDefaults instantiates a new FilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar1

`func (o *FilterSpec) GetVar1() int32`

GetVar1 returns the Var1 field if non-nil, zero value otherwise.

### GetVar1Ok

`func (o *FilterSpec) GetVar1Ok() (*int32, bool)`

GetVar1Ok returns a tuple with the Var1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1

`func (o *FilterSpec) SetVar1(v int32)`

SetVar1 sets Var1 field to given value.

### HasVar1

`func (o *FilterSpec) HasVar1() bool`

HasVar1 returns a boolean if a field has been set.

### GetVar2

`func (o *FilterSpec) GetVar2() []Condition`

GetVar2 returns the Var2 field if non-nil, zero value otherwise.

### GetVar2Ok

`func (o *FilterSpec) GetVar2Ok() (*[]Condition, bool)`

GetVar2Ok returns a tuple with the Var2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2

`func (o *FilterSpec) SetVar2(v []Condition)`

SetVar2 sets Var2 field to given value.

### HasVar2

`func (o *FilterSpec) HasVar2() bool`

HasVar2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


