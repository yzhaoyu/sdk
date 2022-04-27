# BatchUpdateDocDataReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileID** | **string** |  | 
**Requests** | [**[]Request**](Request.md) |  | 

## Methods

### NewBatchUpdateDocDataReq

`func NewBatchUpdateDocDataReq(fileID string, requests []Request, ) *BatchUpdateDocDataReq`

NewBatchUpdateDocDataReq instantiates a new BatchUpdateDocDataReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateDocDataReqWithDefaults

`func NewBatchUpdateDocDataReqWithDefaults() *BatchUpdateDocDataReq`

NewBatchUpdateDocDataReqWithDefaults instantiates a new BatchUpdateDocDataReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileID

`func (o *BatchUpdateDocDataReq) GetFileID() string`

GetFileID returns the FileID field if non-nil, zero value otherwise.

### GetFileIDOk

`func (o *BatchUpdateDocDataReq) GetFileIDOk() (*string, bool)`

GetFileIDOk returns a tuple with the FileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileID

`func (o *BatchUpdateDocDataReq) SetFileID(v string)`

SetFileID sets FileID field to given value.


### GetRequests

`func (o *BatchUpdateDocDataReq) GetRequests() []Request`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *BatchUpdateDocDataReq) GetRequestsOk() (*[]Request, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *BatchUpdateDocDataReq) SetRequests(v []Request)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


