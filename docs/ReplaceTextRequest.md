# ReplaceTextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] 
**Range** | Pointer to [**[]Range**](Range.md) |  | [optional] 

## Methods

### NewReplaceTextRequest

`func NewReplaceTextRequest() *ReplaceTextRequest`

NewReplaceTextRequest instantiates a new ReplaceTextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceTextRequestWithDefaults

`func NewReplaceTextRequestWithDefaults() *ReplaceTextRequest`

NewReplaceTextRequestWithDefaults instantiates a new ReplaceTextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ReplaceTextRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ReplaceTextRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ReplaceTextRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ReplaceTextRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetRange

`func (o *ReplaceTextRequest) GetRange() []Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *ReplaceTextRequest) GetRangeOk() (*[]Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *ReplaceTextRequest) SetRange(v []Range)`

SetRange sets Range field to given value.

### HasRange

`func (o *ReplaceTextRequest) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


