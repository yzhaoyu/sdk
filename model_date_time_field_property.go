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

// DateTimeFieldProperty struct for DateTimeFieldProperty
type DateTimeFieldProperty struct {
	Format *string `json:"format,omitempty"`
	AutoFill *bool `json:"autoFill,omitempty"`
}

// NewDateTimeFieldProperty instantiates a new DateTimeFieldProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateTimeFieldProperty() *DateTimeFieldProperty {
	this := DateTimeFieldProperty{}
	return &this
}

// NewDateTimeFieldPropertyWithDefaults instantiates a new DateTimeFieldProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateTimeFieldPropertyWithDefaults() *DateTimeFieldProperty {
	this := DateTimeFieldProperty{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *DateTimeFieldProperty) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldProperty) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *DateTimeFieldProperty) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *DateTimeFieldProperty) SetFormat(v string) {
	o.Format = &v
}

// GetAutoFill returns the AutoFill field value if set, zero value otherwise.
func (o *DateTimeFieldProperty) GetAutoFill() bool {
	if o == nil || o.AutoFill == nil {
		var ret bool
		return ret
	}
	return *o.AutoFill
}

// GetAutoFillOk returns a tuple with the AutoFill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFieldProperty) GetAutoFillOk() (*bool, bool) {
	if o == nil || o.AutoFill == nil {
		return nil, false
	}
	return o.AutoFill, true
}

// HasAutoFill returns a boolean if a field has been set.
func (o *DateTimeFieldProperty) HasAutoFill() bool {
	if o != nil && o.AutoFill != nil {
		return true
	}

	return false
}

// SetAutoFill gets a reference to the given bool and assigns it to the AutoFill field.
func (o *DateTimeFieldProperty) SetAutoFill(v bool) {
	o.AutoFill = &v
}

func (o DateTimeFieldProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.AutoFill != nil {
		toSerialize["autoFill"] = o.AutoFill
	}
	return json.Marshal(toSerialize)
}

type NullableDateTimeFieldProperty struct {
	value *DateTimeFieldProperty
	isSet bool
}

func (v NullableDateTimeFieldProperty) Get() *DateTimeFieldProperty {
	return v.value
}

func (v *NullableDateTimeFieldProperty) Set(val *DateTimeFieldProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableDateTimeFieldProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableDateTimeFieldProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateTimeFieldProperty(val *DateTimeFieldProperty) *NullableDateTimeFieldProperty {
	return &NullableDateTimeFieldProperty{value: val, isSet: true}
}

func (v NullableDateTimeFieldProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateTimeFieldProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


