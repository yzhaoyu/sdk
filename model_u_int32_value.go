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

// UInt32Value Wrapper message for `uint32`. The JSON representation for `UInt32Value` is JSON number.
type UInt32Value struct {
	// The uint32 value.
	Value *int32 `json:"value,omitempty"`
}

// NewUInt32Value instantiates a new UInt32Value object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUInt32Value() *UInt32Value {
	this := UInt32Value{}
	return &this
}

// NewUInt32ValueWithDefaults instantiates a new UInt32Value object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUInt32ValueWithDefaults() *UInt32Value {
	this := UInt32Value{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UInt32Value) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UInt32Value) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UInt32Value) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *UInt32Value) SetValue(v int32) {
	o.Value = &v
}

func (o UInt32Value) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUInt32Value struct {
	value *UInt32Value
	isSet bool
}

func (v NullableUInt32Value) Get() *UInt32Value {
	return v.value
}

func (v *NullableUInt32Value) Set(val *UInt32Value) {
	v.value = val
	v.isSet = true
}

func (v NullableUInt32Value) IsSet() bool {
	return v.isSet
}

func (v *NullableUInt32Value) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUInt32Value(val *UInt32Value) *NullableUInt32Value {
	return &NullableUInt32Value{value: val, isSet: true}
}

func (v NullableUInt32Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUInt32Value) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


