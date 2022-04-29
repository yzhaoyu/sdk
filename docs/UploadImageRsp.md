# UploadImageRsp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ret** | **int32** |  | 
**Msg** | **string** |  | 
**Data** | Pointer to [**Response3**](Response3.md) |  | [optional] 

## Methods

### NewUploadImageRsp

`func NewUploadImageRsp(ret int32, msg string, ) *UploadImageRsp`

NewUploadImageRsp instantiates a new UploadImageRsp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadImageRspWithDefaults

`func NewUploadImageRspWithDefaults() *UploadImageRsp`

NewUploadImageRspWithDefaults instantiates a new UploadImageRsp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRet

`func (o *UploadImageRsp) GetRet() int32`

GetRet returns the Ret field if non-nil, zero value otherwise.

### GetRetOk

`func (o *UploadImageRsp) GetRetOk() (*int32, bool)`

GetRetOk returns a tuple with the Ret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRet

`func (o *UploadImageRsp) SetRet(v int32)`

SetRet sets Ret field to given value.


### GetMsg

`func (o *UploadImageRsp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UploadImageRsp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UploadImageRsp) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *UploadImageRsp) GetData() Response3`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UploadImageRsp) GetDataOk() (*Response3, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UploadImageRsp) SetData(v Response3)`

SetData sets Data field to given value.

### HasData

`func (o *UploadImageRsp) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


