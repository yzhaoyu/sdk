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

// Status The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
type Status struct {
	// The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Code *int32 `json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message *string `json:"message,omitempty"`
	// A list of messages that carry the error details.  There is a common set of message types for APIs to use.
	Details []GoogleProtobufAny `json:"details,omitempty"`
}

// NewStatus instantiates a new Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus() *Status {
	this := Status{}
	return &this
}

// NewStatusWithDefaults instantiates a new Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusWithDefaults() *Status {
	this := Status{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Status) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Status) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Status) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Status) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Status) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Status) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Status) GetDetails() []GoogleProtobufAny {
	if o == nil || o.Details == nil {
		var ret []GoogleProtobufAny
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetDetailsOk() ([]GoogleProtobufAny, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Status) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []GoogleProtobufAny and assigns it to the Details field.
func (o *Status) SetDetails(v []GoogleProtobufAny) {
	o.Details = v
}

func (o Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


