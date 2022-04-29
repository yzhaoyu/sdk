# Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var1** | Pointer to [**ListValue**](ListValue.md) |  | [optional] 
**Var2** | Pointer to [**DoubleValue**](DoubleValue.md) |  | [optional] 
**Var3** | Pointer to [**BoolValue**](BoolValue.md) |  | [optional] 
**Var4** | Pointer to [**Int64Value**](Int64Value.md) |  | [optional] 
**Var5** | Pointer to [**ListValue**](ListValue.md) |  | [optional] 
**Var6** | Pointer to **interface{}** | Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values. | [optional] 
**Var7** | Pointer to **interface{}** | Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values. | [optional] 
**Var8** | Pointer to [**ListValue**](ListValue.md) |  | [optional] 
**Var9** | Pointer to [**ListValue**](ListValue.md) |  | [optional] 
**Var10** | Pointer to **interface{}** | Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values. | [optional] 
**Var11** | Pointer to **interface{}** | Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values. | [optional] 
**Var12** | Pointer to **interface{}** | Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values. | [optional] 
**Var13** | Pointer to **interface{}** | Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values. | [optional] 
**Var14** | Pointer to **interface{}** | Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values. | [optional] 
**Var15** | Pointer to [**StringValue**](StringValue.md) |  | [optional] 
**Var16** | Pointer to [**StringValue**](StringValue.md) |  | [optional] 
**Var17** | Pointer to [**ListValue**](ListValue.md) |  | [optional] 

## Methods

### NewValue

`func NewValue() *Value`

NewValue instantiates a new Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueWithDefaults

`func NewValueWithDefaults() *Value`

NewValueWithDefaults instantiates a new Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar1

`func (o *Value) GetVar1() ListValue`

GetVar1 returns the Var1 field if non-nil, zero value otherwise.

### GetVar1Ok

`func (o *Value) GetVar1Ok() (*ListValue, bool)`

GetVar1Ok returns a tuple with the Var1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1

`func (o *Value) SetVar1(v ListValue)`

SetVar1 sets Var1 field to given value.

### HasVar1

`func (o *Value) HasVar1() bool`

HasVar1 returns a boolean if a field has been set.

### GetVar2

`func (o *Value) GetVar2() DoubleValue`

GetVar2 returns the Var2 field if non-nil, zero value otherwise.

### GetVar2Ok

`func (o *Value) GetVar2Ok() (*DoubleValue, bool)`

GetVar2Ok returns a tuple with the Var2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2

`func (o *Value) SetVar2(v DoubleValue)`

SetVar2 sets Var2 field to given value.

### HasVar2

`func (o *Value) HasVar2() bool`

HasVar2 returns a boolean if a field has been set.

### GetVar3

`func (o *Value) GetVar3() BoolValue`

GetVar3 returns the Var3 field if non-nil, zero value otherwise.

### GetVar3Ok

`func (o *Value) GetVar3Ok() (*BoolValue, bool)`

GetVar3Ok returns a tuple with the Var3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3

`func (o *Value) SetVar3(v BoolValue)`

SetVar3 sets Var3 field to given value.

### HasVar3

`func (o *Value) HasVar3() bool`

HasVar3 returns a boolean if a field has been set.

### GetVar4

`func (o *Value) GetVar4() Int64Value`

GetVar4 returns the Var4 field if non-nil, zero value otherwise.

### GetVar4Ok

`func (o *Value) GetVar4Ok() (*Int64Value, bool)`

GetVar4Ok returns a tuple with the Var4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar4

`func (o *Value) SetVar4(v Int64Value)`

SetVar4 sets Var4 field to given value.

### HasVar4

`func (o *Value) HasVar4() bool`

HasVar4 returns a boolean if a field has been set.

### GetVar5

`func (o *Value) GetVar5() ListValue`

GetVar5 returns the Var5 field if non-nil, zero value otherwise.

### GetVar5Ok

`func (o *Value) GetVar5Ok() (*ListValue, bool)`

GetVar5Ok returns a tuple with the Var5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5

`func (o *Value) SetVar5(v ListValue)`

SetVar5 sets Var5 field to given value.

### HasVar5

