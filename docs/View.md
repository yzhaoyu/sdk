# View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**RecordIds** | Pointer to **[]string** |  | [optional] 
**FieldIds** | Pointer to **[]string** |  | [optional] 
**FieldMetas** | Pointer to [**map[string]ViewFieldMeta**](ViewFieldMeta.md) |  | [optional] 
**FilterSpec** | Pointer to [**FilterSpec**](FilterSpec.md) |  | [optional] 
**AutoSort** | Pointer to **bool** |  | [optional] 
**SortSpec** | Pointer to [**SortSpec**](SortSpec.md) |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PublicLevel** | Pointer to **int32** |  | [optional] 
**FrozenFieldsCount** | Pointer to **int32** |  | [optional] 
**GroupSpec** | Pointer to [**GroupSpec**](GroupSpec.md) |  | [optional] 
**RowHeightLevel** | Pointer to **int32** |  | [optional] 
**IsFieldStatEnabled** | Pointer to **bool** |  | [optional] 
**CardConfig** | Pointer to [**CardConfig**](CardConfig.md) |  | [optional] 

## Methods

### NewView

`func NewView() *View`

NewView instantiates a new View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewWithDefaults

`func NewViewWithDefaults() *View`

NewViewWithDefaults instantiates a new View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *View) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *View) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *View) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *View) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *View) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *View) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *View) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *View) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *View) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *View) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *View) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *View) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRecordIds

`func (o *View) GetRecordIds() []string`

GetRecordIds returns the RecordIds field if non-nil, zero value otherwise.

### GetRecordIdsOk

`func (o *View) GetRecordIdsOk() (*[]string, bool)`

GetRecordIdsOk returns a tuple with the RecordIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIds

`func (o *View) SetRecordIds(v []string)`

SetRecordIds sets RecordIds field to given value.

### HasRecordIds

`func (o *View) HasRecordIds() bool`

HasRecordIds returns a boolean if a field has been set.

### GetFieldIds

`func (o *View) GetFieldIds() []string`

GetFieldIds returns the FieldIds field if non-nil, zero value otherwise.

### GetFieldIdsOk

`func (o *View) GetFieldIdsOk() (*[]string, bool)`

GetFieldIdsOk returns a tuple with the FieldIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIds

`func (o *View) SetFieldIds(v []string)`

SetFieldIds sets FieldIds field to given value.

### HasFieldIds

`func (o *View) HasFieldIds() bool`

HasFieldIds returns a boolean if a field has been set.

### GetFieldMetas

`func (o *View) GetFieldMetas() map[string]ViewFieldMeta`

GetFieldMetas returns the FieldMetas field if non-nil, zero value otherwise.

### GetFieldMetasOk

`func (o *View) GetFieldMetasOk() (*map[string]ViewFieldMeta, bool)`

GetFieldMetasOk returns a tuple with the FieldMetas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMetas

`func (o *View) SetFieldMetas(v map[string]ViewFieldMeta)`

SetFieldMetas sets FieldMetas field to given value.

### HasFieldMetas

`func (o *View) HasFieldMetas() bool`

HasFieldMetas returns a boolean if a field has been set.

### GetFilterSpec

`func (o *View) GetFilterSpec() FilterSpec`

GetFilterSpec returns the FilterSpec field if non-nil, zero value otherwise.

### GetFilterSpecOk

`func (o *View) GetFilterSpecOk() (*FilterSpec, bool)`

GetFilterSpecOk returns a tuple with the FilterSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSpec

`func (o *View) SetFilterSpec(v FilterSpec)`

SetFilterSpec sets FilterSpec field to given value.

### HasFilterSpec

`func (o *View) HasFilterSpec() bool`

HasFilterSpec returns a boolean if a field has been set.

### GetAutoSort

`func (o *View) GetAutoSort() bool`

GetAutoSort returns the AutoSort field if non-nil, zero value otherwise.

### GetAutoSortOk

`func (o *View) GetAutoSortOk() (*bool, bool)`

GetAutoSortOk returns a tuple with the AutoSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSort

`func (o *View) SetAutoSort(v bool)`

SetAutoSort sets AutoSort field to given value.

### HasAutoSort

`func (o *View) HasAutoSort() bool`

HasAutoSort returns a boolean if a field has been set.

### GetSortSpec

`func (o *View) GetSortSpec() SortSpec`

