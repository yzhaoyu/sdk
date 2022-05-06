# DriveAPIResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ret** | **int32** |  | 
**Msg** | **string** |  | 
**Data** | Pointer to [**FileInfo**](FileInfo.md) |  | [optional] 

## Methods

### NewDriveAPIResponse1

`func NewDriveAPIResponse1(ret int32, msg string, ) *DriveAPIResponse1`

NewDriveAPIResponse1 instantiates a new DriveAPIResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveAPIResponse1WithDefaults

`func NewDriveAPIResponse1WithDefaults() *DriveAPIResponse1`

NewDriveAPIResponse1WithDefaults instantiates a new DriveAPIResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRet

`func (o *DriveAPIResponse1) GetRet() int32`

GetRet returns the Ret field if non-nil, zero value otherwise.

### GetRetOk

`func (o *DriveAPIResponse1) GetRetOk() (*int32, bool)`

GetRetOk returns a tuple with the Ret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRet

`func (o *DriveAPIResponse1) SetRet(v int32)`

SetRet sets Ret field to given value.


### GetMsg

`func (o *DriveAPIResponse1) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *DriveAPIResponse1) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *DriveAPIResponse1) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *DriveAPIResponse1) GetData() FileInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DriveAPIResponse1) GetDataOk() (*FileInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DriveAPIResponse1) SetData(v FileInfo)`

SetData sets Data field to given value.

### HasData

`func (o *DriveAPIResponse1) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


