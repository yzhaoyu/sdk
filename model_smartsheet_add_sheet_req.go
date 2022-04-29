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

// SmartsheetAddSheetReq struct for SmartsheetAddSheetReq
type SmartsheetAddSheetReq struct {
	Properties *SheetProperties `json:"properties,omitempty"`
}

// NewSmartsheetAddSheetReq instantiates a new SmartsheetAddSheetReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetAddSheetReq() *SmartsheetAddSheetReq {
	this := SmartsheetAddSheetReq{}
	return &this
}

// NewSmartsheetAddSheetReqWithDefaults instantiates a new SmartsheetAddSheetReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetAddSheetReqWithDefaults() *SmartsheetAddSheetReq {
	this := SmartsheetAddSheetReq{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *SmartsheetAddSheetReq) GetProperties() SheetProperties {
	if o == nil || o.Properties == nil {
		var ret SheetProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetAddSheetReq) GetPropertiesOk() (*SheetProperties, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *SmartsheetAddSheetReq) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given SheetProperties and assigns it to the Properties field.
func (o *SmartsheetAddSheetReq) SetProperties(v SheetProperties) {
	o.Properties = &v
}

func (o SmartsheetAddSheetReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetAddSheetReq struct {
	value *SmartsheetAddSheetReq
	isSet bool
}

func (v NullableSmartsheetAddSheetReq) Get() *SmartsheetAddSheetReq {
	return v.value
}

func (v *NullableSmartsheetAddSheetReq) Set(val *SmartsheetAddSheetReq) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetAddSheetReq) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetAddSheetReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetAddSheetReq(val *SmartsheetAddSheetReq) *NullableSmartsheetAddSheetReq {
	return &NullableSmartsheetAddSheetReq{value: val, isSet: true}
}

func (v NullableSmartsheetAddSheetReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetAddSheetReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


