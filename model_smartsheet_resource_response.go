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

// SmartsheetResourceResponse SmartsheetResourceResponse 处理 Smartsheet 视图、记录、字段资源的响应参数
type SmartsheetResourceResponse struct {
	GetViews *GetViewsResponse `json:"getViews,omitempty"`
	GetRecords *GetRecordsResponse `json:"getRecords,omitempty"`
	GetFields *GetFieldsResponse `json:"getFields,omitempty"`
	AddView *AddViewResponse `json:"addView,omitempty"`
	// 删除指定表格指定子表中的指定视图回包
	DeleteViews map[string]interface{} `json:"deleteViews,omitempty"`
	AddRecords *AddRecordsResponse `json:"addRecords,omitempty"`
	UpdateRecords *UpdateRecordsResponse `json:"updateRecords,omitempty"`
	// 删除指定表格和子表中的一到多条记录回包
	DeleteRecords map[string]interface{} `json:"deleteRecords,omitempty"`
	AddFields *AddFieldsResponse `json:"addFields,omitempty"`
	UpdateFields *UpdateFieldsResponse `json:"updateFields,omitempty"`
	// 删除字段回包
	DeleteFields map[string]interface{} `json:"deleteFields,omitempty"`
}

// NewSmartsheetResourceResponse instantiates a new SmartsheetResourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetResourceResponse() *SmartsheetResourceResponse {
	this := SmartsheetResourceResponse{}
	return &this
}

// NewSmartsheetResourceResponseWithDefaults instantiates a new SmartsheetResourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetResourceResponseWithDefaults() *SmartsheetResourceResponse {
	this := SmartsheetResourceResponse{}
	return &this
}

// GetGetViews returns the GetViews field value if set, zero value otherwise.
func (o *SmartsheetResourceResponse) GetGetViews() GetViewsResponse {
	if o == nil || o.GetViews == nil {
		var ret GetViewsResponse
		return ret
	}
	return *o.GetViews
}

// GetGetViewsOk returns a tuple with the GetViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetResourceResponse) GetGetViewsOk() (*GetViewsResponse, bool) {
	if o == nil || o.GetViews == nil {
		return nil, false
	}
	return o.GetViews, true
}

// HasGetViews returns a boolean if a field has been set.
func (o *SmartsheetResourceResponse) HasGetViews() bool {
	if o != nil && o.GetViews != nil {
		return true
	}

	return false
}

// SetGetViews gets a reference to the given GetViewsResponse and assigns it to the GetViews field.
func (o *SmartsheetResourceResponse) SetGetViews(v GetViewsResponse) {
	o.GetViews = &v
}

// GetGetRecords returns the GetRecords field value if set, zero value otherwise.
func (o *SmartsheetResourceResponse) GetGetRecords() GetRecordsResponse {
	if o == nil || o.GetRecords == nil {
		var ret GetRecordsResponse
		return ret
	}
	return *o.GetRecords
}

// GetGetRecordsOk returns a tuple with the GetRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetResourceResponse) GetGetRecordsOk() (*GetRecordsResponse, bool) {
	if o == nil || o.GetRecords == nil {
		return nil, false
	}
	return o.GetRecords, true
}

// HasGetRecords returns a boolean if a field has been set.
func (o *SmartsheetResourceResponse) HasGetRecords() bool {
	if o != nil && o.GetRecords != nil {
		return true
	}

	return false
}

// SetGetRecords gets a reference to the given GetRecordsResponse and assigns it to the GetRecords field.
func (o *SmartsheetResourceResponse) SetGetRecords(v GetRecordsResponse) {
	o.GetRecords = &v
}

// GetGetFields returns the GetFields field value if set, zero value otherwise.
func (o *SmartsheetResourceResponse) GetGetFields() GetFieldsResponse {
	if o == nil || o.GetFields == nil {
		var ret GetFieldsResponse
		return ret
	}
	return *o.GetFields
}

// GetGetFieldsOk returns a tuple with the GetFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetResourceResponse) GetGetFieldsOk() (*GetFieldsResponse, bool) {
	if o == nil || o.GetFields == nil {
		return nil, false
	}
	return o.GetFields, true
}

// HasGetFields returns a boolean if a field has been set.
func (o *SmartsheetResourceResponse) HasGetFields() bool {
	if o != nil && o.GetFields != nil {
		return true
	}

	return false
}

// SetGetFields gets a reference to the given GetFieldsResponse and assigns it to the GetFields field.
func (o *SmartsheetResourceResponse) SetGetFields(v GetFieldsResponse) {
	o.GetFields = &v
}

