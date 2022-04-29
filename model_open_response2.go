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

// OpenResponse2 struct for OpenResponse2
type OpenResponse2 struct {
	Ret int32 `json:"ret"`
	Msg string `json:"msg"`
	Data *GetDocFullTextRsp `json:"data,omitempty"`
}

// NewOpenResponse2 instantiates a new OpenResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenResponse2(ret int32, msg string) *OpenResponse2 {
	this := OpenResponse2{}
	this.Ret = ret
	this.Msg = msg
	return &this
}

// NewOpenResponse2WithDefaults instantiates a new OpenResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenResponse2WithDefaults() *OpenResponse2 {
	this := OpenResponse2{}
	return &this
}

// GetRet returns the Ret field value
func (o *OpenResponse2) GetRet() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ret
}

// GetRetOk returns a tuple with the Ret field value
// and a boolean to check if the value has been set.
func (o *OpenResponse2) GetRetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ret, true
}

// SetRet sets field value
func (o *OpenResponse2) SetRet(v int32) {
	o.Ret = v
}

// GetMsg returns the Msg field value
func (o *OpenResponse2) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *OpenResponse2) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *OpenResponse2) SetMsg(v string) {
	o.Msg = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OpenResponse2) GetData() GetDocFullTextRsp {
	if o == nil || o.Data == nil {
		var ret GetDocFullTextRsp
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenResponse2) GetDataOk() (*GetDocFullTextRsp, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OpenResponse2) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetDocFullTextRsp and assigns it to the Data field.
func (o *OpenResponse2) SetData(v GetDocFullTextRsp) {
	o.Data = &v
}

func (o OpenResponse2) MarshalJSON() ([]byte, error) {
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

type NullableOpenResponse2 struct {
	value *OpenResponse2
	isSet bool
}

func (v NullableOpenResponse2) Get() *OpenResponse2 {
	return v.value
}

func (v *NullableOpenResponse2) Set(val *OpenResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenResponse2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenResponse2(val *OpenResponse2) *NullableOpenResponse2 {
	return &NullableOpenResponse2{value: val, isSet: true}
}

func (v NullableOpenResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


