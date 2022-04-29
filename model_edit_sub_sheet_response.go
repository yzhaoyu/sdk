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

// EditSubSheetResponse EditSheetResponse 编辑子表响应，创建、删除 Smartsheet 子表
type EditSubSheetResponse struct {
	AddSheet *SmartsheetAddSheetRsp `json:"addSheet,omitempty"`
	DeleteSheet map[string]interface{} `json:"deleteSheet,omitempty"`
}

// NewEditSubSheetResponse instantiates a new EditSubSheetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditSubSheetResponse() *EditSubSheetResponse {
	this := EditSubSheetResponse{}
	return &this
}

// NewEditSubSheetResponseWithDefaults instantiates a new EditSubSheetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditSubSheetResponseWithDefaults() *EditSubSheetResponse {
	this := EditSubSheetResponse{}
	return &this
}

// GetAddSheet returns the AddSheet field value if set, zero value otherwise.
func (o *EditSubSheetResponse) GetAddSheet() SmartsheetAddSheetRsp {
	if o == nil || o.AddSheet == nil {
		var ret SmartsheetAddSheetRsp
		return ret
	}
	return *o.AddSheet
}

// GetAddSheetOk returns a tuple with the AddSheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSubSheetResponse) GetAddSheetOk() (*SmartsheetAddSheetRsp, bool) {
	if o == nil || o.AddSheet == nil {
		return nil, false
	}
	return o.AddSheet, true
}

// HasAddSheet returns a boolean if a field has been set.
func (o *EditSubSheetResponse) HasAddSheet() bool {
	if o != nil && o.AddSheet != nil {
		return true
	}

	return false
}

// SetAddSheet gets a reference to the given SmartsheetAddSheetRsp and assigns it to the AddSheet field.
func (o *EditSubSheetResponse) SetAddSheet(v SmartsheetAddSheetRsp) {
	o.AddSheet = &v
}

// GetDeleteSheet returns the DeleteSheet field value if set, zero value otherwise.
func (o *EditSubSheetResponse) GetDeleteSheet() map[string]interface{} {
	if o == nil || o.DeleteSheet == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DeleteSheet
}

// GetDeleteSheetOk returns a tuple with the DeleteSheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSubSheetResponse) GetDeleteSheetOk() (map[string]interface{}, bool) {
	if o == nil || o.DeleteSheet == nil {
		return nil, false
	}
	return o.DeleteSheet, true
}

// HasDeleteSheet returns a boolean if a field has been set.
func (o *EditSubSheetResponse) HasDeleteSheet() bool {
	if o != nil && o.DeleteSheet != nil {
		return true
	}

	return false
}

// SetDeleteSheet gets a reference to the given map[string]interface{} and assigns it to the DeleteSheet field.
func (o *EditSubSheetResponse) SetDeleteSheet(v map[string]interface{}) {
	o.DeleteSheet = v
}

func (o EditSubSheetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddSheet != nil {
		toSerialize["addSheet"] = o.AddSheet
	}
	if o.DeleteSheet != nil {
		toSerialize["deleteSheet"] = o.DeleteSheet
	}
	return json.Marshal(toSerialize)
}

type NullableEditSubSheetResponse struct {
	value *EditSubSheetResponse
	isSet bool
}

func (v NullableEditSubSheetResponse) Get() *EditSubSheetResponse {
	return v.value
}

func (v *NullableEditSubSheetResponse) Set(val *EditSubSheetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEditSubSheetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEditSubSheetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditSubSheetResponse(val *EditSubSheetResponse) *NullableEditSubSheetResponse {
	return &NullableEditSubSheetResponse{value: val, isSet: true}
}

func (v NullableEditSubSheetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditSubSheetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


