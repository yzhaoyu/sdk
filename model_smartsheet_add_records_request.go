/*
Tencent Docs SDK

Tencent docs sdk contains DocAPI, SmartsheetAPI, SheetAPI, DriveAPI and OAuthAPI

API version: 0.0.1
Contact: tencentdocs@tencent.com
*/

// Code generated by Tencent Docs (https://docs.qq.com/open/wiki); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SmartsheetAddRecordsRequest 在指定表格和子表中创建一到多条新记录请求
type SmartsheetAddRecordsRequest struct {
	Records []SmartsheetAddRecord `json:"records,omitempty"`
}

// NewSmartsheetAddRecordsRequest instantiates a new SmartsheetAddRecordsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetAddRecordsRequest() *SmartsheetAddRecordsRequest {
	this := SmartsheetAddRecordsRequest{}
	return &this
}

// NewSmartsheetAddRecordsRequestWithDefaults instantiates a new SmartsheetAddRecordsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetAddRecordsRequestWithDefaults() *SmartsheetAddRecordsRequest {
	this := SmartsheetAddRecordsRequest{}
	return &this
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *SmartsheetAddRecordsRequest) GetRecords() []SmartsheetAddRecord {
	if o == nil || o.Records == nil {
		var ret []SmartsheetAddRecord
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetAddRecordsRequest) GetRecordsOk() ([]SmartsheetAddRecord, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *SmartsheetAddRecordsRequest) HasRecords() bool {
	if o != nil && o.Records != nil {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []SmartsheetAddRecord and assigns it to the Records field.
func (o *SmartsheetAddRecordsRequest) SetRecords(v []SmartsheetAddRecord) {
	o.Records = v
}

func (o SmartsheetAddRecordsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetAddRecordsRequest struct {
	value *SmartsheetAddRecordsRequest
	isSet bool
}

func (v NullableSmartsheetAddRecordsRequest) Get() *SmartsheetAddRecordsRequest {
	return v.value
}

func (v *NullableSmartsheetAddRecordsRequest) Set(val *SmartsheetAddRecordsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetAddRecordsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetAddRecordsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetAddRecordsRequest(val *SmartsheetAddRecordsRequest) *NullableSmartsheetAddRecordsRequest {
	return &NullableSmartsheetAddRecordsRequest{value: val, isSet: true}
}

func (v NullableSmartsheetAddRecordsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetAddRecordsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


