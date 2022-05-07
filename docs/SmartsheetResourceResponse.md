# SmartsheetResourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GetViews** | Pointer to [**SmartsheetGetViewsResponse**](SmartsheetGetViewsResponse.md) |  | [optional] 
**GetRecords** | Pointer to [**SmartsheetGetRecordsRsp**](SmartsheetGetRecordsRsp.md) |  | [optional] 
**GetFields** | Pointer to [**SmartsheetGetFieldsResponse**](SmartsheetGetFieldsResponse.md) |  | [optional] 
**AddView** | Pointer to [**SmartsheetAddViewResponse**](SmartsheetAddViewResponse.md) |  | [optional] 
**DeleteViews** | Pointer to **map[string]interface{}** | 删除指定表格指定子表中的指定视图回包 | [optional] 
**AddRecords** | Pointer to [**SmartsheetAddRecordsResponse**](SmartsheetAddRecordsResponse.md) |  | [optional] 
**UpdateRecords** | Pointer to [**SmartsheetUpdateRecordsResponse**](SmartsheetUpdateRecordsResponse.md) |  | [optional] 
**DeleteRecords** | Pointer to **map[string]interface{}** | 删除指定表格和子表中的一到多条记录回包 | [optional] 
**AddFields** | Pointer to [**SmartsheetAddFieldsResponse**](SmartsheetAddFieldsResponse.md) |  | [optional] 
**UpdateFields** | Pointer to [**SmartsheetUpdateFieldsResponse**](SmartsheetUpdateFieldsResponse.md) |  | [optional] 
**DeleteFields** | Pointer to **map[string]interface{}** | 删除字段回包 | [optional] 

## Methods

### NewSmartsheetResourceResponse

`func NewSmartsheetResourceResponse() *SmartsheetResourceResponse`

NewSmartsheetResourceResponse instantiates a new SmartsheetResourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetResourceResponseWithDefaults

`func NewSmartsheetResourceResponseWithDefaults() *SmartsheetResourceResponse`

NewSmartsheetResourceResponseWithDefaults instantiates a new SmartsheetResourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGetViews

`func (o *SmartsheetResourceResponse) GetGetViews() SmartsheetGetViewsResponse`

GetGetViews returns the GetViews field if non-nil, zero value otherwise.

### GetGetViewsOk

`func (o *SmartsheetResourceResponse) GetGetViewsOk() (*SmartsheetGetViewsResponse, bool)`

GetGetViewsOk returns a tuple with the GetViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetViews

`func (o *SmartsheetResourceResponse) SetGetViews(v SmartsheetGetViewsResponse)`

SetGetViews sets GetViews field to given value.

### HasGetViews

`func (o *SmartsheetResourceResponse) HasGetViews() bool`

HasGetViews returns a boolean if a field has been set.

### GetGetRecords

`func (o *SmartsheetResourceResponse) GetGetRecords() SmartsheetGetRecordsRsp`

GetGetRecords returns the GetRecords field if non-nil, zero value otherwise.

### GetGetRecordsOk

`func (o *SmartsheetResourceResponse) GetGetRecordsOk() (*SmartsheetGetRecordsRsp, bool)`

GetGetRecordsOk returns a tuple with the GetRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetRecords

`func (o *SmartsheetResourceResponse) SetGetRecords(v SmartsheetGetRecordsRsp)`

SetGetRecords sets GetRecords field to given value.

### HasGetRecords

`func (o *SmartsheetResourceResponse) HasGetRecords() bool`

HasGetRecords returns a boolean if a field has been set.

### GetGetFields

`func (o *SmartsheetResourceResponse) GetGetFields() SmartsheetGetFieldsResponse`

GetGetFields returns the GetFields field if non-nil, zero value otherwise.

### GetGetFieldsOk

`func (o *SmartsheetResourceResponse) GetGetFieldsOk() (*SmartsheetGetFieldsResponse, bool)`

GetGetFieldsOk returns a tuple with the GetFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetFields

`func (o *SmartsheetResourceResponse) SetGetFields(v SmartsheetGetFieldsResponse)`

SetGetFields sets GetFields field to given value.

### HasGetFields

`func (o *SmartsheetResourceResponse) HasGetFields() bool`

HasGetFields returns a boolean if a field has been set.

### GetAddView

`func (o *SmartsheetResourceResponse) GetAddView() SmartsheetAddViewResponse`

GetAddView returns the AddView field if non-nil, zero value otherwise.

### GetAddViewOk

`func (o *SmartsheetResourceResponse) GetAddViewOk() (*SmartsheetAddViewResponse, bool)`

GetAddViewOk returns a tuple with the AddView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddView

`func (o *SmartsheetResourceResponse) SetAddView(v SmartsheetAddViewResponse)`

SetAddView sets AddView field to given value.

### HasAddView

`func (o *SmartsheetResourceResponse) HasAddView() bool`

HasAddView returns a boolean if a field has been set.

### GetDeleteViews

`func (o *SmartsheetResourceResponse) GetDeleteViews() map[string]interface{}`

GetDeleteViews returns the DeleteViews field if non-nil, zero value otherwise.

