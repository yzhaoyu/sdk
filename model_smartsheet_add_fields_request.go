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

// SmartsheetAddFieldsRequest 新增字段请求
type SmartsheetAddFieldsRequest struct {
	Fields []SmartsheetAddField `json:"fields,omitempty"`
}

// NewSmartsheetAddFieldsRequest instantiates a new SmartsheetAddFieldsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetAddFieldsRequest() *SmartsheetAddFieldsRequest {
	this := SmartsheetAddFieldsRequest{}
	return &this
}

// NewSmartsheetAddFieldsRequestWithDefaults instantiates a new SmartsheetAddFieldsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetAddFieldsRequestWithDefaults() *SmartsheetAddFieldsRequest {
	this := SmartsheetAddFieldsRequest{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *SmartsheetAddFieldsRequest) GetFields() []SmartsheetAddField {
	if o == nil || o.Fields == nil {
		var ret []SmartsheetAddField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetAddFieldsRequest) GetFieldsOk() ([]SmartsheetAddField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *SmartsheetAddFieldsRequest) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []SmartsheetAddField and assigns it to the Fields field.
func (o *SmartsheetAddFieldsRequest) SetFields(v []SmartsheetAddField) {
	o.Fields = v
}

func (o SmartsheetAddFieldsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetAddFieldsRequest struct {
	value *SmartsheetAddFieldsRequest
	isSet bool
}

func (v NullableSmartsheetAddFieldsRequest) Get() *SmartsheetAddFieldsRequest {
	return v.value
}

func (v *NullableSmartsheetAddFieldsRequest) Set(val *SmartsheetAddFieldsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetAddFieldsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetAddFieldsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetAddFieldsRequest(val *SmartsheetAddFieldsRequest) *NullableSmartsheetAddFieldsRequest {
	return &NullableSmartsheetAddFieldsRequest{value: val, isSet: true}
}

func (v NullableSmartsheetAddFieldsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetAddFieldsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


