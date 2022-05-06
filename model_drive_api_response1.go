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

// DriveAPIResponse1 struct for DriveAPIResponse1
type DriveAPIResponse1 struct {
	Ret int32 `json:"ret"`
	Msg string `json:"msg"`
	Data *FileInfo `json:"data,omitempty"`
}

// NewDriveAPIResponse1 instantiates a new DriveAPIResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDriveAPIResponse1(ret int32, msg string) *DriveAPIResponse1 {
	this := DriveAPIResponse1{}
	this.Ret = ret
	this.Msg = msg
	return &this
}

// NewDriveAPIResponse1WithDefaults instantiates a new DriveAPIResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriveAPIResponse1WithDefaults() *DriveAPIResponse1 {
	this := DriveAPIResponse1{}
	return &this
}

// GetRet returns the Ret field value
func (o *DriveAPIResponse1) GetRet() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ret
}

// GetRetOk returns a tuple with the Ret field value
// and a boolean to check if the value has been set.
func (o *DriveAPIResponse1) GetRetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ret, true
}

// SetRet sets field value
func (o *DriveAPIResponse1) SetRet(v int32) {
	o.Ret = v
}

// GetMsg returns the Msg field value
func (o *DriveAPIResponse1) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *DriveAPIResponse1) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *DriveAPIResponse1) SetMsg(v string) {
	o.Msg = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DriveAPIResponse1) GetData() FileInfo {
	if o == nil || o.Data == nil {
		var ret FileInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveAPIResponse1) GetDataOk() (*FileInfo, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DriveAPIResponse1) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given FileInfo and assigns it to the Data field.
func (o *DriveAPIResponse1) SetData(v FileInfo) {
	o.Data = &v
}

func (o DriveAPIResponse1) MarshalJSON() ([]byte, error) {
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

type NullableDriveAPIResponse1 struct {
	value *DriveAPIResponse1
	isSet bool
}

func (v NullableDriveAPIResponse1) Get() *DriveAPIResponse1 {
	return v.value
}

func (v *NullableDriveAPIResponse1) Set(val *DriveAPIResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableDriveAPIResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableDriveAPIResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDriveAPIResponse1(val *DriveAPIResponse1) *NullableDriveAPIResponse1 {
	return &NullableDriveAPIResponse1{value: val, isSet: true}
}

func (v NullableDriveAPIResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDriveAPIResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


