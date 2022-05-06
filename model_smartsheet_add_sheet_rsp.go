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

// SmartsheetAddSheetRsp struct for SmartsheetAddSheetRsp
type SmartsheetAddSheetRsp struct {
	Properties *SheetProperties `json:"properties,omitempty"`
}

// NewSmartsheetAddSheetRsp instantiates a new SmartsheetAddSheetRsp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetAddSheetRsp() *SmartsheetAddSheetRsp {
	this := SmartsheetAddSheetRsp{}
	return &this
}

// NewSmartsheetAddSheetRspWithDefaults instantiates a new SmartsheetAddSheetRsp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetAddSheetRspWithDefaults() *SmartsheetAddSheetRsp {
	this := SmartsheetAddSheetRsp{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *SmartsheetAddSheetRsp) GetProperties() SheetProperties {
	if o == nil || o.Properties == nil {
		var ret SheetProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetAddSheetRsp) GetPropertiesOk() (*SheetProperties, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *SmartsheetAddSheetRsp) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given SheetProperties and assigns it to the Properties field.
func (o *SmartsheetAddSheetRsp) SetProperties(v SheetProperties) {
	o.Properties = &v
}

func (o SmartsheetAddSheetRsp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetAddSheetRsp struct {
	value *SmartsheetAddSheetRsp
	isSet bool
}

func (v NullableSmartsheetAddSheetRsp) Get() *SmartsheetAddSheetRsp {
	return v.value
}

func (v *NullableSmartsheetAddSheetRsp) Set(val *SmartsheetAddSheetRsp) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetAddSheetRsp) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetAddSheetRsp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetAddSheetRsp(val *SmartsheetAddSheetRsp) *NullableSmartsheetAddSheetRsp {
	return &NullableSmartsheetAddSheetRsp{value: val, isSet: true}
}

func (v NullableSmartsheetAddSheetRsp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetAddSheetRsp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

