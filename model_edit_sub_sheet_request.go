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

// EditSubSheetRequest EditSheetRequest 编辑子表请求，创建、删除 Smartsheet 子表
type EditSubSheetRequest struct {
	FileID string `json:"fileID"`
	AddSheet *SmartsheetAddSheetReq `json:"addSheet,omitempty"`
	DeleteSheet *SmartsheetDeleteSheetReq `json:"deleteSheet,omitempty"`
}

// NewEditSubSheetRequest instantiates a new EditSubSheetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditSubSheetRequest(fileID string) *EditSubSheetRequest {
	this := EditSubSheetRequest{}
	this.FileID = fileID
	return &this
}

// NewEditSubSheetRequestWithDefaults instantiates a new EditSubSheetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditSubSheetRequestWithDefaults() *EditSubSheetRequest {
	this := EditSubSheetRequest{}
	return &this
}

// GetFileID returns the FileID field value
func (o *EditSubSheetRequest) GetFileID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileID
}

// GetFileIDOk returns a tuple with the FileID field value
// and a boolean to check if the value has been set.
func (o *EditSubSheetRequest) GetFileIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileID, true
}

// SetFileID sets field value
func (o *EditSubSheetRequest) SetFileID(v string) {
	o.FileID = v
}

// GetAddSheet returns the AddSheet field value if set, zero value otherwise.
func (o *EditSubSheetRequest) GetAddSheet() SmartsheetAddSheetReq {
	if o == nil || o.AddSheet == nil {
		var ret SmartsheetAddSheetReq
		return ret
	}
	return *o.AddSheet
}

// GetAddSheetOk returns a tuple with the AddSheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSubSheetRequest) GetAddSheetOk() (*SmartsheetAddSheetReq, bool) {
	if o == nil || o.AddSheet == nil {
		return nil, false
	}
	return o.AddSheet, true
}

// HasAddSheet returns a boolean if a field has been set.
func (o *EditSubSheetRequest) HasAddSheet() bool {
	if o != nil && o.AddSheet != nil {
		return true
	}

	return false
}

// SetAddSheet gets a reference to the given SmartsheetAddSheetReq and assigns it to the AddSheet field.
func (o *EditSubSheetRequest) SetAddSheet(v SmartsheetAddSheetReq) {
	o.AddSheet = &v
}

// GetDeleteSheet returns the DeleteSheet field value if set, zero value otherwise.
func (o *EditSubSheetRequest) GetDeleteSheet() SmartsheetDeleteSheetReq {
	if o == nil || o.DeleteSheet == nil {
		var ret SmartsheetDeleteSheetReq
		return ret
	}
	return *o.DeleteSheet
}

// GetDeleteSheetOk returns a tuple with the DeleteSheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSubSheetRequest) GetDeleteSheetOk() (*SmartsheetDeleteSheetReq, bool) {
	if o == nil || o.DeleteSheet == nil {
		return nil, false
	}
	return o.DeleteSheet, true
}

// HasDeleteSheet returns a boolean if a field has been set.
func (o *EditSubSheetRequest) HasDeleteSheet() bool {
	if o != nil && o.DeleteSheet != nil {
		return true
	}

	return false
}

// SetDeleteSheet gets a reference to the given SmartsheetDeleteSheetReq and assigns it to the DeleteSheet field.
func (o *EditSubSheetRequest) SetDeleteSheet(v SmartsheetDeleteSheetReq) {
	o.DeleteSheet = &v
}

func (o EditSubSheetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fileID"] = o.FileID
	}
	if o.AddSheet != nil {
		toSerialize["addSheet"] = o.AddSheet
	}
	if o.DeleteSheet != nil {
		toSerialize["deleteSheet"] = o.DeleteSheet
	}
	return json.Marshal(toSerialize)
}

type NullableEditSubSheetRequest struct {
	value *EditSubSheetRequest
	isSet bool
}

func (v NullableEditSubSheetRequest) Get() *EditSubSheetRequest {
	return v.value
}

func (v *NullableEditSubSheetRequest) Set(val *EditSubSheetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditSubSheetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditSubSheetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditSubSheetRequest(val *EditSubSheetRequest) *NullableEditSubSheetRequest {
	return &NullableEditSubSheetRequest{value: val, isSet: true}
}

func (v NullableEditSubSheetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditSubSheetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

