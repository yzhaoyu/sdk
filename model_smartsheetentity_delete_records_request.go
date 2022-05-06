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

// SmartsheetentityDeleteRecordsRequest 删除指定表格和子表中的一到多条记录请求
type SmartsheetentityDeleteRecordsRequest struct {
	RecordIDs []string `json:"recordIDs,omitempty"`
}

// NewSmartsheetentityDeleteRecordsRequest instantiates a new SmartsheetentityDeleteRecordsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetentityDeleteRecordsRequest() *SmartsheetentityDeleteRecordsRequest {
	this := SmartsheetentityDeleteRecordsRequest{}
	return &this
}

// NewSmartsheetentityDeleteRecordsRequestWithDefaults instantiates a new SmartsheetentityDeleteRecordsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetentityDeleteRecordsRequestWithDefaults() *SmartsheetentityDeleteRecordsRequest {
	this := SmartsheetentityDeleteRecordsRequest{}
	return &this
}

// GetRecordIDs returns the RecordIDs field value if set, zero value otherwise.
func (o *SmartsheetentityDeleteRecordsRequest) GetRecordIDs() []string {
	if o == nil || o.RecordIDs == nil {
		var ret []string
		return ret
	}
	return o.RecordIDs
}

// GetRecordIDsOk returns a tuple with the RecordIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetentityDeleteRecordsRequest) GetRecordIDsOk() ([]string, bool) {
	if o == nil || o.RecordIDs == nil {
		return nil, false
	}
	return o.RecordIDs, true
}

// HasRecordIDs returns a boolean if a field has been set.
func (o *SmartsheetentityDeleteRecordsRequest) HasRecordIDs() bool {
	if o != nil && o.RecordIDs != nil {
		return true
	}

	return false
}

// SetRecordIDs gets a reference to the given []string and assigns it to the RecordIDs field.
func (o *SmartsheetentityDeleteRecordsRequest) SetRecordIDs(v []string) {
	o.RecordIDs = v
}

func (o SmartsheetentityDeleteRecordsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecordIDs != nil {
		toSerialize["recordIDs"] = o.RecordIDs
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetentityDeleteRecordsRequest struct {
	value *SmartsheetentityDeleteRecordsRequest
	isSet bool
}

func (v NullableSmartsheetentityDeleteRecordsRequest) Get() *SmartsheetentityDeleteRecordsRequest {
	return v.value
}

func (v *NullableSmartsheetentityDeleteRecordsRequest) Set(val *SmartsheetentityDeleteRecordsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetentityDeleteRecordsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetentityDeleteRecordsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetentityDeleteRecordsRequest(val *SmartsheetentityDeleteRecordsRequest) *NullableSmartsheetentityDeleteRecordsRequest {
	return &NullableSmartsheetentityDeleteRecordsRequest{value: val, isSet: true}
}

func (v NullableSmartsheetentityDeleteRecordsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetentityDeleteRecordsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

