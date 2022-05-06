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

// SmartsheetentityDeleteViewsRequest 删除指定表格指定子表中的指定视图请求
type SmartsheetentityDeleteViewsRequest struct {
	ViewIDs []string `json:"viewIDs,omitempty"`
}

// NewSmartsheetentityDeleteViewsRequest instantiates a new SmartsheetentityDeleteViewsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetentityDeleteViewsRequest() *SmartsheetentityDeleteViewsRequest {
	this := SmartsheetentityDeleteViewsRequest{}
	return &this
}

// NewSmartsheetentityDeleteViewsRequestWithDefaults instantiates a new SmartsheetentityDeleteViewsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetentityDeleteViewsRequestWithDefaults() *SmartsheetentityDeleteViewsRequest {
	this := SmartsheetentityDeleteViewsRequest{}
	return &this
}

// GetViewIDs returns the ViewIDs field value if set, zero value otherwise.
func (o *SmartsheetentityDeleteViewsRequest) GetViewIDs() []string {
	if o == nil || o.ViewIDs == nil {
		var ret []string
		return ret
	}
	return o.ViewIDs
}

// GetViewIDsOk returns a tuple with the ViewIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetentityDeleteViewsRequest) GetViewIDsOk() ([]string, bool) {
	if o == nil || o.ViewIDs == nil {
		return nil, false
	}
	return o.ViewIDs, true
}

// HasViewIDs returns a boolean if a field has been set.
func (o *SmartsheetentityDeleteViewsRequest) HasViewIDs() bool {
	if o != nil && o.ViewIDs != nil {
		return true
	}

	return false
}

// SetViewIDs gets a reference to the given []string and assigns it to the ViewIDs field.
func (o *SmartsheetentityDeleteViewsRequest) SetViewIDs(v []string) {
	o.ViewIDs = v
}

func (o SmartsheetentityDeleteViewsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ViewIDs != nil {
		toSerialize["viewIDs"] = o.ViewIDs
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetentityDeleteViewsRequest struct {
	value *SmartsheetentityDeleteViewsRequest
	isSet bool
}

func (v NullableSmartsheetentityDeleteViewsRequest) Get() *SmartsheetentityDeleteViewsRequest {
	return v.value
}

func (v *NullableSmartsheetentityDeleteViewsRequest) Set(val *SmartsheetentityDeleteViewsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetentityDeleteViewsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetentityDeleteViewsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetentityDeleteViewsRequest(val *SmartsheetentityDeleteViewsRequest) *NullableSmartsheetentityDeleteViewsRequest {
	return &NullableSmartsheetentityDeleteViewsRequest{value: val, isSet: true}
}

func (v NullableSmartsheetentityDeleteViewsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetentityDeleteViewsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

