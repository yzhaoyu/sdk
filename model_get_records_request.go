/*
Tencent Docs SDK

Tencent docs sdk contains DocAPI, SmartsheetAPI, SheetAPI, DriveAPI and OAuthAPI

API version: 0.0.1
Contact: tencentdocs@tencent.com
*/

// Code generated by Tencent Docs (https://docs.qq.com); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetRecordsRequest 获取指定表格和子表下指定视图中的所有记录、指定返回记录里要包含的字段、获取指定筛选条件和指定数量的记录三个接口合一
type GetRecordsRequest struct {
	ViewID *string `json:"viewID,omitempty"`
	RecordIDs []string `json:"recordIDs,omitempty"`
	FieldTitles []string `json:"fieldTitles,omitempty"`
	Sort []SortInfo `json:"sort,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
}

// NewGetRecordsRequest instantiates a new GetRecordsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecordsRequest() *GetRecordsRequest {
	this := GetRecordsRequest{}
	return &this
}

// NewGetRecordsRequestWithDefaults instantiates a new GetRecordsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecordsRequestWithDefaults() *GetRecordsRequest {
	this := GetRecordsRequest{}
	return &this
}

// GetViewID returns the ViewID field value if set, zero value otherwise.
func (o *GetRecordsRequest) GetViewID() string {
	if o == nil || o.ViewID == nil {
		var ret string
		return ret
	}
	return *o.ViewID
}

// GetViewIDOk returns a tuple with the ViewID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsRequest) GetViewIDOk() (*string, bool) {
	if o == nil || o.ViewID == nil {
		return nil, false
	}
	return o.ViewID, true
}

// HasViewID returns a boolean if a field has been set.
func (o *GetRecordsRequest) HasViewID() bool {
	if o != nil && o.ViewID != nil {
		return true
	}

	return false
}

// SetViewID gets a reference to the given string and assigns it to the ViewID field.
func (o *GetRecordsRequest) SetViewID(v string) {
	o.ViewID = &v
}

// GetRecordIDs returns the RecordIDs field value if set, zero value otherwise.
func (o *GetRecordsRequest) GetRecordIDs() []string {
	if o == nil || o.RecordIDs == nil {
		var ret []string
		return ret
	}
	return o.RecordIDs
}

// GetRecordIDsOk returns a tuple with the RecordIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsRequest) GetRecordIDsOk() ([]string, bool) {
	if o == nil || o.RecordIDs == nil {
		return nil, false
	}
	return o.RecordIDs, true
}

// HasRecordIDs returns a boolean if a field has been set.
func (o *GetRecordsRequest) HasRecordIDs() bool {
	if o != nil && o.RecordIDs != nil {
		return true
	}

	return false
}

// SetRecordIDs gets a reference to the given []string and assigns it to the RecordIDs field.
func (o *GetRecordsRequest) SetRecordIDs(v []string) {
	o.RecordIDs = v
}

// GetFieldTitles returns the FieldTitles field value if set, zero value otherwise.
func (o *GetRecordsRequest) GetFieldTitles() []string {
	if o == nil || o.FieldTitles == nil {
		var ret []string
		return ret
	}
	return o.FieldTitles
}

// GetFieldTitlesOk returns a tuple with the FieldTitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsRequest) GetFieldTitlesOk() ([]string, bool) {
	if o == nil || o.FieldTitles == nil {
		return nil, false
	}
	return o.FieldTitles, true
}

// HasFieldTitles returns a boolean if a field has been set.
func (o *GetRecordsRequest) HasFieldTitles() bool {
	if o != nil && o.FieldTitles != nil {
		return true
	}

	return false
}

// SetFieldTitles gets a reference to the given []string and assigns it to the FieldTitles field.
func (o *GetRecordsRequest) SetFieldTitles(v []string) {
	o.FieldTitles = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *GetRecordsRequest) GetSort() []SortInfo {
	if o == nil || o.Sort == nil {
		var ret []SortInfo
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsRequest) GetSortOk() ([]SortInfo, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *GetRecordsRequest) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []SortInfo and assigns it to the Sort field.
func (o *GetRecordsRequest) SetSort(v []SortInfo) {
	o.Sort = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetRecordsRequest) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetRecordsRequest) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GetRecordsRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetRecordsRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetRecordsRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetRecordsRequest) SetLimit(v int32) {
	o.Limit = &v
}

func (o GetRecordsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ViewID != nil {
		toSerialize["viewID"] = o.ViewID
	}
	if o.RecordIDs != nil {
		toSerialize["recordIDs"] = o.RecordIDs
	}
	if o.FieldTitles != nil {
		toSerialize["fieldTitles"] = o.FieldTitles
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableGetRecordsRequest struct {
	value *GetRecordsRequest
	isSet bool
}

func (v NullableGetRecordsRequest) Get() *GetRecordsRequest {
	return v.value
}

func (v *NullableGetRecordsRequest) Set(val *GetRecordsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecordsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecordsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecordsRequest(val *GetRecordsRequest) *NullableGetRecordsRequest {
	return &NullableGetRecordsRequest{value: val, isSet: true}
}

func (v NullableGetRecordsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecordsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