GetSortSpec returns the SortSpec field if non-nil, zero value otherwise.

### GetSortSpecOk

`func (o *View) GetSortSpecOk() (*SortSpec, bool)`

GetSortSpecOk returns a tuple with the SortSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortSpec

`func (o *View) SetSortSpec(v SortSpec)`

SetSortSpec sets SortSpec field to given value.

### HasSortSpec

`func (o *View) HasSortSpec() bool`

HasSortSpec returns a boolean if a field has been set.

### GetOwner

`func (o *View) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *View) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *View) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *View) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPublicLevel

`func (o *View) GetPublicLevel() int32`

GetPublicLevel returns the PublicLevel field if non-nil, zero value otherwise.

### GetPublicLevelOk

`func (o *View) GetPublicLevelOk() (*int32, bool)`

GetPublicLevelOk returns a tuple with the PublicLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLevel

`func (o *View) SetPublicLevel(v int32)`

SetPublicLevel sets PublicLevel field to given value.

### HasPublicLevel

`func (o *View) HasPublicLevel() bool`

HasPublicLevel returns a boolean if a field has been set.

### GetFrozenFieldsCount

`func (o *View) GetFrozenFieldsCount() int32`

GetFrozenFieldsCount returns the FrozenFieldsCount field if non-nil, zero value otherwise.

### GetFrozenFieldsCountOk

`func (o *View) GetFrozenFieldsCountOk() (*int32, bool)`

GetFrozenFieldsCountOk returns a tuple with the FrozenFieldsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenFieldsCount

`func (o *View) SetFrozenFieldsCount(v int32)`

SetFrozenFieldsCount sets FrozenFieldsCount field to given value.

### HasFrozenFieldsCount

`func (o *View) HasFrozenFieldsCount() bool`

HasFrozenFieldsCount returns a boolean if a field has been set.

### GetGroupSpec

`func (o *View) GetGroupSpec() GroupSpec`

GetGroupSpec returns the GroupSpec field if non-nil, zero value otherwise.

### GetGroupSpecOk

`func (o *View) GetGroupSpecOk() (*GroupSpec, bool)`

GetGroupSpecOk returns a tuple with the GroupSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSpec

`func (o *View) SetGroupSpec(v GroupSpec)`

SetGroupSpec sets GroupSpec field to given value.

### HasGroupSpec

`func (o *View) HasGroupSpec() bool`

HasGroupSpec returns a boolean if a field has been set.

### GetRowHeightLevel

`func (o *View) GetRowHeightLevel() int32`

GetRowHeightLevel returns the RowHeightLevel field if non-nil, zero value otherwise.

### GetRowHeightLevelOk

`func (o *View) GetRowHeightLevelOk() (*int32, bool)`

GetRowHeightLevelOk returns a tuple with the RowHeightLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowHeightLevel

`func (o *View) SetRowHeightLevel(v int32)`

SetRowHeightLevel sets RowHeightLevel field to given value.

### HasRowHeightLevel

`func (o *View) HasRowHeightLevel() bool`

HasRowHeightLevel returns a boolean if a field has been set.

### GetIsFieldStatEnabled

`func (o *View) GetIsFieldStatEnabled() bool`

GetIsFieldStatEnabled returns the IsFieldStatEnabled field if non-nil, zero value otherwise.

### GetIsFieldStatEnabledOk

`func (o *View) GetIsFieldStatEnabledOk() (*bool, bool)`

GetIsFieldStatEnabledOk returns a tuple with the IsFieldStatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFieldStatEnabled

`func (o *View) SetIsFieldStatEnabled(v bool)`

SetIsFieldStatEnabled sets IsFieldStatEnabled field to given value.

### HasIsFieldStatEnabled

`func (o *View) HasIsFieldStatEnabled() bool`

HasIsFieldStatEnabled returns a boolean if a field has been set.

### GetCardConfig

`func (o *View) GetCardConfig() CardConfig`

GetCardConfig returns the CardConfig field if non-nil, zero value otherwise.

### GetCardConfigOk

`func (o *View) GetCardConfigOk() (*CardConfig, bool)`

GetCardConfigOk returns a tuple with the CardConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardConfig

`func (o *View) SetCardConfig(v CardConfig)`

SetCardConfig sets CardConfig field to given value.

### HasCardConfig

`func (o *View) HasCardConfig() bool`

HasCardConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


