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

// ProgressFieldProperty struct for ProgressFieldProperty
type ProgressFieldProperty struct {
	DecimalPlaces *int32 `json:"decimalPlaces,omitempty"`
}

// NewProgressFieldProperty instantiates a new ProgressFieldProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgressFieldProperty() *ProgressFieldProperty {
	this := ProgressFieldProperty{}
	return &this
}

// NewProgressFieldPropertyWithDefaults instantiates a new ProgressFieldProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgressFieldPropertyWithDefaults() *ProgressFieldProperty {
	this := ProgressFieldProperty{}
	return &this
}

// GetDecimalPlaces returns the DecimalPlaces field value if set, zero value otherwise.
func (o *ProgressFieldProperty) GetDecimalPlaces() int32 {
	if o == nil || o.DecimalPlaces == nil {
		var ret int32
		return ret
	}
	return *o.DecimalPlaces
}

// GetDecimalPlacesOk returns a tuple with the DecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgressFieldProperty) GetDecimalPlacesOk() (*int32, bool) {
	if o == nil || o.DecimalPlaces == nil {
		return nil, false
	}
	return o.DecimalPlaces, true
}

// HasDecimalPlaces returns a boolean if a field has been set.
func (o *ProgressFieldProperty) HasDecimalPlaces() bool {
	if o != nil && o.DecimalPlaces != nil {
		return true
	}

	return false
}

// SetDecimalPlaces gets a reference to the given int32 and assigns it to the DecimalPlaces field.
func (o *ProgressFieldProperty) SetDecimalPlaces(v int32) {
	o.DecimalPlaces = &v
}

func (o ProgressFieldProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DecimalPlaces != nil {
		toSerialize["decimalPlaces"] = o.DecimalPlaces
	}
	return json.Marshal(toSerialize)
}

type NullableProgressFieldProperty struct {
	value *ProgressFieldProperty
	isSet bool
}

func (v NullableProgressFieldProperty) Get() *ProgressFieldProperty {
	return v.value
}

func (v *NullableProgressFieldProperty) Set(val *ProgressFieldProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableProgressFieldProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableProgressFieldProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgressFieldProperty(val *ProgressFieldProperty) *NullableProgressFieldProperty {
	return &NullableProgressFieldProperty{value: val, isSet: true}
}

func (v NullableProgressFieldProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgressFieldProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