`func (o *Value) HasVar5() bool`

HasVar5 returns a boolean if a field has been set.

### GetVar6

`func (o *Value) GetVar6() interface{}`

GetVar6 returns the Var6 field if non-nil, zero value otherwise.

### GetVar6Ok

`func (o *Value) GetVar6Ok() (*interface{}, bool)`

GetVar6Ok returns a tuple with the Var6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6

`func (o *Value) SetVar6(v interface{})`

SetVar6 sets Var6 field to given value.

### HasVar6

`func (o *Value) HasVar6() bool`

HasVar6 returns a boolean if a field has been set.

### SetVar6Nil

`func (o *Value) SetVar6Nil(b bool)`

 SetVar6Nil sets the value for Var6 to be an explicit nil

### UnsetVar6
`func (o *Value) UnsetVar6()`

UnsetVar6 ensures that no value is present for Var6, not even an explicit nil
### GetVar7

`func (o *Value) GetVar7() interface{}`

GetVar7 returns the Var7 field if non-nil, zero value otherwise.

### GetVar7Ok

`func (o *Value) GetVar7Ok() (*interface{}, bool)`

GetVar7Ok returns a tuple with the Var7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7

`func (o *Value) SetVar7(v interface{})`

SetVar7 sets Var7 field to given value.

### HasVar7

`func (o *Value) HasVar7() bool`

HasVar7 returns a boolean if a field has been set.

### SetVar7Nil

`func (o *Value) SetVar7Nil(b bool)`

 SetVar7Nil sets the value for Var7 to be an explicit nil

### UnsetVar7
`func (o *Value) UnsetVar7()`

UnsetVar7 ensures that no value is present for Var7, not even an explicit nil
### GetVar8

`func (o *Value) GetVar8() ListValue`

GetVar8 returns the Var8 field if non-nil, zero value otherwise.

### GetVar8Ok

`func (o *Value) GetVar8Ok() (*ListValue, bool)`

GetVar8Ok returns a tuple with the Var8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar8

`func (o *Value) SetVar8(v ListValue)`

SetVar8 sets Var8 field to given value.

### HasVar8

`func (o *Value) HasVar8() bool`

HasVar8 returns a boolean if a field has been set.

### GetVar9

`func (o *Value) GetVar9() ListValue`

GetVar9 returns the Var9 field if non-nil, zero value otherwise.

### GetVar9Ok

`func (o *Value) GetVar9Ok() (*ListValue, bool)`

GetVar9Ok returns a tuple with the Var9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar9

`func (o *Value) SetVar9(v ListValue)`

SetVar9 sets Var9 field to given value.

### HasVar9

`func (o *Value) HasVar9() bool`

HasVar9 returns a boolean if a field has been set.

### GetVar10

`func (o *Value) GetVar10() interface{}`

GetVar10 returns the Var10 field if non-nil, zero value otherwise.

### GetVar10Ok

`func (o *Value) GetVar10Ok() (*interface{}, bool)`

GetVar10Ok returns a tuple with the Var10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar10

`func (o *Value) SetVar10(v interface{})`

SetVar10 sets Var10 field to given value.

### HasVar10

`func (o *Value) HasVar10() bool`

HasVar10 returns a boolean if a field has been set.

### SetVar10Nil

`func (o *Value) SetVar10Nil(b bool)`

 SetVar10Nil sets the value for Var10 to be an explicit nil

### UnsetVar10
`func (o *Value) UnsetVar10()`

UnsetVar10 ensures that no value is present for Var10, not even an explicit nil
### GetVar11

`func (o *Value) GetVar11() interface{}`

GetVar11 returns the Var11 field if non-nil, zero value otherwise.

### GetVar11Ok

`func (o *Value) GetVar11Ok() (*interface{}, bool)`

GetVar11Ok returns a tuple with the Var11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar11

`func (o *Value) SetVar11(v interface{})`

SetVar11 sets Var11 field to given value.

### HasVar11

`func (o *Value) HasVar11() bool`

HasVar11 returns a boolean if a field has been set.

### SetVar11Nil

`func (o *Value) SetVar11Nil(b bool)`

 SetVar11Nil sets the value for Var11 to be an explicit nil

