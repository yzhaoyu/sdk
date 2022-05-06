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

// SmartsheetentityGetRecordsRequest 获取指定表格和子表下指定视图中的所有记录、指定返回记录里要包含的字段、获取指定筛选条件和指定数量的记录三个接口合一
type SmartsheetentityGetRecordsRequest struct {
	ViewID *string `json:"viewID,omitempty"`
	RecordIDs []string `json:"recordIDs,omitempty"`
	FieldTitles []string `json:"fieldTitles,omitempty"`
	Sort []SmartsheetentitySort `json:"sort,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
}

// NewSmartsheetentityGetRecordsRequest instantiates a new SmartsheetentityGetRecordsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetentityGetRecordsRequest() *SmartsheetentityGetRecordsRequest {
	this := SmartsheetentityGetRecordsRequest{}
	return &this
}

// NewSmartsheetentityGetRecordsRequestWithDefaults instantiates a new SmartsheetentityGetRecordsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetentityGetRecordsRequestWithDefaults() *SmartsheetentityGetRecordsRequest {
	this := SmartsheetentityGetRecordsRequest{}
	return &this
}

// GetViewID returns the ViewID field value if set, zero value otherwise.
func (o *SmartsheetentityGetRecordsRequest) GetViewID() string {
	if o == nil || o.ViewID == nil {
		var ret string
		return ret
	}
	return *o.ViewID
}

// GetViewIDOk returns a tuple with the ViewID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetentityGetRecordsRequest) GetViewIDOk() (*string, bool) {
	if o == nil || o.ViewID == nil {
		return nil, false
	}
	return o.ViewID, true
}

// HasViewID returns a boolean if a field has been set.
func (o *SmartsheetentityGetRecordsRequest) HasViewID() bool {
	if o != nil && o.ViewID != nil {
		return true
	}

	return false
}

// SetViewID gets a reference to the given string and assigns it to the ViewID field.
func (o *SmartsheetentityGetRecordsRequest) SetViewID(v string) {
	o.ViewID = &v
}

// GetRecordIDs returns the RecordIDs field value if set, zero value otherwise.
func (o *SmartsheetentityGetRecordsRequest) GetRecordIDs() []string {
	if o == nil || o.RecordIDs == nil {
		var ret []string
		return ret
	}
	return o.RecordIDs
}

// GetRecordIDsOk returns a tuple with the RecordIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetentityGetRecordsRequest) GetRecordIDsOk() ([]string, bool) {
	if o == nil || o.RecordIDs == nil {
		return nil, false
	}
	return o.RecordIDs, true
}

// HasRecordIDs returns a boolean if a field has been set.
func (o *SmartsheetentityGetRecordsRequest) HasRecordIDs() bool {
	if o != nil && o.RecordIDs != nil {
		return true
	}

	return false
}

// SetRecordIDs gets a reference to the given []string and assigns it to the RecordIDs field.
func (o *SmartsheetentityGetRecordsRequest) SetRecordIDs(v []string) {
	o.RecordIDs = v
}

// GetFieldTitles returns the FieldTitles field value if set, zero value otherwise.
func (o *SmartsheetentityGetRecordsRequest) GetFieldTitles() []string {
	if o == nil || o.FieldTitles == nil {
		var ret []string
		return ret
	}
	return o.FieldTitles
}

// GetFieldTitlesOk returns a tuple with the FieldTitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetentityGetRecordsRequest) GetFieldTitlesOk() ([]string, bool) {
	if o == nil || o.FieldTitles == nil {
		return nil, false
	}
	return o.FieldTitles, true
}

// HasFieldTitles returns a boolean if a field has been set.
func (o *SmartsheetentityGetRecordsRequest) HasFieldTitles() bool {
	if o != nil && o.FieldTitles != nil {
		return true
	}

	return false
}

// SetFieldTitles gets a reference to the given []string and assigns it to the FieldTitles field.
func (o *SmartsheetentityGetRecordsRequest) SetFieldTitles(v []string) {
	o.FieldTitles = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *SmartsheetentityGetRecordsRequest) GetSort() []SmartsheetentitySort {
	if o == nil || o.Sort == nil {
		var ret []SmartsheetentitySort
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetentityGetRecordsRequest) GetSortOk() ([]SmartsheetentitySort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *SmartsheetentityGetRecordsRequest) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []SmartsheetentitySort and assigns it to the Sort field.
func (o *SmartsheetentityGetRecordsRequest) SetSort(v []SmartsheetentitySort) {
	o.Sort = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *SmartsheetentityGetRecordsRequest) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetentityGetRecordsRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *SmartsheetentityGetRecordsRequest) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *SmartsheetentityGetRecordsRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SmartsheetentityGetRecordsRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetentityGetRecordsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SmartsheetentityGetRecordsRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *SmartsheetentityGetRecordsRequest) SetLimit(v int32) {
	o.Limit = &v
}

func (o SmartsheetentityGetRecordsRequest) MarshalJSON() ([]byte, error) {
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

type NullableSmartsheetentityGetRecordsRequest struct {
	value *SmartsheetentityGetRecordsRequest
	isSet bool
}

func (v NullableSmartsheetentityGetRecordsRequest) Get() *SmartsheetentityGetRecordsRequest {
	return v.value
}

func (v *NullableSmartsheetentityGetRecordsRequest) Set(val *SmartsheetentityGetRecordsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetentityGetRecordsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetentityGetRecordsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetentityGetRecordsRequest(val *SmartsheetentityGetRecordsRequest) *NullableSmartsheetentityGetRecordsRequest {
	return &NullableSmartsheetentityGetRecordsRequest{value: val, isSet: true}
}

func (v NullableSmartsheetentityGetRecordsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetentityGetRecordsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

