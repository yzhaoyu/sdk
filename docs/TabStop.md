# TabStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to [**GoogleProtobufDoubleValue**](GoogleProtobufDoubleValue.md) |  | [optional] 
**Alignment** | Pointer to **string** |  | [optional] 

## Methods

### NewTabStop

`func NewTabStop() *TabStop`

NewTabStop instantiates a new TabStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTabStopWithDefaults

`func NewTabStopWithDefaults() *TabStop`

NewTabStopWithDefaults instantiates a new TabStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *TabStop) GetOffset() GoogleProtobufDoubleValue`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TabStop) GetOffsetOk() (*GoogleProtobufDoubleValue, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TabStop) SetOffset(v GoogleProtobufDoubleValue)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TabStop) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetAlignment

`func (o *TabStop) GetAlignment() string`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *TabStop) GetAlignmentOk() (*string, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *TabStop) SetAlignment(v string)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *TabStop) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


