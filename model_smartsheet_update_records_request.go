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

// SmartsheetUpdateRecordsRequest 更新指定表格和子表中的一到多条记录请求
type SmartsheetUpdateRecordsRequest struct {
	Records []SmartsheetCommonRecord `json:"records,omitempty"`
}

// NewSmartsheetUpdateRecordsRequest instantiates a new SmartsheetUpdateRecordsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetUpdateRecordsRequest() *SmartsheetUpdateRecordsRequest {
	this := SmartsheetUpdateRecordsRequest{}
	return &this
}

// NewSmartsheetUpdateRecordsRequestWithDefaults instantiates a new SmartsheetUpdateRecordsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetUpdateRecordsRequestWithDefaults() *SmartsheetUpdateRecordsRequest {
	this := SmartsheetUpdateRecordsRequest{}
	return &this
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *SmartsheetUpdateRecordsRequest) GetRecords() []SmartsheetCommonRecord {
	if o == nil || o.Records == nil {
		var ret []SmartsheetCommonRecord
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetUpdateRecordsRequest) GetRecordsOk() ([]SmartsheetCommonRecord, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *SmartsheetUpdateRecordsRequest) HasRecords() bool {
	if o != nil && o.Records != nil {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []SmartsheetCommonRecord and assigns it to the Records field.
func (o *SmartsheetUpdateRecordsRequest) SetRecords(v []SmartsheetCommonRecord) {
	o.Records = v
}

func (o SmartsheetUpdateRecordsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetUpdateRecordsRequest struct {
	value *SmartsheetUpdateRecordsRequest
	isSet bool
}

func (v NullableSmartsheetUpdateRecordsRequest) Get() *SmartsheetUpdateRecordsRequest {
	return v.value
}

func (v *NullableSmartsheetUpdateRecordsRequest) Set(val *SmartsheetUpdateRecordsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetUpdateRecordsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetUpdateRecordsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetUpdateRecordsRequest(val *SmartsheetUpdateRecordsRequest) *NullableSmartsheetUpdateRecordsRequest {
	return &NullableSmartsheetUpdateRecordsRequest{value: val, isSet: true}
}

func (v NullableSmartsheetUpdateRecordsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetUpdateRecordsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

