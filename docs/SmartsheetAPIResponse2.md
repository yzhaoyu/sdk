# SmartsheetAPIResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ret** | **int32** |  | 
**Msg** | **string** |  | 
**Data** | Pointer to [**GetSubSheetResponse**](GetSubSheetResponse.md) |  | [optional] 

## Methods

### NewSmartsheetAPIResponse2

`func NewSmartsheetAPIResponse2(ret int32, msg string, ) *SmartsheetAPIResponse2`

NewSmartsheetAPIResponse2 instantiates a new SmartsheetAPIResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetAPIResponse2WithDefaults

`func NewSmartsheetAPIResponse2WithDefaults() *SmartsheetAPIResponse2`

NewSmartsheetAPIResponse2WithDefaults instantiates a new SmartsheetAPIResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRet

`func (o *SmartsheetAPIResponse2) GetRet() int32`

GetRet returns the Ret field if non-nil, zero value otherwise.

### GetRetOk

`func (o *SmartsheetAPIResponse2) GetRetOk() (*int32, bool)`

GetRetOk returns a tuple with the Ret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRet

`func (o *SmartsheetAPIResponse2) SetRet(v int32)`

SetRet sets Ret field to given value.


### GetMsg

`func (o *SmartsheetAPIResponse2) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SmartsheetAPIResponse2) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SmartsheetAPIResponse2) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *SmartsheetAPIResponse2) GetData() GetSubSheetResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SmartsheetAPIResponse2) GetDataOk() (*GetSubSheetResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SmartsheetAPIResponse2) SetData(v GetSubSheetResponse)`

SetData sets Data field to given value.

### HasData

`func (o *SmartsheetAPIResponse2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


