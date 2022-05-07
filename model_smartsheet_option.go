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

// SmartsheetOption struct for SmartsheetOption
type SmartsheetOption struct {
	Id *string `json:"id,omitempty"`
	Text *string `json:"text,omitempty"`
	Style *int32 `json:"style,omitempty"`
}

// NewSmartsheetOption instantiates a new SmartsheetOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetOption() *SmartsheetOption {
	this := SmartsheetOption{}
	return &this
}

// NewSmartsheetOptionWithDefaults instantiates a new SmartsheetOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetOptionWithDefaults() *SmartsheetOption {
	this := SmartsheetOption{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SmartsheetOption) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetOption) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SmartsheetOption) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SmartsheetOption) SetId(v string) {
	o.Id = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *SmartsheetOption) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetOption) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *SmartsheetOption) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *SmartsheetOption) SetText(v string) {
	o.Text = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *SmartsheetOption) GetStyle() int32 {
	if o == nil || o.Style == nil {
		var ret int32
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetOption) GetStyleOk() (*int32, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *SmartsheetOption) HasStyle() bool {
	if o != nil && o.Style != nil {
		return true
	}

	return false
}

// SetStyle gets a reference to the given int32 and assigns it to the Style field.
func (o *SmartsheetOption) SetStyle(v int32) {
	o.Style = &v
}

func (o SmartsheetOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetOption struct {
	value *SmartsheetOption
	isSet bool
}

func (v NullableSmartsheetOption) Get() *SmartsheetOption {
	return v.value
}

func (v *NullableSmartsheetOption) Set(val *SmartsheetOption) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetOption) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetOption(val *SmartsheetOption) *NullableSmartsheetOption {
	return &NullableSmartsheetOption{value: val, isSet: true}
}

func (v NullableSmartsheetOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


