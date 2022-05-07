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

// ResourcesAPIResponse1 struct for ResourcesAPIResponse1
type ResourcesAPIResponse1 struct {
	Ret int32 `json:"ret"`
	Msg string `json:"msg"`
	Data *UploadImageRsp `json:"data,omitempty"`
}

// NewResourcesAPIResponse1 instantiates a new ResourcesAPIResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcesAPIResponse1(ret int32, msg string) *ResourcesAPIResponse1 {
	this := ResourcesAPIResponse1{}
	this.Ret = ret
	this.Msg = msg
	return &this
}

// NewResourcesAPIResponse1WithDefaults instantiates a new ResourcesAPIResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcesAPIResponse1WithDefaults() *ResourcesAPIResponse1 {
	this := ResourcesAPIResponse1{}
	return &this
}

// GetRet returns the Ret field value
func (o *ResourcesAPIResponse1) GetRet() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ret
}

// GetRetOk returns a tuple with the Ret field value
// and a boolean to check if the value has been set.
func (o *ResourcesAPIResponse1) GetRetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ret, true
}

// SetRet sets field value
func (o *ResourcesAPIResponse1) SetRet(v int32) {
	o.Ret = v
}

// GetMsg returns the Msg field value
func (o *ResourcesAPIResponse1) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *ResourcesAPIResponse1) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *ResourcesAPIResponse1) SetMsg(v string) {
	o.Msg = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResourcesAPIResponse1) GetData() UploadImageRsp {
	if o == nil || o.Data == nil {
		var ret UploadImageRsp
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcesAPIResponse1) GetDataOk() (*UploadImageRsp, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResourcesAPIResponse1) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given UploadImageRsp and assigns it to the Data field.
func (o *ResourcesAPIResponse1) SetData(v UploadImageRsp) {
	o.Data = &v
}

func (o ResourcesAPIResponse1) MarshalJSON() ([]byte, error) {
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

type NullableResourcesAPIResponse1 struct {
	value *ResourcesAPIResponse1
	isSet bool
}

func (v NullableResourcesAPIResponse1) Get() *ResourcesAPIResponse1 {
	return v.value
}

func (v *NullableResourcesAPIResponse1) Set(val *ResourcesAPIResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcesAPIResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcesAPIResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcesAPIResponse1(val *ResourcesAPIResponse1) *NullableResourcesAPIResponse1 {
	return &NullableResourcesAPIResponse1{value: val, isSet: true}
}

func (v NullableResourcesAPIResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcesAPIResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


