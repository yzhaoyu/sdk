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

// SmartsheetDeleteViewsRequest 删除指定表格指定子表中的指定视图请求
type SmartsheetDeleteViewsRequest struct {
	ViewIDs []string `json:"viewIDs,omitempty"`
}

// NewSmartsheetDeleteViewsRequest instantiates a new SmartsheetDeleteViewsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetDeleteViewsRequest() *SmartsheetDeleteViewsRequest {
	this := SmartsheetDeleteViewsRequest{}
	return &this
}

// NewSmartsheetDeleteViewsRequestWithDefaults instantiates a new SmartsheetDeleteViewsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetDeleteViewsRequestWithDefaults() *SmartsheetDeleteViewsRequest {
	this := SmartsheetDeleteViewsRequest{}
	return &this
}

// GetViewIDs returns the ViewIDs field value if set, zero value otherwise.
func (o *SmartsheetDeleteViewsRequest) GetViewIDs() []string {
	if o == nil || o.ViewIDs == nil {
		var ret []string
		return ret
	}
	return o.ViewIDs
}

// GetViewIDsOk returns a tuple with the ViewIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetDeleteViewsRequest) GetViewIDsOk() ([]string, bool) {
	if o == nil || o.ViewIDs == nil {
		return nil, false
	}
	return o.ViewIDs, true
}

// HasViewIDs returns a boolean if a field has been set.
func (o *SmartsheetDeleteViewsRequest) HasViewIDs() bool {
	if o != nil && o.ViewIDs != nil {
		return true
	}

	return false
}

// SetViewIDs gets a reference to the given []string and assigns it to the ViewIDs field.
func (o *SmartsheetDeleteViewsRequest) SetViewIDs(v []string) {
	o.ViewIDs = v
}

func (o SmartsheetDeleteViewsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ViewIDs != nil {
		toSerialize["viewIDs"] = o.ViewIDs
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetDeleteViewsRequest struct {
	value *SmartsheetDeleteViewsRequest
	isSet bool
}

func (v NullableSmartsheetDeleteViewsRequest) Get() *SmartsheetDeleteViewsRequest {
	return v.value
}

func (v *NullableSmartsheetDeleteViewsRequest) Set(val *SmartsheetDeleteViewsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetDeleteViewsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetDeleteViewsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetDeleteViewsRequest(val *SmartsheetDeleteViewsRequest) *NullableSmartsheetDeleteViewsRequest {
	return &NullableSmartsheetDeleteViewsRequest{value: val, isSet: true}
}

func (v NullableSmartsheetDeleteViewsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetDeleteViewsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