### GetDeleteViewsOk

`func (o *SmartsheetResourceResponse) GetDeleteViewsOk() (*map[string]interface{}, bool)`

GetDeleteViewsOk returns a tuple with the DeleteViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteViews

`func (o *SmartsheetResourceResponse) SetDeleteViews(v map[string]interface{})`

SetDeleteViews sets DeleteViews field to given value.

### HasDeleteViews

`func (o *SmartsheetResourceResponse) HasDeleteViews() bool`

HasDeleteViews returns a boolean if a field has been set.

### GetAddRecords

`func (o *SmartsheetResourceResponse) GetAddRecords() SmartsheetAddRecordsResponse`

GetAddRecords returns the AddRecords field if non-nil, zero value otherwise.

### GetAddRecordsOk

`func (o *SmartsheetResourceResponse) GetAddRecordsOk() (*SmartsheetAddRecordsResponse, bool)`

GetAddRecordsOk returns a tuple with the AddRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRecords

`func (o *SmartsheetResourceResponse) SetAddRecords(v SmartsheetAddRecordsResponse)`

SetAddRecords sets AddRecords field to given value.

### HasAddRecords

`func (o *SmartsheetResourceResponse) HasAddRecords() bool`

HasAddRecords returns a boolean if a field has been set.

### GetUpdateRecords

`func (o *SmartsheetResourceResponse) GetUpdateRecords() SmartsheetUpdateRecordsResponse`

GetUpdateRecords returns the UpdateRecords field if non-nil, zero value otherwise.

### GetUpdateRecordsOk

`func (o *SmartsheetResourceResponse) GetUpdateRecordsOk() (*SmartsheetUpdateRecordsResponse, bool)`

GetUpdateRecordsOk returns a tuple with the UpdateRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRecords

`func (o *SmartsheetResourceResponse) SetUpdateRecords(v SmartsheetUpdateRecordsResponse)`

SetUpdateRecords sets UpdateRecords field to given value.

### HasUpdateRecords

`func (o *SmartsheetResourceResponse) HasUpdateRecords() bool`

HasUpdateRecords returns a boolean if a field has been set.

### GetDeleteRecords

`func (o *SmartsheetResourceResponse) GetDeleteRecords() map[string]interface{}`

GetDeleteRecords returns the DeleteRecords field if non-nil, zero value otherwise.

### GetDeleteRecordsOk

`func (o *SmartsheetResourceResponse) GetDeleteRecordsOk() (*map[string]interface{}, bool)`

GetDeleteRecordsOk returns a tuple with the DeleteRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRecords

`func (o *SmartsheetResourceResponse) SetDeleteRecords(v map[string]interface{})`

SetDeleteRecords sets DeleteRecords field to given value.

### HasDeleteRecords

`func (o *SmartsheetResourceResponse) HasDeleteRecords() bool`

HasDeleteRecords returns a boolean if a field has been set.

### GetAddFields

`func (o *SmartsheetResourceResponse) GetAddFields() SmartsheetAddFieldsResponse`

GetAddFields returns the AddFields field if non-nil, zero value otherwise.

### GetAddFieldsOk

`func (o *SmartsheetResourceResponse) GetAddFieldsOk() (*SmartsheetAddFieldsResponse, bool)`

GetAddFieldsOk returns a tuple with the AddFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFields

`func (o *SmartsheetResourceResponse) SetAddFields(v SmartsheetAddFieldsResponse)`

SetAddFields sets AddFields field to given value.

### HasAddFields

`func (o *SmartsheetResourceResponse) HasAddFields() bool`

HasAddFields returns a boolean if a field has been set.

### GetUpdateFields

`func (o *SmartsheetResourceResponse) GetUpdateFields() SmartsheetUpdateFieldsResponse`

GetUpdateFields returns the UpdateFields field if non-nil, zero value otherwise.

### GetUpdateFieldsOk

`func (o *SmartsheetResourceResponse) GetUpdateFieldsOk() (*SmartsheetUpdateFieldsResponse, bool)`

GetUpdateFieldsOk returns a tuple with the UpdateFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFields

`func (o *SmartsheetResourceResponse) SetUpdateFields(v SmartsheetUpdateFieldsResponse)`

SetUpdateFields sets UpdateFields field to given value.

### HasUpdateFields

`func (o *SmartsheetResourceResponse) HasUpdateFields() bool`

HasUpdateFields returns a boolean if a field has been set.

### GetDeleteFields

`func (o *SmartsheetResourceResponse) GetDeleteFields() map[string]interface{}`

GetDeleteFields returns the DeleteFields field if non-nil, zero value otherwise.

### GetDeleteFieldsOk

`func (o *SmartsheetResourceResponse) GetDeleteFieldsOk() (*map[string]interface{}, bool)`

GetDeleteFieldsOk returns a tuple with the DeleteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFields

`func (o *SmartsheetResourceResponse) SetDeleteFields(v map[string]interface{})`

SetDeleteFields sets DeleteFields field to given value.

### HasDeleteFields

`func (o *SmartsheetResourceResponse) HasDeleteFields() bool`

HasDeleteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


