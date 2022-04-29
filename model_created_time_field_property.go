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

// CreatedTimeFieldProperty struct for CreatedTimeFieldProperty
type CreatedTimeFieldProperty struct {
	Format *string `json:"format,omitempty"`
}

// NewCreatedTimeFieldProperty instantiates a new CreatedTimeFieldProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedTimeFieldProperty() *CreatedTimeFieldProperty {
	this := CreatedTimeFieldProperty{}
	return &this
}

// NewCreatedTimeFieldPropertyWithDefaults instantiates a new CreatedTimeFieldProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedTimeFieldPropertyWithDefaults() *CreatedTimeFieldProperty {
	this := CreatedTimeFieldProperty{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CreatedTimeFieldProperty) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedTimeFieldProperty) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CreatedTimeFieldProperty) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CreatedTimeFieldProperty) SetFormat(v string) {
	o.Format = &v
}

func (o CreatedTimeFieldProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	return json.Marshal(toSerialize)
}

type NullableCreatedTimeFieldProperty struct {
	value *CreatedTimeFieldProperty
	isSet bool
}

func (v NullableCreatedTimeFieldProperty) Get() *CreatedTimeFieldProperty {
	return v.value
}

func (v *NullableCreatedTimeFieldProperty) Set(val *CreatedTimeFieldProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedTimeFieldProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedTimeFieldProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedTimeFieldProperty(val *CreatedTimeFieldProperty) *NullableCreatedTimeFieldProperty {
	return &NullableCreatedTimeFieldProperty{value: val, isSet: true}
}

func (v NullableCreatedTimeFieldProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedTimeFieldProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


