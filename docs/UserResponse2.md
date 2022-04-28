# UserResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ret** | **int32** |  | 
**Msg** | **string** |  | 
**Data** | Pointer to [**Response2**](Response2.md) |  | [optional] 

## Methods

### NewUserResponse2

`func NewUserResponse2(ret int32, msg string, ) *UserResponse2`

NewUserResponse2 instantiates a new UserResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponse2WithDefaults

`func NewUserResponse2WithDefaults() *UserResponse2`

NewUserResponse2WithDefaults instantiates a new UserResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRet

`func (o *UserResponse2) GetRet() int32`

GetRet returns the Ret field if non-nil, zero value otherwise.

### GetRetOk

`func (o *UserResponse2) GetRetOk() (*int32, bool)`

GetRetOk returns a tuple with the Ret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRet

`func (o *UserResponse2) SetRet(v int32)`

SetRet sets Ret field to given value.


### GetMsg

`func (o *UserResponse2) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UserResponse2) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UserResponse2) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *UserResponse2) GetData() Response2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserResponse2) GetDataOk() (*Response2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserResponse2) SetData(v Response2)`

SetData sets Data field to given value.

### HasData

`func (o *UserResponse2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


