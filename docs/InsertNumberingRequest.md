# InsertNumberingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to [**Range**](Range.md) |  | [optional] 
**AbstractNumID** | Pointer to **string** |  | [optional] 
**ParagraphIndex** | Pointer to **int32** |  | [optional] 
**PStyle** | Pointer to [**ParagraphStyle**](ParagraphStyle.md) |  | [optional] 

## Methods

### NewInsertNumberingRequest

`func NewInsertNumberingRequest() *InsertNumberingRequest`

NewInsertNumberingRequest instantiates a new InsertNumberingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsertNumberingRequestWithDefaults

`func NewInsertNumberingRequestWithDefaults() *InsertNumberingRequest`

NewInsertNumberingRequestWithDefaults instantiates a new InsertNumberingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *InsertNumberingRequest) GetRange() Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *InsertNumberingRequest) GetRangeOk() (*Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *InsertNumberingRequest) SetRange(v Range)`

SetRange sets Range field to given value.

### HasRange

`func (o *InsertNumberingRequest) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetAbstractNumID

`func (o *InsertNumberingRequest) GetAbstractNumID() string`

GetAbstractNumID returns the AbstractNumID field if non-nil, zero value otherwise.

### GetAbstractNumIDOk

`func (o *InsertNumberingRequest) GetAbstractNumIDOk() (*string, bool)`

GetAbstractNumIDOk returns a tuple with the AbstractNumID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstractNumID

`func (o *InsertNumberingRequest) SetAbstractNumID(v string)`

SetAbstractNumID sets AbstractNumID field to given value.

### HasAbstractNumID

`func (o *InsertNumberingRequest) HasAbstractNumID() bool`

HasAbstractNumID returns a boolean if a field has been set.

### GetParagraphIndex

`func (o *InsertNumberingRequest) GetParagraphIndex() int32`

GetParagraphIndex returns the ParagraphIndex field if non-nil, zero value otherwise.

### GetParagraphIndexOk

`func (o *InsertNumberingRequest) GetParagraphIndexOk() (*int32, bool)`

GetParagraphIndexOk returns a tuple with the ParagraphIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParagraphIndex

`func (o *InsertNumberingRequest) SetParagraphIndex(v int32)`

SetParagraphIndex sets ParagraphIndex field to given value.

### HasParagraphIndex

`func (o *InsertNumberingRequest) HasParagraphIndex() bool`

HasParagraphIndex returns a boolean if a field has been set.

### GetPStyle

`func (o *InsertNumberingRequest) GetPStyle() ParagraphStyle`

GetPStyle returns the PStyle field if non-nil, zero value otherwise.

### GetPStyleOk

`func (o *InsertNumberingRequest) GetPStyleOk() (*ParagraphStyle, bool)`

GetPStyleOk returns a tuple with the PStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPStyle

`func (o *InsertNumberingRequest) SetPStyle(v ParagraphStyle)`

SetPStyle sets PStyle field to given value.

### HasPStyle

`func (o *InsertNumberingRequest) HasPStyle() bool`

HasPStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


