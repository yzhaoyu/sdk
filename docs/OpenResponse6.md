# OpenResponse6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ret** | **int32** |  | 
**Msg** | **string** |  | 
**Data** | Pointer to [**GetSubSheetResponse**](GetSubSheetResponse.md) |  | [optional] 

## Methods

### NewOpenResponse6

`func NewOpenResponse6(ret int32, msg string, ) *OpenResponse6`

NewOpenResponse6 instantiates a new OpenResponse6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenResponse6WithDefaults

`func NewOpenResponse6WithDefaults() *OpenResponse6`

NewOpenResponse6WithDefaults instantiates a new OpenResponse6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRet

`func (o *OpenResponse6) GetRet() int32`

GetRet returns the Ret field if non-nil, zero value otherwise.

### GetRetOk

`func (o *OpenResponse6) GetRetOk() (*int32, bool)`

GetRetOk returns a tuple with the Ret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRet

`func (o *OpenResponse6) SetRet(v int32)`

SetRet sets Ret field to given value.


### GetMsg

`func (o *OpenResponse6) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OpenResponse6) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OpenResponse6) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *OpenResponse6) GetData() GetSubSheetResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OpenResponse6) GetDataOk() (*GetSubSheetResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OpenResponse6) SetData(v GetSubSheetResponse)`

SetData sets Data field to given value.

### HasData

`func (o *OpenResponse6) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


