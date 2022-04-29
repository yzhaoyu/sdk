# GetDocFullTextRsp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to [**Node**](Node.md) |  | [optional] 

## Methods

### NewGetDocFullTextRsp

`func NewGetDocFullTextRsp() *GetDocFullTextRsp`

NewGetDocFullTextRsp instantiates a new GetDocFullTextRsp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocFullTextRspWithDefaults

`func NewGetDocFullTextRspWithDefaults() *GetDocFullTextRsp`

NewGetDocFullTextRspWithDefaults instantiates a new GetDocFullTextRsp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *GetDocFullTextRsp) GetDocument() Node`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *GetDocFullTextRsp) GetDocumentOk() (*Node, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *GetDocFullTextRsp) SetDocument(v Node)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *GetDocFullTextRsp) HasDocument() bool`

HasDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


