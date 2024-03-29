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

// SmartsheetGetFieldsRequest 获取指定表格和子表中的所有字段请求
type SmartsheetGetFieldsRequest struct {
	ViewID *string `json:"viewID,omitempty"`
	FieldIDs []string `json:"fieldIDs,omitempty"`
	FieldTitles []string `json:"fieldTitles,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
}

// NewSmartsheetGetFieldsRequest instantiates a new SmartsheetGetFieldsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetGetFieldsRequest() *SmartsheetGetFieldsRequest {
	this := SmartsheetGetFieldsRequest{}
	return &this
}

// NewSmartsheetGetFieldsRequestWithDefaults instantiates a new SmartsheetGetFieldsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetGetFieldsRequestWithDefaults() *SmartsheetGetFieldsRequest {
	this := SmartsheetGetFieldsRequest{}
	return &this
}

// GetViewID returns the ViewID field value if set, zero value otherwise.
func (o *SmartsheetGetFieldsRequest) GetViewID() string {
	if o == nil || o.ViewID == nil {
		var ret string
		return ret
	}
	return *o.ViewID
}

// GetViewIDOk returns a tuple with the ViewID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetGetFieldsRequest) GetViewIDOk() (*string, bool) {
	if o == nil || o.ViewID == nil {
		return nil, false
	}
	return o.ViewID, true
}

// HasViewID returns a boolean if a field has been set.
func (o *SmartsheetGetFieldsRequest) HasViewID() bool {
	if o != nil && o.ViewID != nil {
		return true
	}

	return false
}

// SetViewID gets a reference to the given string and assigns it to the ViewID field.
func (o *SmartsheetGetFieldsRequest) SetViewID(v string) {
	o.ViewID = &v
}

// GetFieldIDs returns the FieldIDs field value if set, zero value otherwise.
func (o *SmartsheetGetFieldsRequest) GetFieldIDs() []string {
	if o == nil || o.FieldIDs == nil {
		var ret []string
		return ret
	}
	return o.FieldIDs
}

// GetFieldIDsOk returns a tuple with the FieldIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetGetFieldsRequest) GetFieldIDsOk() ([]string, bool) {
	if o == nil || o.FieldIDs == nil {
		return nil, false
	}
	return o.FieldIDs, true
}

// HasFieldIDs returns a boolean if a field has been set.
func (o *SmartsheetGetFieldsRequest) HasFieldIDs() bool {
	if o != nil && o.FieldIDs != nil {
		return true
	}

	return false
}

// SetFieldIDs gets a reference to the given []string and assigns it to the FieldIDs field.
func (o *SmartsheetGetFieldsRequest) SetFieldIDs(v []string) {
	o.FieldIDs = v
}

// GetFieldTitles returns the FieldTitles field value if set, zero value otherwise.
func (o *SmartsheetGetFieldsRequest) GetFieldTitles() []string {
	if o == nil || o.FieldTitles == nil {
		var ret []string
		return ret
	}
	return o.FieldTitles
}

// GetFieldTitlesOk returns a tuple with the FieldTitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetGetFieldsRequest) GetFieldTitlesOk() ([]string, bool) {
	if o == nil || o.FieldTitles == nil {
		return nil, false
	}
	return o.FieldTitles, true
}

// HasFieldTitles returns a boolean if a field has been set.
func (o *SmartsheetGetFieldsRequest) HasFieldTitles() bool {
	if o != nil && o.FieldTitles != nil {
		return true
	}

	return false
}

// SetFieldTitles gets a reference to the given []string and assigns it to the FieldTitles field.
func (o *SmartsheetGetFieldsRequest) SetFieldTitles(v []string) {
	o.FieldTitles = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *SmartsheetGetFieldsRequest) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetGetFieldsRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *SmartsheetGetFieldsRequest) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *SmartsheetGetFieldsRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SmartsheetGetFieldsRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetGetFieldsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SmartsheetGetFieldsRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *SmartsheetGetFieldsRequest) SetLimit(v int32) {
	o.Limit = &v
}

func (o SmartsheetGetFieldsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ViewID != nil {
		toSerialize["viewID"] = o.ViewID
	}
	if o.FieldIDs != nil {
		toSerialize["fieldIDs"] = o.FieldIDs
	}
	if o.FieldTitles != nil {
		toSerialize["fieldTitles"] = o.FieldTitles
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetGetFieldsRequest struct {
	value *SmartsheetGetFieldsRequest
	isSet bool
}

func (v NullableSmartsheetGetFieldsRequest) Get() *SmartsheetGetFieldsRequest {
	return v.value
}

func (v *NullableSmartsheetGetFieldsRequest) Set(val *SmartsheetGetFieldsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetGetFieldsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetGetFieldsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetGetFieldsRequest(val *SmartsheetGetFieldsRequest) *NullableSmartsheetGetFieldsRequest {
	return &NullableSmartsheetGetFieldsRequest{value: val, isSet: true}
}

func (v NullableSmartsheetGetFieldsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetGetFieldsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


