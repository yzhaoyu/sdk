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

// BatchUpdateDocDataReq struct for BatchUpdateDocDataReq
type BatchUpdateDocDataReq struct {
	FileID string `json:"fileID"`
	Requests []Request `json:"requests"`
}

// NewBatchUpdateDocDataReq instantiates a new BatchUpdateDocDataReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchUpdateDocDataReq(fileID string, requests []Request) *BatchUpdateDocDataReq {
	this := BatchUpdateDocDataReq{}
	this.FileID = fileID
	this.Requests = requests
	return &this
}

// NewBatchUpdateDocDataReqWithDefaults instantiates a new BatchUpdateDocDataReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchUpdateDocDataReqWithDefaults() *BatchUpdateDocDataReq {
	this := BatchUpdateDocDataReq{}
	return &this
}

// GetFileID returns the FileID field value
func (o *BatchUpdateDocDataReq) GetFileID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileID
}

// GetFileIDOk returns a tuple with the FileID field value
// and a boolean to check if the value has been set.
func (o *BatchUpdateDocDataReq) GetFileIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileID, true
}

// SetFileID sets field value
func (o *BatchUpdateDocDataReq) SetFileID(v string) {
	o.FileID = v
}

// GetRequests returns the Requests field value
func (o *BatchUpdateDocDataReq) GetRequests() []Request {
	if o == nil {
		var ret []Request
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *BatchUpdateDocDataReq) GetRequestsOk() ([]Request, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requests, true
}

// SetRequests sets field value
func (o *BatchUpdateDocDataReq) SetRequests(v []Request) {
	o.Requests = v
}

func (o BatchUpdateDocDataReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fileID"] = o.FileID
	}
	if true {
		toSerialize["requests"] = o.Requests
	}
	return json.Marshal(toSerialize)
}

type NullableBatchUpdateDocDataReq struct {
	value *BatchUpdateDocDataReq
	isSet bool
}

func (v NullableBatchUpdateDocDataReq) Get() *BatchUpdateDocDataReq {
	return v.value
}

func (v *NullableBatchUpdateDocDataReq) Set(val *BatchUpdateDocDataReq) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchUpdateDocDataReq) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchUpdateDocDataReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchUpdateDocDataReq(val *BatchUpdateDocDataReq) *NullableBatchUpdateDocDataReq {
	return &NullableBatchUpdateDocDataReq{value: val, isSet: true}
}

func (v NullableBatchUpdateDocDataReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchUpdateDocDataReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


