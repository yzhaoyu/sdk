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

// SmartsheetAddRecord struct for SmartsheetAddRecord
type SmartsheetAddRecord struct {
	Values *map[string]interface{} `json:"values,omitempty"`
}

// NewSmartsheetAddRecord instantiates a new SmartsheetAddRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetAddRecord() *SmartsheetAddRecord {
	this := SmartsheetAddRecord{}
	return &this
}

// NewSmartsheetAddRecordWithDefaults instantiates a new SmartsheetAddRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetAddRecordWithDefaults() *SmartsheetAddRecord {
	this := SmartsheetAddRecord{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SmartsheetAddRecord) GetValues() map[string]interface{} {
	if o == nil || o.Values == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetAddRecord) GetValuesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SmartsheetAddRecord) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]interface{} and assigns it to the Values field.
func (o *SmartsheetAddRecord) SetValues(v map[string]interface{}) {
	o.Values = &v
}

func (o SmartsheetAddRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetAddRecord struct {
	value *SmartsheetAddRecord
	isSet bool
}

func (v NullableSmartsheetAddRecord) Get() *SmartsheetAddRecord {
	return v.value
}

func (v *NullableSmartsheetAddRecord) Set(val *SmartsheetAddRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetAddRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetAddRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetAddRecord(val *SmartsheetAddRecord) *NullableSmartsheetAddRecord {
	return &NullableSmartsheetAddRecord{value: val, isSet: true}
}

func (v NullableSmartsheetAddRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetAddRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


