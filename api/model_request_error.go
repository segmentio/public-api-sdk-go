/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the RequestError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestError{}

// RequestError Represents any error that could have occurred while performing a request.
type RequestError struct {
	// The type for this error (validation, server, unknown, etc).
	Type string `json:"type"`
	// An error message attached to this error.
	Message *string `json:"message,omitempty"`
	// The name of an input field from the request that triggered this error.
	Field *string `json:"field,omitempty"`
	// Any extra data associated with this error.
	Data interface{} `json:"data,omitempty"`
	// Http status code.
	Status *float32 `json:"status,omitempty"`
}

// NewRequestError instantiates a new RequestError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestError(type_ string) *RequestError {
	this := RequestError{}
	this.Type = type_
	return &this
}

// NewRequestErrorWithDefaults instantiates a new RequestError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestErrorWithDefaults() *RequestError {
	this := RequestError{}
	return &this
}

// GetType returns the Type field value
func (o *RequestError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestError) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RequestError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RequestError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RequestError) SetMessage(v string) {
	o.Message = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *RequestError) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestError) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *RequestError) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *RequestError) SetField(v string) {
	o.Field = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestError) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestError) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RequestError) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *RequestError) SetData(v interface{}) {
	o.Data = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RequestError) GetStatus() float32 {
	if o == nil || IsNil(o.Status) {
		var ret float32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestError) GetStatusOk() (*float32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RequestError) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given float32 and assigns it to the Status field.
func (o *RequestError) SetStatus(v float32) {
	o.Status = &v
}

func (o RequestError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableRequestError struct {
	value *RequestError
	isSet bool
}

func (v NullableRequestError) Get() *RequestError {
	return v.value
}

func (v *NullableRequestError) Set(val *RequestError) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestError) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestError(val *RequestError) *NullableRequestError {
	return &NullableRequestError{value: val, isSet: true}
}

func (v NullableRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