// GetAddView returns the AddView field value if set, zero value otherwise.
func (o *SmartsheetResourceResponse) GetAddView() AddViewResponse {
	if o == nil || o.AddView == nil {
		var ret AddViewResponse
		return ret
	}
	return *o.AddView
}

// GetAddViewOk returns a tuple with the AddView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetResourceResponse) GetAddViewOk() (*AddViewResponse, bool) {
	if o == nil || o.AddView == nil {
		return nil, false
	}
	return o.AddView, true
}

// HasAddView returns a boolean if a field has been set.
func (o *SmartsheetResourceResponse) HasAddView() bool {
	if o != nil && o.AddView != nil {
		return true
	}

	return false
}

// SetAddView gets a reference to the given AddViewResponse and assigns it to the AddView field.
func (o *SmartsheetResourceResponse) SetAddView(v AddViewResponse) {
	o.AddView = &v
}

// GetDeleteViews returns the DeleteViews field value if set, zero value otherwise.
func (o *SmartsheetResourceResponse) GetDeleteViews() map[string]interface{} {
	if o == nil || o.DeleteViews == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DeleteViews
}

// GetDeleteViewsOk returns a tuple with the DeleteViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetResourceResponse) GetDeleteViewsOk() (map[string]interface{}, bool) {
	if o == nil || o.DeleteViews == nil {
		return nil, false
	}
	return o.DeleteViews, true
}

// HasDeleteViews returns a boolean if a field has been set.
func (o *SmartsheetResourceResponse) HasDeleteViews() bool {
	if o != nil && o.DeleteViews != nil {
		return true
	}

	return false
}

// SetDeleteViews gets a reference to the given map[string]interface{} and assigns it to the DeleteViews field.
func (o *SmartsheetResourceResponse) SetDeleteViews(v map[string]interface{}) {
	o.DeleteViews = v
}

// GetAddRecords returns the AddRecords field value if set, zero value otherwise.
func (o *SmartsheetResourceResponse) GetAddRecords() AddRecordsResponse {
	if o == nil || o.AddRecords == nil {
		var ret AddRecordsResponse
		return ret
	}
	return *o.AddRecords
}

// GetAddRecordsOk returns a tuple with the AddRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetResourceResponse) GetAddRecordsOk() (*AddRecordsResponse, bool) {
	if o == nil || o.AddRecords == nil {
		return nil, false
	}
	return o.AddRecords, true
}

// HasAddRecords returns a boolean if a field has been set.
func (o *SmartsheetResourceResponse) HasAddRecords() bool {
	if o != nil && o.AddRecords != nil {
		return true
	}

	return false
}

// SetAddRecords gets a reference to the given AddRecordsResponse and assigns it to the AddRecords field.
func (o *SmartsheetResourceResponse) SetAddRecords(v AddRecordsResponse) {
	o.AddRecords = &v
}

// GetUpdateRecords returns the UpdateRecords field value if set, zero value otherwise.
func (o *SmartsheetResourceResponse) GetUpdateRecords() UpdateRecordsResponse {
	if o == nil || o.UpdateRecords == nil {
		var ret UpdateRecordsResponse
		return ret
	}
	return *o.UpdateRecords
}

// GetUpdateRecordsOk returns a tuple with the UpdateRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetResourceResponse) GetUpdateRecordsOk() (*UpdateRecordsResponse, bool) {
	if o == nil || o.UpdateRecords == nil {
		return nil, false
	}
	return o.UpdateRecords, true
}

// HasUpdateRecords returns a boolean if a field has been set.
func (o *SmartsheetResourceResponse) HasUpdateRecords() bool {
	if o != nil && o.UpdateRecords != nil {
		return true
	}

	return false
}

// SetUpdateRecords gets a reference to the given UpdateRecordsResponse and assigns it to the UpdateRecords field.
func (o *SmartsheetResourceResponse) SetUpdateRecords(v UpdateRecordsResponse) {
	o.UpdateRecords = &v
}

// GetDeleteRecords returns the DeleteRecords field value if set, zero value otherwise.
func (o *SmartsheetResourceResponse) GetDeleteRecords() map[string]interface{} {
	if o == nil || o.DeleteRecords == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DeleteRecords
}

// GetDeleteRecordsOk returns a tuple with the DeleteRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetResourceResponse) GetDeleteRecordsOk() (map[string]interface{}, bool) {
	if o == nil || o.DeleteRecords == nil {
		return nil, false
	}
	return o.DeleteRecords, true
}

// HasDeleteRecords returns a boolean if a field has been set.
func (o *SmartsheetResourceResponse) HasDeleteRecords() bool {
	if o != nil && o.DeleteRecords != nil {
		return true
	}

	return false
}

