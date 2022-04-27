/*
Tencent Docs SDK

Tencent docs sdk contains DocAPI, SmartsheetAPI, SheetAPI, DriveAPI and OAuthAPI

API version: 0.0.1
Contact: tencentdocs@tencent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UserResponse1 struct for UserResponse1
type UserResponse1 struct {
	Ret *int32 `json:"ret,omitempty"`
	Msg *string `json:"msg,omitempty"`
	Data *Response1 `json:"data,omitempty"`
}

// NewUserResponse1 instantiates a new UserResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse1() *UserResponse1 {
	this := UserResponse1{}
	return &this
}

// NewUserResponse1WithDefaults instantiates a new UserResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponse1WithDefaults() *UserResponse1 {
	this := UserResponse1{}
	return &this
}

// GetRet returns the Ret field value if set, zero value otherwise.
func (o *UserResponse1) GetRet() int32 {
	if o == nil || o.Ret == nil {
		var ret int32
		return ret
	}
	return *o.Ret
}

// GetRetOk returns a tuple with the Ret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse1) GetRetOk() (*int32, bool) {
	if o == nil || o.Ret == nil {
		return nil, false
	}
	return o.Ret, true
}

// HasRet returns a boolean if a field has been set.
func (o *UserResponse1) HasRet() bool {
	if o != nil && o.Ret != nil {
		return true
	}

	return false
}

// SetRet gets a reference to the given int32 and assigns it to the Ret field.
func (o *UserResponse1) SetRet(v int32) {
	o.Ret = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *UserResponse1) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse1) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *UserResponse1) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *UserResponse1) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserResponse1) GetData() Response1 {
	if o == nil || o.Data == nil {
		var ret Response1
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse1) GetDataOk() (*Response1, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserResponse1) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Response1 and assigns it to the Data field.
func (o *UserResponse1) SetData(v Response1) {
	o.Data = &v
}

func (o UserResponse1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ret != nil {
		toSerialize["ret"] = o.Ret
	}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUserResponse1 struct {
	value *UserResponse1
	isSet bool
}

func (v NullableUserResponse1) Get() *UserResponse1 {
	return v.value
}

func (v *NullableUserResponse1) Set(val *UserResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponse1(val *UserResponse1) *NullableUserResponse1 {
	return &NullableUserResponse1{value: val, isSet: true}
}

func (v NullableUserResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


