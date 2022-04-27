# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsertText** | Pointer to [**InsertTextRequest**](InsertTextRequest.md) |  | [optional] 
**InsertImage** | Pointer to [**InsertImageRequest**](InsertImageRequest.md) |  | [optional] 
**InsertWebBlock** | Pointer to [**InsertWebBlockRequest**](InsertWebBlockRequest.md) |  | [optional] 

## Methods

### NewRequest

`func NewRequest() *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsertText

`func (o *Request) GetInsertText() InsertTextRequest`

GetInsertText returns the InsertText field if non-nil, zero value otherwise.

### GetInsertTextOk

`func (o *Request) GetInsertTextOk() (*InsertTextRequest, bool)`

GetInsertTextOk returns a tuple with the InsertText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertText

`func (o *Request) SetInsertText(v InsertTextRequest)`

SetInsertText sets InsertText field to given value.

### HasInsertText

`func (o *Request) HasInsertText() bool`

HasInsertText returns a boolean if a field has been set.

### GetInsertImage

`func (o *Request) GetInsertImage() InsertImageRequest`

GetInsertImage returns the InsertImage field if non-nil, zero value otherwise.

### GetInsertImageOk

`func (o *Request) GetInsertImageOk() (*InsertImageRequest, bool)`

GetInsertImageOk returns a tuple with the InsertImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertImage

`func (o *Request) SetInsertImage(v InsertImageRequest)`

SetInsertImage sets InsertImage field to given value.

### HasInsertImage

`func (o *Request) HasInsertImage() bool`

HasInsertImage returns a boolean if a field has been set.

### GetInsertWebBlock

`func (o *Request) GetInsertWebBlock() InsertWebBlockRequest`

GetInsertWebBlock returns the InsertWebBlock field if non-nil, zero value otherwise.

### GetInsertWebBlockOk

`func (o *Request) GetInsertWebBlockOk() (*InsertWebBlockRequest, bool)`

GetInsertWebBlockOk returns a tuple with the InsertWebBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertWebBlock

`func (o *Request) SetInsertWebBlock(v InsertWebBlockRequest)`

SetInsertWebBlock sets InsertWebBlock field to given value.

### HasInsertWebBlock

`func (o *Request) HasInsertWebBlock() bool`

HasInsertWebBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