// SetDeleteRecords gets a reference to the given map[string]interface{} and assigns it to the DeleteRecords field.
func (o *SmartsheetResourceResponse) SetDeleteRecords(v map[string]interface{}) {
	o.DeleteRecords = v
}

// GetAddFields returns the AddFields field value if set, zero value otherwise.
func (o *SmartsheetResourceResponse) GetAddFields() AddFieldsResponse {
	if o == nil || o.AddFields == nil {
		var ret AddFieldsResponse
		return ret
	}
	return *o.AddFields
}

// GetAddFieldsOk returns a tuple with the AddFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetResourceResponse) GetAddFieldsOk() (*AddFieldsResponse, bool) {
	if o == nil || o.AddFields == nil {
		return nil, false
	}
	return o.AddFields, true
}

// HasAddFields returns a boolean if a field has been set.
func (o *SmartsheetResourceResponse) HasAddFields() bool {
	if o != nil && o.AddFields != nil {
		return true
	}

	return false
}

// SetAddFields gets a reference to the given AddFieldsResponse and assigns it to the AddFields field.
func (o *SmartsheetResourceResponse) SetAddFields(v AddFieldsResponse) {
	o.AddFields = &v
}

// GetUpdateFields returns the UpdateFields field value if set, zero value otherwise.
func (o *SmartsheetResourceResponse) GetUpdateFields() UpdateFieldsResponse {
	if o == nil || o.UpdateFields == nil {
		var ret UpdateFieldsResponse
		return ret
	}
	return *o.UpdateFields
}

// GetUpdateFieldsOk returns a tuple with the UpdateFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetResourceResponse) GetUpdateFieldsOk() (*UpdateFieldsResponse, bool) {
	if o == nil || o.UpdateFields == nil {
		return nil, false
	}
	return o.UpdateFields, true
}

// HasUpdateFields returns a boolean if a field has been set.
func (o *SmartsheetResourceResponse) HasUpdateFields() bool {
	if o != nil && o.UpdateFields != nil {
		return true
	}

	return false
}

// SetUpdateFields gets a reference to the given UpdateFieldsResponse and assigns it to the UpdateFields field.
func (o *SmartsheetResourceResponse) SetUpdateFields(v UpdateFieldsResponse) {
	o.UpdateFields = &v
}

// GetDeleteFields returns the DeleteFields field value if set, zero value otherwise.
func (o *SmartsheetResourceResponse) GetDeleteFields() map[string]interface{} {
	if o == nil || o.DeleteFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DeleteFields
}

// GetDeleteFieldsOk returns a tuple with the DeleteFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetResourceResponse) GetDeleteFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.DeleteFields == nil {
		return nil, false
	}
	return o.DeleteFields, true
}

// HasDeleteFields returns a boolean if a field has been set.
func (o *SmartsheetResourceResponse) HasDeleteFields() bool {
	if o != nil && o.DeleteFields != nil {
		return true
	}

	return false
}

// SetDeleteFields gets a reference to the given map[string]interface{} and assigns it to the DeleteFields field.
func (o *SmartsheetResourceResponse) SetDeleteFields(v map[string]interface{}) {
	o.DeleteFields = v
}

func (o SmartsheetResourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GetViews != nil {
		toSerialize["getViews"] = o.GetViews
	}
	if o.GetRecords != nil {
		toSerialize["getRecords"] = o.GetRecords
	}
	if o.GetFields != nil {
		toSerialize["getFields"] = o.GetFields
	}
	if o.AddView != nil {
		toSerialize["addView"] = o.AddView
	}
	if o.DeleteViews != nil {
		toSerialize["deleteViews"] = o.DeleteViews
	}
	if o.AddRecords != nil {
		toSerialize["addRecords"] = o.AddRecords
	}
	if o.UpdateRecords != nil {
		toSerialize["updateRecords"] = o.UpdateRecords
	}
	if o.DeleteRecords != nil {
		toSerialize["deleteRecords"] = o.DeleteRecords
	}
	if o.AddFields != nil {
		toSerialize["addFields"] = o.AddFields
	}
	if o.UpdateFields != nil {
		toSerialize["updateFields"] = o.UpdateFields
	}
	if o.DeleteFields != nil {
		toSerialize["deleteFields"] = o.DeleteFields
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetResourceResponse struct {
	value *SmartsheetResourceResponse
	isSet bool
}

func (v NullableSmartsheetResourceResponse) Get() *SmartsheetResourceResponse {
	return v.value
}

func (v *NullableSmartsheetResourceResponse) Set(val *SmartsheetResourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetResourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetResourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetResourceResponse(val *SmartsheetResourceResponse) *NullableSmartsheetResourceResponse {
	return &NullableSmartsheetResourceResponse{value: val, isSet: true}
}

func (v NullableSmartsheetResourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetResourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


