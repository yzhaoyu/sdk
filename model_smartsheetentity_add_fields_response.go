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

// SmartsheetentityAddFieldsResponse 新增字段回包
type SmartsheetentityAddFieldsResponse struct {
	Fields []SmartsheetentityField `json:"fields,omitempty"`
}

// NewSmartsheetentityAddFieldsResponse instantiates a new SmartsheetentityAddFieldsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetentityAddFieldsResponse() *SmartsheetentityAddFieldsResponse {
	this := SmartsheetentityAddFieldsResponse{}
	return &this
}

// NewSmartsheetentityAddFieldsResponseWithDefaults instantiates a new SmartsheetentityAddFieldsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetentityAddFieldsResponseWithDefaults() *SmartsheetentityAddFieldsResponse {
	this := SmartsheetentityAddFieldsResponse{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *SmartsheetentityAddFieldsResponse) GetFields() []SmartsheetentityField {
	if o == nil || o.Fields == nil {
		var ret []SmartsheetentityField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetentityAddFieldsResponse) GetFieldsOk() ([]SmartsheetentityField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *SmartsheetentityAddFieldsResponse) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []SmartsheetentityField and assigns it to the Fields field.
func (o *SmartsheetentityAddFieldsResponse) SetFields(v []SmartsheetentityField) {
	o.Fields = v
}

func (o SmartsheetentityAddFieldsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetentityAddFieldsResponse struct {
	value *SmartsheetentityAddFieldsResponse
	isSet bool
}

func (v NullableSmartsheetentityAddFieldsResponse) Get() *SmartsheetentityAddFieldsResponse {
	return v.value
}

func (v *NullableSmartsheetentityAddFieldsResponse) Set(val *SmartsheetentityAddFieldsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetentityAddFieldsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetentityAddFieldsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetentityAddFieldsResponse(val *SmartsheetentityAddFieldsResponse) *NullableSmartsheetentityAddFieldsResponse {
	return &NullableSmartsheetentityAddFieldsResponse{value: val, isSet: true}
}

func (v NullableSmartsheetentityAddFieldsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetentityAddFieldsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

