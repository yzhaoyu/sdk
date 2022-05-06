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

// SmartsheetAddViewRequest 新建视图请求，创建一个视图
type SmartsheetAddViewRequest struct {
	ViewTitle *string `json:"viewTitle,omitempty"`
	ViewType *string `json:"viewType,omitempty"`
}

// NewSmartsheetAddViewRequest instantiates a new SmartsheetAddViewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetAddViewRequest() *SmartsheetAddViewRequest {
	this := SmartsheetAddViewRequest{}
	return &this
}

// NewSmartsheetAddViewRequestWithDefaults instantiates a new SmartsheetAddViewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetAddViewRequestWithDefaults() *SmartsheetAddViewRequest {
	this := SmartsheetAddViewRequest{}
	return &this
}

// GetViewTitle returns the ViewTitle field value if set, zero value otherwise.
func (o *SmartsheetAddViewRequest) GetViewTitle() string {
	if o == nil || o.ViewTitle == nil {
		var ret string
		return ret
	}
	return *o.ViewTitle
}

// GetViewTitleOk returns a tuple with the ViewTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetAddViewRequest) GetViewTitleOk() (*string, bool) {
	if o == nil || o.ViewTitle == nil {
		return nil, false
	}
	return o.ViewTitle, true
}

// HasViewTitle returns a boolean if a field has been set.
func (o *SmartsheetAddViewRequest) HasViewTitle() bool {
	if o != nil && o.ViewTitle != nil {
		return true
	}

	return false
}

// SetViewTitle gets a reference to the given string and assigns it to the ViewTitle field.
func (o *SmartsheetAddViewRequest) SetViewTitle(v string) {
	o.ViewTitle = &v
}

// GetViewType returns the ViewType field value if set, zero value otherwise.
func (o *SmartsheetAddViewRequest) GetViewType() string {
	if o == nil || o.ViewType == nil {
		var ret string
		return ret
	}
	return *o.ViewType
}

// GetViewTypeOk returns a tuple with the ViewType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetAddViewRequest) GetViewTypeOk() (*string, bool) {
	if o == nil || o.ViewType == nil {
		return nil, false
	}
	return o.ViewType, true
}

// HasViewType returns a boolean if a field has been set.
func (o *SmartsheetAddViewRequest) HasViewType() bool {
	if o != nil && o.ViewType != nil {
		return true
	}

	return false
}

// SetViewType gets a reference to the given string and assigns it to the ViewType field.
func (o *SmartsheetAddViewRequest) SetViewType(v string) {
	o.ViewType = &v
}

func (o SmartsheetAddViewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ViewTitle != nil {
		toSerialize["viewTitle"] = o.ViewTitle
	}
	if o.ViewType != nil {
		toSerialize["viewType"] = o.ViewType
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetAddViewRequest struct {
	value *SmartsheetAddViewRequest
	isSet bool
}

func (v NullableSmartsheetAddViewRequest) Get() *SmartsheetAddViewRequest {
	return v.value
}

func (v *NullableSmartsheetAddViewRequest) Set(val *SmartsheetAddViewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetAddViewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetAddViewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetAddViewRequest(val *SmartsheetAddViewRequest) *NullableSmartsheetAddViewRequest {
	return &NullableSmartsheetAddViewRequest{value: val, isSet: true}
}

func (v NullableSmartsheetAddViewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetAddViewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

