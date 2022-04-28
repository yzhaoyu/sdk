# UserResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ret** | **int32** |  | 
**Msg** | **string** |  | 
**Data** | Pointer to [**Response1**](Response1.md) |  | [optional] 

## Methods

### NewUserResponse1

`func NewUserResponse1(ret int32, msg string, ) *UserResponse1`

NewUserResponse1 instantiates a new UserResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponse1WithDefaults

`func NewUserResponse1WithDefaults() *UserResponse1`

NewUserResponse1WithDefaults instantiates a new UserResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRet

`func (o *UserResponse1) GetRet() int32`

GetRet returns the Ret field if non-nil, zero value otherwise.

### GetRetOk

`func (o *UserResponse1) GetRetOk() (*int32, bool)`

GetRetOk returns a tuple with the Ret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRet

`func (o *UserResponse1) SetRet(v int32)`

SetRet sets Ret field to given value.


### GetMsg

`func (o *UserResponse1) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UserResponse1) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UserResponse1) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *UserResponse1) GetData() Response1`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserResponse1) GetDataOk() (*Response1, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserResponse1) SetData(v Response1)`

SetData sets Data field to given value.

### HasData

`func (o *UserResponse1) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


