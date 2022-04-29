# RecordResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordID** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**Values** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRecordResource

`func NewRecordResource() *RecordResource`

NewRecordResource instantiates a new RecordResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordResourceWithDefaults

`func NewRecordResourceWithDefaults() *RecordResource`

NewRecordResourceWithDefaults instantiates a new RecordResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordID

`func (o *RecordResource) GetRecordID() string`

GetRecordID returns the RecordID field if non-nil, zero value otherwise.

### GetRecordIDOk

`func (o *RecordResource) GetRecordIDOk() (*string, bool)`

GetRecordIDOk returns a tuple with the RecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordID

`func (o *RecordResource) SetRecordID(v string)`

SetRecordID sets RecordID field to given value.

### HasRecordID

`func (o *RecordResource) HasRecordID() bool`

HasRecordID returns a boolean if a field has been set.

### GetCreateTime

`func (o *RecordResource) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *RecordResource) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *RecordResource) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *RecordResource) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *RecordResource) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *RecordResource) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *RecordResource) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *RecordResource) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetValues

`func (o *RecordResource) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RecordResource) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RecordResource) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *RecordResource) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


