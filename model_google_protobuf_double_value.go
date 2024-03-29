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

// GoogleProtobufDoubleValue Wrapper message for `double`. The JSON representation for `DoubleValue` is JSON number.
type GoogleProtobufDoubleValue struct {
	// The double value.
	Value *float64 `json:"value,omitempty"`
}

// NewGoogleProtobufDoubleValue instantiates a new GoogleProtobufDoubleValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleProtobufDoubleValue() *GoogleProtobufDoubleValue {
	this := GoogleProtobufDoubleValue{}
	return &this
}

// NewGoogleProtobufDoubleValueWithDefaults instantiates a new GoogleProtobufDoubleValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleProtobufDoubleValueWithDefaults() *GoogleProtobufDoubleValue {
	this := GoogleProtobufDoubleValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GoogleProtobufDoubleValue) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleProtobufDoubleValue) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GoogleProtobufDoubleValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *GoogleProtobufDoubleValue) SetValue(v float64) {
	o.Value = &v
}

func (o GoogleProtobufDoubleValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleProtobufDoubleValue struct {
	value *GoogleProtobufDoubleValue
	isSet bool
}

func (v NullableGoogleProtobufDoubleValue) Get() *GoogleProtobufDoubleValue {
	return v.value
}

func (v *NullableGoogleProtobufDoubleValue) Set(val *GoogleProtobufDoubleValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleProtobufDoubleValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleProtobufDoubleValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleProtobufDoubleValue(val *GoogleProtobufDoubleValue) *NullableGoogleProtobufDoubleValue {
	return &NullableGoogleProtobufDoubleValue{value: val, isSet: true}
}

func (v NullableGoogleProtobufDoubleValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleProtobufDoubleValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


