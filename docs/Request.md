# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplaceText** | Pointer to [**ReplaceTextRequest**](ReplaceTextRequest.md) |  | [optional] 
**InsertText** | Pointer to [**InsertTextRequest**](InsertTextRequest.md) |  | [optional] 
**DeleteContent** | Pointer to [**DeleteContentRequest**](DeleteContentRequest.md) |  | [optional] 
**InsertImage** | Pointer to [**InsertImageRequest**](InsertImageRequest.md) |  | [optional] 
**InsertPageBreak** | Pointer to [**InsertPageBreakRequest**](InsertPageBreakRequest.md) |  | [optional] 
**InsertTable** | Pointer to [**InsertTableRequest**](InsertTableRequest.md) |  | [optional] 
**InsertNumbering** | Pointer to [**InsertNumberingRequest**](InsertNumberingRequest.md) |  | [optional] 
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

### GetReplaceText

`func (o *Request) GetReplaceText() ReplaceTextRequest`

GetReplaceText returns the ReplaceText field if non-nil, zero value otherwise.

### GetReplaceTextOk

`func (o *Request) GetReplaceTextOk() (*ReplaceTextRequest, bool)`

GetReplaceTextOk returns a tuple with the ReplaceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceText

`func (o *Request) SetReplaceText(v ReplaceTextRequest)`

SetReplaceText sets ReplaceText field to given value.

### HasReplaceText

`func (o *Request) HasReplaceText() bool`

HasReplaceText returns a boolean if a field has been set.

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

### GetDeleteContent

`func (o *Request) GetDeleteContent() DeleteContentRequest`

GetDeleteContent returns the DeleteContent field if non-nil, zero value otherwise.

### GetDeleteContentOk

`func (o *Request) GetDeleteContentOk() (*DeleteContentRequest, bool)`

GetDeleteContentOk returns a tuple with the DeleteContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteContent

`func (o *Request) SetDeleteContent(v DeleteContentRequest)`

SetDeleteContent sets DeleteContent field to given value.

### HasDeleteContent

`func (o *Request) HasDeleteContent() bool`

HasDeleteContent returns a boolean if a field has been set.

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

### GetInsertPageBreak

`func (o *Request) GetInsertPageBreak() InsertPageBreakRequest`

GetInsertPageBreak returns the InsertPageBreak field if non-nil, zero value otherwise.

### GetInsertPageBreakOk

`func (o *Request) GetInsertPageBreakOk() (*InsertPageBreakRequest, bool)`

GetInsertPageBreakOk returns a tuple with the InsertPageBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertPageBreak

`func (o *Request) SetInsertPageBreak(v InsertPageBreakRequest)`

SetInsertPageBreak sets InsertPageBreak field to given value.

### HasInsertPageBreak

`func (o *Request) HasInsertPageBreak() bool`

HasInsertPageBreak returns a boolean if a field has been set.

### GetInsertTable

`func (o *Request) GetInsertTable() InsertTableRequest`

GetInsertTable returns the InsertTable field if non-nil, zero value otherwise.

### GetInsertTableOk

`func (o *Request) GetInsertTableOk() (*InsertTableRequest, bool)`

GetInsertTableOk returns a tuple with the InsertTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertTable

`func (o *Request) SetInsertTable(v InsertTableRequest)`

SetInsertTable sets InsertTable field to given value.

### HasInsertTable

`func (o *Request) HasInsertTable() bool`

HasInsertTable returns a boolean if a field has been set.

### GetInsertNumbering

`func (o *Request) GetInsertNumbering() InsertNumberingRequest`

GetInsertNumbering returns the InsertNumbering field if non-nil, zero value otherwise.

### GetInsertNumberingOk

`func (o *Request) GetInsertNumberingOk() (*InsertNumberingRequest, bool)`

GetInsertNumberingOk returns a tuple with the InsertNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertNumbering

`func (o *Request) SetInsertNumbering(v InsertNumberingRequest)`

SetInsertNumbering sets InsertNumbering field to given value.

### HasInsertNumbering

`func (o *Request) HasInsertNumbering() bool`

HasInsertNumbering returns a boolean if a field has been set.

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


