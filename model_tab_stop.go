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

// TabStop struct for TabStop
type TabStop struct {
	Offset *GoogleProtobufDoubleValue `json:"offset,omitempty"`
	Alignment *string `json:"alignment,omitempty"`
}

// NewTabStop instantiates a new TabStop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTabStop() *TabStop {
	this := TabStop{}
	return &this
}

// NewTabStopWithDefaults instantiates a new TabStop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTabStopWithDefaults() *TabStop {
	this := TabStop{}
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TabStop) GetOffset() GoogleProtobufDoubleValue {
	if o == nil || o.Offset == nil {
		var ret GoogleProtobufDoubleValue
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TabStop) GetOffsetOk() (*GoogleProtobufDoubleValue, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TabStop) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given GoogleProtobufDoubleValue and assigns it to the Offset field.
func (o *TabStop) SetOffset(v GoogleProtobufDoubleValue) {
	o.Offset = &v
}

// GetAlignment returns the Alignment field value if set, zero value otherwise.
func (o *TabStop) GetAlignment() string {
	if o == nil || o.Alignment == nil {
		var ret string
		return ret
	}
	return *o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TabStop) GetAlignmentOk() (*string, bool) {
	if o == nil || o.Alignment == nil {
		return nil, false
	}
	return o.Alignment, true
}

// HasAlignment returns a boolean if a field has been set.
func (o *TabStop) HasAlignment() bool {
	if o != nil && o.Alignment != nil {
		return true
	}

	return false
}

// SetAlignment gets a reference to the given string and assigns it to the Alignment field.
func (o *TabStop) SetAlignment(v string) {
	o.Alignment = &v
}

func (o TabStop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Alignment != nil {
		toSerialize["alignment"] = o.Alignment
	}
	return json.Marshal(toSerialize)
}

type NullableTabStop struct {
	value *TabStop
	isSet bool
}

func (v NullableTabStop) Get() *TabStop {
	return v.value
}

func (v *NullableTabStop) Set(val *TabStop) {
	v.value = val
	v.isSet = true
}

func (v NullableTabStop) IsSet() bool {
	return v.isSet
}

func (v *NullableTabStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTabStop(val *TabStop) *NullableTabStop {
	return &NullableTabStop{value: val, isSet: true}
}

func (v NullableTabStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTabStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


