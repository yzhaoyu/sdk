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

// UpdateFieldsResponse 更新字段回包
type UpdateFieldsResponse struct {
	Fields []FieldResource `json:"fields,omitempty"`
}

// NewUpdateFieldsResponse instantiates a new UpdateFieldsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFieldsResponse() *UpdateFieldsResponse {
	this := UpdateFieldsResponse{}
	return &this
}

// NewUpdateFieldsResponseWithDefaults instantiates a new UpdateFieldsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFieldsResponseWithDefaults() *UpdateFieldsResponse {
	this := UpdateFieldsResponse{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *UpdateFieldsResponse) GetFields() []FieldResource {
	if o == nil || o.Fields == nil {
		var ret []FieldResource
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFieldsResponse) GetFieldsOk() ([]FieldResource, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *UpdateFieldsResponse) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []FieldResource and assigns it to the Fields field.
func (o *UpdateFieldsResponse) SetFields(v []FieldResource) {
	o.Fields = v
}

func (o UpdateFieldsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFieldsResponse struct {
	value *UpdateFieldsResponse
	isSet bool
}

func (v NullableUpdateFieldsResponse) Get() *UpdateFieldsResponse {
	return v.value
}

func (v *NullableUpdateFieldsResponse) Set(val *UpdateFieldsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFieldsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFieldsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFieldsResponse(val *UpdateFieldsResponse) *NullableUpdateFieldsResponse {
	return &NullableUpdateFieldsResponse{value: val, isSet: true}
}

func (v NullableUpdateFieldsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFieldsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


