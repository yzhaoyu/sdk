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

// AddFieldsResponse 新增字段回包
type AddFieldsResponse struct {
	Fields []FieldResource `json:"fields,omitempty"`
}

// NewAddFieldsResponse instantiates a new AddFieldsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFieldsResponse() *AddFieldsResponse {
	this := AddFieldsResponse{}
	return &this
}

// NewAddFieldsResponseWithDefaults instantiates a new AddFieldsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFieldsResponseWithDefaults() *AddFieldsResponse {
	this := AddFieldsResponse{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *AddFieldsResponse) GetFields() []FieldResource {
	if o == nil || o.Fields == nil {
		var ret []FieldResource
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFieldsResponse) GetFieldsOk() ([]FieldResource, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *AddFieldsResponse) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []FieldResource and assigns it to the Fields field.
func (o *AddFieldsResponse) SetFields(v []FieldResource) {
	o.Fields = v
}

func (o AddFieldsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableAddFieldsResponse struct {
	value *AddFieldsResponse
	isSet bool
}

func (v NullableAddFieldsResponse) Get() *AddFieldsResponse {
	return v.value
}

func (v *NullableAddFieldsResponse) Set(val *AddFieldsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFieldsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFieldsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFieldsResponse(val *AddFieldsResponse) *NullableAddFieldsResponse {
	return &NullableAddFieldsResponse{value: val, isSet: true}
}

func (v NullableAddFieldsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFieldsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


