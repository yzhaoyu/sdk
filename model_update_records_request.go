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

// UpdateRecordsRequest 更新指定表格和子表中的一到多条记录请求
type UpdateRecordsRequest struct {
	Records []CommonRecord `json:"records,omitempty"`
}

// NewUpdateRecordsRequest instantiates a new UpdateRecordsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRecordsRequest() *UpdateRecordsRequest {
	this := UpdateRecordsRequest{}
	return &this
}

// NewUpdateRecordsRequestWithDefaults instantiates a new UpdateRecordsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRecordsRequestWithDefaults() *UpdateRecordsRequest {
	this := UpdateRecordsRequest{}
	return &this
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *UpdateRecordsRequest) GetRecords() []CommonRecord {
	if o == nil || o.Records == nil {
		var ret []CommonRecord
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRecordsRequest) GetRecordsOk() ([]CommonRecord, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *UpdateRecordsRequest) HasRecords() bool {
	if o != nil && o.Records != nil {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []CommonRecord and assigns it to the Records field.
func (o *UpdateRecordsRequest) SetRecords(v []CommonRecord) {
	o.Records = v
}

func (o UpdateRecordsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRecordsRequest struct {
	value *UpdateRecordsRequest
	isSet bool
}

func (v NullableUpdateRecordsRequest) Get() *UpdateRecordsRequest {
	return v.value
}

func (v *NullableUpdateRecordsRequest) Set(val *UpdateRecordsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRecordsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRecordsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRecordsRequest(val *UpdateRecordsRequest) *NullableUpdateRecordsRequest {
	return &NullableUpdateRecordsRequest{value: val, isSet: true}
}

func (v NullableUpdateRecordsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRecordsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