### UnsetVar11
`func (o *Value) UnsetVar11()`

UnsetVar11 ensures that no value is present for Var11, not even an explicit nil
### GetVar12

`func (o *Value) GetVar12() interface{}`

GetVar12 returns the Var12 field if non-nil, zero value otherwise.

### GetVar12Ok

`func (o *Value) GetVar12Ok() (*interface{}, bool)`

GetVar12Ok returns a tuple with the Var12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar12

`func (o *Value) SetVar12(v interface{})`

SetVar12 sets Var12 field to given value.

### HasVar12

`func (o *Value) HasVar12() bool`

HasVar12 returns a boolean if a field has been set.

### SetVar12Nil

`func (o *Value) SetVar12Nil(b bool)`

 SetVar12Nil sets the value for Var12 to be an explicit nil

### UnsetVar12
`func (o *Value) UnsetVar12()`

UnsetVar12 ensures that no value is present for Var12, not even an explicit nil
### GetVar13

`func (o *Value) GetVar13() interface{}`

GetVar13 returns the Var13 field if non-nil, zero value otherwise.

### GetVar13Ok

`func (o *Value) GetVar13Ok() (*interface{}, bool)`

GetVar13Ok returns a tuple with the Var13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar13

`func (o *Value) SetVar13(v interface{})`

SetVar13 sets Var13 field to given value.

### HasVar13

`func (o *Value) HasVar13() bool`

HasVar13 returns a boolean if a field has been set.

### SetVar13Nil

`func (o *Value) SetVar13Nil(b bool)`

 SetVar13Nil sets the value for Var13 to be an explicit nil

### UnsetVar13
`func (o *Value) UnsetVar13()`

UnsetVar13 ensures that no value is present for Var13, not even an explicit nil
### GetVar14

`func (o *Value) GetVar14() interface{}`

GetVar14 returns the Var14 field if non-nil, zero value otherwise.

### GetVar14Ok

`func (o *Value) GetVar14Ok() (*interface{}, bool)`

GetVar14Ok returns a tuple with the Var14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar14

`func (o *Value) SetVar14(v interface{})`

SetVar14 sets Var14 field to given value.

### HasVar14

`func (o *Value) HasVar14() bool`

HasVar14 returns a boolean if a field has been set.

### SetVar14Nil

`func (o *Value) SetVar14Nil(b bool)`

 SetVar14Nil sets the value for Var14 to be an explicit nil

### UnsetVar14
`func (o *Value) UnsetVar14()`

UnsetVar14 ensures that no value is present for Var14, not even an explicit nil
### GetVar15

`func (o *Value) GetVar15() StringValue`

GetVar15 returns the Var15 field if non-nil, zero value otherwise.

### GetVar15Ok

`func (o *Value) GetVar15Ok() (*StringValue, bool)`

GetVar15Ok returns a tuple with the Var15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar15

`func (o *Value) SetVar15(v StringValue)`

SetVar15 sets Var15 field to given value.

### HasVar15

`func (o *Value) HasVar15() bool`

HasVar15 returns a boolean if a field has been set.

### GetVar16

`func (o *Value) GetVar16() StringValue`

GetVar16 returns the Var16 field if non-nil, zero value otherwise.

### GetVar16Ok

`func (o *Value) GetVar16Ok() (*StringValue, bool)`

GetVar16Ok returns a tuple with the Var16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar16

`func (o *Value) SetVar16(v StringValue)`

SetVar16 sets Var16 field to given value.

### HasVar16

`func (o *Value) HasVar16() bool`

HasVar16 returns a boolean if a field has been set.

### GetVar17

`func (o *Value) GetVar17() ListValue`

GetVar17 returns the Var17 field if non-nil, zero value otherwise.

### GetVar17Ok

`func (o *Value) GetVar17Ok() (*ListValue, bool)`

GetVar17Ok returns a tuple with the Var17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar17

`func (o *Value) SetVar17(v ListValue)`

SetVar17 sets Var17 field to given value.

### HasVar17

`func (o *Value) HasVar17() bool`

HasVar17 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


