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

// SmartsheetCreatedTimeFieldProperty struct for SmartsheetCreatedTimeFieldProperty
type SmartsheetCreatedTimeFieldProperty struct {
	Format *string `json:"format,omitempty"`
}

// NewSmartsheetCreatedTimeFieldProperty instantiates a new SmartsheetCreatedTimeFieldProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetCreatedTimeFieldProperty() *SmartsheetCreatedTimeFieldProperty {
	this := SmartsheetCreatedTimeFieldProperty{}
	return &this
}

// NewSmartsheetCreatedTimeFieldPropertyWithDefaults instantiates a new SmartsheetCreatedTimeFieldProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetCreatedTimeFieldPropertyWithDefaults() *SmartsheetCreatedTimeFieldProperty {
	this := SmartsheetCreatedTimeFieldProperty{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *SmartsheetCreatedTimeFieldProperty) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetCreatedTimeFieldProperty) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *SmartsheetCreatedTimeFieldProperty) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *SmartsheetCreatedTimeFieldProperty) SetFormat(v string) {
	o.Format = &v
}

func (o SmartsheetCreatedTimeFieldProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetCreatedTimeFieldProperty struct {
	value *SmartsheetCreatedTimeFieldProperty
	isSet bool
}

func (v NullableSmartsheetCreatedTimeFieldProperty) Get() *SmartsheetCreatedTimeFieldProperty {
	return v.value
}

func (v *NullableSmartsheetCreatedTimeFieldProperty) Set(val *SmartsheetCreatedTimeFieldProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetCreatedTimeFieldProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetCreatedTimeFieldProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetCreatedTimeFieldProperty(val *SmartsheetCreatedTimeFieldProperty) *NullableSmartsheetCreatedTimeFieldProperty {
	return &NullableSmartsheetCreatedTimeFieldProperty{value: val, isSet: true}
}

func (v NullableSmartsheetCreatedTimeFieldProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetCreatedTimeFieldProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


