/*
Tencent Docs SDK

Tencent docs sdk contains DocAPI, SmartsheetAPI, SheetAPI, DriveAPI and OAuthAPI

API version: 0.0.1
Contact: tencentdocs@tencent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FilterDateTimeValue struct for FilterDateTimeValue
type FilterDateTimeValue struct {
	Var1 *int32 `json:"1,omitempty"`
	Var2 *ListValue `json:"2,omitempty"`
}

// NewFilterDateTimeValue instantiates a new FilterDateTimeValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterDateTimeValue() *FilterDateTimeValue {
	this := FilterDateTimeValue{}
	return &this
}

// NewFilterDateTimeValueWithDefaults instantiates a new FilterDateTimeValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterDateTimeValueWithDefaults() *FilterDateTimeValue {
	this := FilterDateTimeValue{}
	return &this
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *FilterDateTimeValue) GetVar1() int32 {
	if o == nil || o.Var1 == nil {
		var ret int32
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterDateTimeValue) GetVar1Ok() (*int32, bool) {
	if o == nil || o.Var1 == nil {
		return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *FilterDateTimeValue) HasVar1() bool {
	if o != nil && o.Var1 != nil {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given int32 and assigns it to the Var1 field.
func (o *FilterDateTimeValue) SetVar1(v int32) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *FilterDateTimeValue) GetVar2() ListValue {
	if o == nil || o.Var2 == nil {
		var ret ListValue
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterDateTimeValue) GetVar2Ok() (*ListValue, bool) {
	if o == nil || o.Var2 == nil {
		return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *FilterDateTimeValue) HasVar2() bool {
	if o != nil && o.Var2 != nil {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given ListValue and assigns it to the Var2 field.
func (o *FilterDateTimeValue) SetVar2(v ListValue) {
	o.Var2 = &v
}

func (o FilterDateTimeValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Var1 != nil {
		toSerialize["1"] = o.Var1
	}
	if o.Var2 != nil {
		toSerialize["2"] = o.Var2
	}
	return json.Marshal(toSerialize)
}

type NullableFilterDateTimeValue struct {
	value *FilterDateTimeValue
	isSet bool
}

func (v NullableFilterDateTimeValue) Get() *FilterDateTimeValue {
	return v.value
}

func (v *NullableFilterDateTimeValue) Set(val *FilterDateTimeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterDateTimeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterDateTimeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterDateTimeValue(val *FilterDateTimeValue) *NullableFilterDateTimeValue {
	return &NullableFilterDateTimeValue{value: val, isSet: true}
}

func (v NullableFilterDateTimeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterDateTimeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


