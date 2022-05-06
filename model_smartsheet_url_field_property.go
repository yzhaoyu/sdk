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

// SmartsheetUrlFieldProperty struct for SmartsheetUrlFieldProperty
type SmartsheetUrlFieldProperty struct {
	Type *string `json:"type,omitempty"`
}

// NewSmartsheetUrlFieldProperty instantiates a new SmartsheetUrlFieldProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetUrlFieldProperty() *SmartsheetUrlFieldProperty {
	this := SmartsheetUrlFieldProperty{}
	return &this
}

// NewSmartsheetUrlFieldPropertyWithDefaults instantiates a new SmartsheetUrlFieldProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetUrlFieldPropertyWithDefaults() *SmartsheetUrlFieldProperty {
	this := SmartsheetUrlFieldProperty{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SmartsheetUrlFieldProperty) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetUrlFieldProperty) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SmartsheetUrlFieldProperty) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SmartsheetUrlFieldProperty) SetType(v string) {
	o.Type = &v
}

func (o SmartsheetUrlFieldProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetUrlFieldProperty struct {
	value *SmartsheetUrlFieldProperty
	isSet bool
}

func (v NullableSmartsheetUrlFieldProperty) Get() *SmartsheetUrlFieldProperty {
	return v.value
}

func (v *NullableSmartsheetUrlFieldProperty) Set(val *SmartsheetUrlFieldProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetUrlFieldProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetUrlFieldProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetUrlFieldProperty(val *SmartsheetUrlFieldProperty) *NullableSmartsheetUrlFieldProperty {
	return &NullableSmartsheetUrlFieldProperty{value: val, isSet: true}
}

func (v NullableSmartsheetUrlFieldProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetUrlFieldProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

