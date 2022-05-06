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

// SmartsheetentityAddRecordsResponse 在指定表格和子表中创建一到多条新记录回包
type SmartsheetentityAddRecordsResponse struct {
	Records []SmartsheetentityCommonRecord `json:"records,omitempty"`
}

// NewSmartsheetentityAddRecordsResponse instantiates a new SmartsheetentityAddRecordsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetentityAddRecordsResponse() *SmartsheetentityAddRecordsResponse {
	this := SmartsheetentityAddRecordsResponse{}
	return &this
}

// NewSmartsheetentityAddRecordsResponseWithDefaults instantiates a new SmartsheetentityAddRecordsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetentityAddRecordsResponseWithDefaults() *SmartsheetentityAddRecordsResponse {
	this := SmartsheetentityAddRecordsResponse{}
	return &this
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *SmartsheetentityAddRecordsResponse) GetRecords() []SmartsheetentityCommonRecord {
	if o == nil || o.Records == nil {
		var ret []SmartsheetentityCommonRecord
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetentityAddRecordsResponse) GetRecordsOk() ([]SmartsheetentityCommonRecord, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *SmartsheetentityAddRecordsResponse) HasRecords() bool {
	if o != nil && o.Records != nil {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []SmartsheetentityCommonRecord and assigns it to the Records field.
func (o *SmartsheetentityAddRecordsResponse) SetRecords(v []SmartsheetentityCommonRecord) {
	o.Records = v
}

func (o SmartsheetentityAddRecordsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetentityAddRecordsResponse struct {
	value *SmartsheetentityAddRecordsResponse
	isSet bool
}

func (v NullableSmartsheetentityAddRecordsResponse) Get() *SmartsheetentityAddRecordsResponse {
	return v.value
}

func (v *NullableSmartsheetentityAddRecordsResponse) Set(val *SmartsheetentityAddRecordsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetentityAddRecordsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetentityAddRecordsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetentityAddRecordsResponse(val *SmartsheetentityAddRecordsResponse) *NullableSmartsheetentityAddRecordsResponse {
	return &NullableSmartsheetentityAddRecordsResponse{value: val, isSet: true}
}

func (v NullableSmartsheetentityAddRecordsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetentityAddRecordsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

