# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Begin** | Pointer to **int32** |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Children** | Pointer to [**[]Node**](Node.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Property** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewNode

`func NewNode() *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBegin

`func (o *Node) GetBegin() int32`

GetBegin returns the Begin field if non-nil, zero value otherwise.

### GetBeginOk

`func (o *Node) GetBeginOk() (*int32, bool)`

GetBeginOk returns a tuple with the Begin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegin

`func (o *Node) SetBegin(v int32)`

SetBegin sets Begin field to given value.

### HasBegin

`func (o *Node) HasBegin() bool`

HasBegin returns a boolean if a field has been set.

### GetEnd

`func (o *Node) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Node) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Node) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Node) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetType

`func (o *Node) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Node) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Node) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Node) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChildren

`func (o *Node) GetChildren() []Node`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Node) GetChildrenOk() (*[]Node, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Node) SetChildren(v []Node)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Node) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetText

`func (o *Node) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Node) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Node) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Node) HasText() bool`

HasText returns a boolean if a field has been set.

### GetProperty

`func (o *Node) GetProperty() map[string]interface{}`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *Node) GetPropertyOk() (*map[string]interface{}, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *Node) SetProperty(v map[string]interface{})`

SetProperty sets Property field to given value.

### HasProperty

`func (o *Node) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


