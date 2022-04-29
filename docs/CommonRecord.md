# CommonRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordID** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCommonRecord

`func NewCommonRecord() *CommonRecord`

NewCommonRecord instantiates a new CommonRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonRecordWithDefaults

`func NewCommonRecordWithDefaults() *CommonRecord`

NewCommonRecordWithDefaults instantiates a new CommonRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordID

`func (o *CommonRecord) GetRecordID() string`

GetRecordID returns the RecordID field if non-nil, zero value otherwise.

### GetRecordIDOk

`func (o *CommonRecord) GetRecordIDOk() (*string, bool)`

GetRecordIDOk returns a tuple with the RecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordID

`func (o *CommonRecord) SetRecordID(v string)`

SetRecordID sets RecordID field to given value.

### HasRecordID

`func (o *CommonRecord) HasRecordID() bool`

HasRecordID returns a boolean if a field has been set.

### GetValues

`func (o *CommonRecord) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CommonRecord) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CommonRecord) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *CommonRecord) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


