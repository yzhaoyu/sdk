/*
Tencent Docs SDK

Tencent docs sdk contains DocAPI, SmartsheetAPI, SheetAPI, DriveAPI and OAuthAPI

API version: 0.0.1
Contact: tencentdocs@tencent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// View View 定义多维表的视图结构，一个多维表可以拥有多个视图。
type View struct {
	Id *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *int32 `json:"type,omitempty"`
	RecordIds []string `json:"recordIds,omitempty"`
	FieldIds []string `json:"fieldIds,omitempty"`
	FieldMetas *map[string]ViewFieldMeta `json:"fieldMetas,omitempty"`
	FilterSpec *FilterSpec `json:"filterSpec,omitempty"`
	AutoSort *bool `json:"autoSort,omitempty"`
	SortSpec *SortSpec `json:"sortSpec,omitempty"`
	Owner *string `json:"owner,omitempty"`
	PublicLevel *int32 `json:"publicLevel,omitempty"`
	FrozenFieldsCount *int32 `json:"frozenFieldsCount,omitempty"`
	GroupSpec *GroupSpec `json:"groupSpec,omitempty"`
	RowHeightLevel *int32 `json:"rowHeightLevel,omitempty"`
	IsFieldStatEnabled *bool `json:"isFieldStatEnabled,omitempty"`
	CardConfig *CardConfig `json:"cardConfig,omitempty"`
}

// NewView instantiates a new View object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewView() *View {
	this := View{}
	return &this
}

// NewViewWithDefaults instantiates a new View object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewWithDefaults() *View {
	this := View{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *View) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *View) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *View) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *View) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *View) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *View) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *View) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *View) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *View) SetType(v int32) {
	o.Type = &v
}

// GetRecordIds returns the RecordIds field value if set, zero value otherwise.
func (o *View) GetRecordIds() []string {
	if o == nil || o.RecordIds == nil {
		var ret []string
		return ret
	}
	return o.RecordIds
}

// GetRecordIdsOk returns a tuple with the RecordIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetRecordIdsOk() ([]string, bool) {
	if o == nil || o.RecordIds == nil {
		return nil, false
	}
	return o.RecordIds, true
}

// HasRecordIds returns a boolean if a field has been set.
func (o *View) HasRecordIds() bool {
	if o != nil && o.RecordIds != nil {
		return true
	}

	return false
}

// SetRecordIds gets a reference to the given []string and assigns it to the RecordIds field.
func (o *View) SetRecordIds(v []string) {
	o.RecordIds = v
}

// GetFieldIds returns the FieldIds field value if set, zero value otherwise.
func (o *View) GetFieldIds() []string {
	if o == nil || o.FieldIds == nil {
		var ret []string
		return ret
	}
	return o.FieldIds
}

// GetFieldIdsOk returns a tuple with the FieldIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetFieldIdsOk() ([]string, bool) {
	if o == nil || o.FieldIds == nil {
		return nil, false
	}
	return o.FieldIds, true
}

// HasFieldIds returns a boolean if a field has been set.
func (o *View) HasFieldIds() bool {
	if o != nil && o.FieldIds != nil {
		return true
	}

	return false
}

// SetFieldIds gets a reference to the given []string and assigns it to the FieldIds field.
func (o *View) SetFieldIds(v []string) {
	o.FieldIds = v
}

// GetFieldMetas returns the FieldMetas field value if set, zero value otherwise.
func (o *View) GetFieldMetas() map[string]ViewFieldMeta {
	if o == nil || o.FieldMetas == nil {
		var ret map[string]ViewFieldMeta
		return ret
	}
	return *o.FieldMetas
}

// GetFieldMetasOk returns a tuple with the FieldMetas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetFieldMetasOk() (*map[string]ViewFieldMeta, bool) {
	if o == nil || o.FieldMetas == nil {
		return nil, false
	}
	return o.FieldMetas, true
}

// HasFieldMetas returns a boolean if a field has been set.
func (o *View) HasFieldMetas() bool {
	if o != nil && o.FieldMetas != nil {
		return true
	}

	return false
}

// SetFieldMetas gets a reference to the given map[string]ViewFieldMeta and assigns it to the FieldMetas field.
func (o *View) SetFieldMetas(v map[string]ViewFieldMeta) {
	o.FieldMetas = &v
}

// GetFilterSpec returns the FilterSpec field value if set, zero value otherwise.
func (o *View) GetFilterSpec() FilterSpec {
	if o == nil || o.FilterSpec == nil {
		var ret FilterSpec
		return ret
	}
	return *o.FilterSpec
}

// GetFilterSpecOk returns a tuple with the FilterSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetFilterSpecOk() (*FilterSpec, bool) {
	if o == nil || o.FilterSpec == nil {
		return nil, false
	}
	return o.FilterSpec, true
}

// HasFilterSpec returns a boolean if a field has been set.
func (o *View) HasFilterSpec() bool {
	if o != nil && o.FilterSpec != nil {
		return true
	}

	return false
}

// SetFilterSpec gets a reference to the given FilterSpec and assigns it to the FilterSpec field.
func (o *View) SetFilterSpec(v FilterSpec) {
	o.FilterSpec = &v
}

// GetAutoSort returns the AutoSort field value if set, zero value otherwise.
func (o *View) GetAutoSort() bool {
	if o == nil || o.AutoSort == nil {
		var ret bool
		return ret
	}
	return *o.AutoSort
}

// GetAutoSortOk returns a tuple with the AutoSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetAutoSortOk() (*bool, bool) {
	if o == nil || o.AutoSort == nil {
		return nil, false
	}
	return o.AutoSort, true
}

// HasAutoSort returns a boolean if a field has been set.
func (o *View) HasAutoSort() bool {
	if o != nil && o.AutoSort != nil {
		return true
	}

	return false
}

// SetAutoSort gets a reference to the given bool and assigns it to the AutoSort field.
func (o *View) SetAutoSort(v bool) {
	o.AutoSort = &v
}

// GetSortSpec returns the SortSpec field value if set, zero value otherwise.
func (o *View) GetSortSpec() SortSpec {
	if o == nil || o.SortSpec == nil {
		var ret SortSpec
		return ret
	}
	return *o.SortSpec
}

// GetSortSpecOk returns a tuple with the SortSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetSortSpecOk() (*SortSpec, bool) {
	if o == nil || o.SortSpec == nil {
		return nil, false
	}
	return o.SortSpec, true
}

// HasSortSpec returns a boolean if a field has been set.
func (o *View) HasSortSpec() bool {
	if o != nil && o.SortSpec != nil {
		return true
	}

	return false
}

// SetSortSpec gets a reference to the given SortSpec and assigns it to the SortSpec field.
func (o *View) SetSortSpec(v SortSpec) {
	o.SortSpec = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *View) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *View) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *View) SetOwner(v string) {
	o.Owner = &v
}

// GetPublicLevel returns the PublicLevel field value if set, zero value otherwise.
func (o *View) GetPublicLevel() int32 {
	if o == nil || o.PublicLevel == nil {
		var ret int32
		return ret
	}
	return *o.PublicLevel
}

// GetPublicLevelOk returns a tuple with the PublicLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetPublicLevelOk() (*int32, bool) {
	if o == nil || o.PublicLevel == nil {
		return nil, false
	}
	return o.PublicLevel, true
}

// HasPublicLevel returns a boolean if a field has been set.
func (o *View) HasPublicLevel() bool {
	if o != nil && o.PublicLevel != nil {
		return true
	}

	return false
}

// SetPublicLevel gets a reference to the given int32 and assigns it to the PublicLevel field.
func (o *View) SetPublicLevel(v int32) {
	o.PublicLevel = &v
}

// GetFrozenFieldsCount returns the FrozenFieldsCount field value if set, zero value otherwise.
func (o *View) GetFrozenFieldsCount() int32 {
	if o == nil || o.FrozenFieldsCount == nil {
		var ret int32
		return ret
	}
	return *o.FrozenFieldsCount
}

// GetFrozenFieldsCountOk returns a tuple with the FrozenFieldsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetFrozenFieldsCountOk() (*int32, bool) {
	if o == nil || o.FrozenFieldsCount == nil {
		return nil, false
	}
	return o.FrozenFieldsCount, true
}

// HasFrozenFieldsCount returns a boolean if a field has been set.
func (o *View) HasFrozenFieldsCount() bool {
	if o != nil && o.FrozenFieldsCount != nil {
		return true
	}

	return false
}

// SetFrozenFieldsCount gets a reference to the given int32 and assigns it to the FrozenFieldsCount field.
func (o *View) SetFrozenFieldsCount(v int32) {
	o.FrozenFieldsCount = &v
}

// GetGroupSpec returns the GroupSpec field value if set, zero value otherwise.
func (o *View) GetGroupSpec() GroupSpec {
	if o == nil || o.GroupSpec == nil {
		var ret GroupSpec
		return ret
	}
	return *o.GroupSpec
}

// GetGroupSpecOk returns a tuple with the GroupSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetGroupSpecOk() (*GroupSpec, bool) {
	if o == nil || o.GroupSpec == nil {
		return nil, false
	}
	return o.GroupSpec, true
}

// HasGroupSpec returns a boolean if a field has been set.
func (o *View) HasGroupSpec() bool {
	if o != nil && o.GroupSpec != nil {
		return true
	}

	return false
}

// SetGroupSpec gets a reference to the given GroupSpec and assigns it to the GroupSpec field.
func (o *View) SetGroupSpec(v GroupSpec) {
	o.GroupSpec = &v
}

// GetRowHeightLevel returns the RowHeightLevel field value if set, zero value otherwise.
func (o *View) GetRowHeightLevel() int32 {
	if o == nil || o.RowHeightLevel == nil {
		var ret int32
		return ret
	}
	return *o.RowHeightLevel
}

// GetRowHeightLevelOk returns a tuple with the RowHeightLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetRowHeightLevelOk() (*int32, bool) {
	if o == nil || o.RowHeightLevel == nil {
		return nil, false
	}
	return o.RowHeightLevel, true
}

// HasRowHeightLevel returns a boolean if a field has been set.
func (o *View) HasRowHeightLevel() bool {
	if o != nil && o.RowHeightLevel != nil {
		return true
	}

	return false
}

// SetRowHeightLevel gets a reference to the given int32 and assigns it to the RowHeightLevel field.
func (o *View) SetRowHeightLevel(v int32) {
	o.RowHeightLevel = &v
}

// GetIsFieldStatEnabled returns the IsFieldStatEnabled field value if set, zero value otherwise.
func (o *View) GetIsFieldStatEnabled() bool {
	if o == nil || o.IsFieldStatEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsFieldStatEnabled
}

// GetIsFieldStatEnabledOk returns a tuple with the IsFieldStatEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetIsFieldStatEnabledOk() (*bool, bool) {
	if o == nil || o.IsFieldStatEnabled == nil {
		return nil, false
	}
	return o.IsFieldStatEnabled, true
}

// HasIsFieldStatEnabled returns a boolean if a field has been set.
func (o *View) HasIsFieldStatEnabled() bool {
	if o != nil && o.IsFieldStatEnabled != nil {
		return true
	}

	return false
}

// SetIsFieldStatEnabled gets a reference to the given bool and assigns it to the IsFieldStatEnabled field.
func (o *View) SetIsFieldStatEnabled(v bool) {
	o.IsFieldStatEnabled = &v
}

// GetCardConfig returns the CardConfig field value if set, zero value otherwise.
func (o *View) GetCardConfig() CardConfig {
	if o == nil || o.CardConfig == nil {
		var ret CardConfig
		return ret
	}
	return *o.CardConfig
}

// GetCardConfigOk returns a tuple with the CardConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetCardConfigOk() (*CardConfig, bool) {
	if o == nil || o.CardConfig == nil {
		return nil, false
	}
	return o.CardConfig, true
}

// HasCardConfig returns a boolean if a field has been set.
func (o *View) HasCardConfig() bool {
	if o != nil && o.CardConfig != nil {
		return true
	}

	return false
}

// SetCardConfig gets a reference to the given CardConfig and assigns it to the CardConfig field.
func (o *View) SetCardConfig(v CardConfig) {
	o.CardConfig = &v
}

func (o View) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.RecordIds != nil {
		toSerialize["recordIds"] = o.RecordIds
	}
	if o.FieldIds != nil {
		toSerialize["fieldIds"] = o.FieldIds
	}
	if o.FieldMetas != nil {
		toSerialize["fieldMetas"] = o.FieldMetas
	}
	if o.FilterSpec != nil {
		toSerialize["filterSpec"] = o.FilterSpec
	}
	if o.AutoSort != nil {
		toSerialize["autoSort"] = o.AutoSort
	}
	if o.SortSpec != nil {
		toSerialize["sortSpec"] = o.SortSpec
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.PublicLevel != nil {
		toSerialize["publicLevel"] = o.PublicLevel
	}
	if o.FrozenFieldsCount != nil {
		toSerialize["frozenFieldsCount"] = o.FrozenFieldsCount
	}
	if o.GroupSpec != nil {
		toSerialize["groupSpec"] = o.GroupSpec
	}
	if o.RowHeightLevel != nil {
		toSerialize["rowHeightLevel"] = o.RowHeightLevel
	}
	if o.IsFieldStatEnabled != nil {
		toSerialize["isFieldStatEnabled"] = o.IsFieldStatEnabled
	}
	if o.CardConfig != nil {
		toSerialize["cardConfig"] = o.CardConfig
	}
	return json.Marshal(toSerialize)
}

type NullableView struct {
	value *View
	isSet bool
}

func (v NullableView) Get() *View {
	return v.value
}

func (v *NullableView) Set(val *View) {
	v.value = val
	v.isSet = true
}

func (v NullableView) IsSet() bool {
	return v.isSet
}

func (v *NullableView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableView(val *View) *NullableView {
	return &NullableView{value: val, isSet: true}
}

func (v NullableView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


