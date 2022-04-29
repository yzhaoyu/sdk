# DocAPIResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ret** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**GetDocFullTextRsp**](GetDocFullTextRsp.md) |  | [optional] 

## Methods

### NewDocAPIResponse2

`func NewDocAPIResponse2() *DocAPIResponse2`

NewDocAPIResponse2 instantiates a new DocAPIResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocAPIResponse2WithDefaults

`func NewDocAPIResponse2WithDefaults() *DocAPIResponse2`

NewDocAPIResponse2WithDefaults instantiates a new DocAPIResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRet

`func (o *DocAPIResponse2) GetRet() int32`

GetRet returns the Ret field if non-nil, zero value otherwise.

### GetRetOk

`func (o *DocAPIResponse2) GetRetOk() (*int32, bool)`

GetRetOk returns a tuple with the Ret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRet

`func (o *DocAPIResponse2) SetRet(v int32)`

SetRet sets Ret field to given value.

### HasRet

`func (o *DocAPIResponse2) HasRet() bool`

HasRet returns a boolean if a field has been set.

### GetMsg

`func (o *DocAPIResponse2) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *DocAPIResponse2) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *DocAPIResponse2) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *DocAPIResponse2) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *DocAPIResponse2) GetData() GetDocFullTextRsp`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DocAPIResponse2) GetDataOk() (*GetDocFullTextRsp, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DocAPIResponse2) SetData(v GetDocFullTextRsp)`

SetData sets Data field to given value.

### HasData

`func (o *DocAPIResponse2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


