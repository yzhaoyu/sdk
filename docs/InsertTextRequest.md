# InsertTextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**Location** | [**Location**](Location.md) |  | 

## Methods

### NewInsertTextRequest

`func NewInsertTextRequest(text string, location Location, ) *InsertTextRequest`

NewInsertTextRequest instantiates a new InsertTextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsertTextRequestWithDefaults

`func NewInsertTextRequestWithDefaults() *InsertTextRequest`

NewInsertTextRequestWithDefaults instantiates a new InsertTextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *InsertTextRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *InsertTextRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *InsertTextRequest) SetText(v string)`

SetText sets Text field to given value.


### GetLocation

`func (o *InsertTextRequest) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InsertTextRequest) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InsertTextRequest) SetLocation(v Location)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


