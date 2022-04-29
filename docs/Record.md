# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to [**map[string]CellValue**](CellValue.md) |  | [optional] 
**CommentIds** | Pointer to **[]string** |  | [optional] 
**CreatedUserId** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewRecord

`func NewRecord() *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *Record) GetValues() map[string]CellValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Record) GetValuesOk() (*map[string]CellValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Record) SetValues(v map[string]CellValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *Record) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetCommentIds

`func (o *Record) GetCommentIds() []string`

GetCommentIds returns the CommentIds field if non-nil, zero value otherwise.

### GetCommentIdsOk

`func (o *Record) GetCommentIdsOk() (*[]string, bool)`

GetCommentIdsOk returns a tuple with the CommentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentIds

`func (o *Record) SetCommentIds(v []string)`

SetCommentIds sets CommentIds field to given value.

### HasCommentIds

`func (o *Record) HasCommentIds() bool`

HasCommentIds returns a boolean if a field has been set.

### GetCreatedUserId

`func (o *Record) GetCreatedUserId() string`

GetCreatedUserId returns the CreatedUserId field if non-nil, zero value otherwise.

### GetCreatedUserIdOk

`func (o *Record) GetCreatedUserIdOk() (*string, bool)`

GetCreatedUserIdOk returns a tuple with the CreatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUserId

`func (o *Record) SetCreatedUserId(v string)`

SetCreatedUserId sets CreatedUserId field to given value.

### HasCreatedUserId

`func (o *Record) HasCreatedUserId() bool`

HasCreatedUserId returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Record) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Record) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Record) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Record) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


