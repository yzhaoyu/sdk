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

// OpenResponse1 struct for OpenResponse1
type OpenResponse1 struct {
	Ret int32 `json:"ret"`
	Msg string `json:"msg"`
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewOpenResponse1 instantiates a new OpenResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenResponse1(ret int32, msg string) *OpenResponse1 {
	this := OpenResponse1{}
	this.Ret = ret
	this.Msg = msg
	return &this
}

// NewOpenResponse1WithDefaults instantiates a new OpenResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenResponse1WithDefaults() *OpenResponse1 {
	this := OpenResponse1{}
	return &this
}

// GetRet returns the Ret field value
func (o *OpenResponse1) GetRet() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ret
}

// GetRetOk returns a tuple with the Ret field value
// and a boolean to check if the value has been set.
func (o *OpenResponse1) GetRetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ret, true
}

// SetRet sets field value
func (o *OpenResponse1) SetRet(v int32) {
	o.Ret = v
}

// GetMsg returns the Msg field value
func (o *OpenResponse1) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *OpenResponse1) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *OpenResponse1) SetMsg(v string) {
	o.Msg = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OpenResponse1) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenResponse1) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OpenResponse1) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *OpenResponse1) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o OpenResponse1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ret"] = o.Ret
	}
	if true {
		toSerialize["msg"] = o.Msg
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOpenResponse1 struct {
	value *OpenResponse1
	isSet bool
}

func (v NullableOpenResponse1) Get() *OpenResponse1 {
	return v.value
}

func (v *NullableOpenResponse1) Set(val *OpenResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenResponse1(val *OpenResponse1) *NullableOpenResponse1 {
	return &NullableOpenResponse1{value: val, isSet: true}
}

func (v NullableOpenResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


